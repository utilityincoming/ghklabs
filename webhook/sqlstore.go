package webhook

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "modernc.org/sqlite"
)

// SQLStore persists endpoints and deliveries in SQLite or Postgres.
type SQLStore struct {
	db       *sql.DB
	postgres bool
}

// OpenSQL opens Postgres when dsn starts with postgres:// or postgresql://.
// Any other dsn is a SQLite path or SQLite URI.
func OpenSQL(ctx context.Context, dsn string) (*SQLStore, error) {
	postgres := strings.HasPrefix(dsn, "postgres://") || strings.HasPrefix(dsn, "postgresql://")
	driver := "sqlite"
	if postgres {
		driver = "pgx"
	} else if !strings.Contains(dsn, "mode=memory") && !strings.Contains(dsn, "?") {
		dsn += "?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)"
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	if !postgres {
		db.SetMaxOpenConns(1)
	}
	s := &SQLStore{db: db, postgres: postgres}
	if err := s.migrate(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}
	return s, nil
}

// Close closes the database.
func (s *SQLStore) Close() error { return s.db.Close() }

func (s *SQLStore) migrate(ctx context.Context) error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS endpoints (
  id TEXT PRIMARY KEY,
  url TEXT NOT NULL,
  secret TEXT NOT NULL,
  types TEXT NOT NULL
)`,
		`CREATE TABLE IF NOT EXISTS deliveries (
  id TEXT PRIMARY KEY,
  event_json TEXT NOT NULL,
  endpoint_id TEXT NOT NULL,
  status TEXT NOT NULL,
  next_at_ms BIGINT NOT NULL,
  attempts_json TEXT NOT NULL
)`,
	}
	for _, stmt := range stmts {
		if _, err := s.db.ExecContext(ctx, stmt); err != nil {
			return err
		}
	}
	return nil
}

func (s *SQLStore) q(query string) string {
	if !s.postgres {
		return query
	}
	var b strings.Builder
	n := 1
	for _, r := range query {
		if r == '?' {
			fmt.Fprintf(&b, "$%d", n)
			n++
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// UpsertEndpoint inserts or replaces an endpoint.
func (s *SQLStore) UpsertEndpoint(ctx context.Context, ep Endpoint) error {
	types := strings.Join(ep.Types, ",")
	_, err := s.db.ExecContext(ctx, s.q(`
INSERT INTO endpoints (id, url, secret, types) VALUES (?, ?, ?, ?)
ON CONFLICT (id) DO UPDATE SET url = excluded.url, secret = excluded.secret, types = excluded.types`),
		ep.ID, ep.URL, ep.Secret, types)
	return err
}

// ListEndpoints returns every endpoint.
func (s *SQLStore) ListEndpoints(ctx context.Context) ([]Endpoint, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT id, url, secret, types FROM endpoints ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Endpoint
	for rows.Next() {
		var ep Endpoint
		var types string
		if err := rows.Scan(&ep.ID, &ep.URL, &ep.Secret, &types); err != nil {
			return nil, err
		}
		if types != "" {
			ep.Types = strings.Split(types, ",")
		}
		out = append(out, ep)
	}
	return out, rows.Err()
}

// Endpoint returns one endpoint.
func (s *SQLStore) Endpoint(ctx context.Context, id string) (Endpoint, bool, error) {
	var ep Endpoint
	var types string
	err := s.db.QueryRowContext(ctx, s.q(`SELECT id, url, secret, types FROM endpoints WHERE id = ?`), id).
		Scan(&ep.ID, &ep.URL, &ep.Secret, &types)
	if err == sql.ErrNoRows {
		return Endpoint{}, false, nil
	}
	if err != nil {
		return Endpoint{}, false, err
	}
	if types != "" {
		ep.Types = strings.Split(types, ",")
	}
	return ep, true, nil
}

// InsertDelivery stores a new pending delivery.
func (s *SQLStore) InsertDelivery(ctx context.Context, dlv Delivery) error {
	ev, err := json.Marshal(dlv.Event)
	if err != nil {
		return err
	}
	attempts, err := json.Marshal(dlv.Attempts)
	if err != nil {
		return err
	}
	if attempts == nil {
		attempts = []byte("[]")
	}
	_, err = s.db.ExecContext(ctx, s.q(`
INSERT INTO deliveries (id, event_json, endpoint_id, status, next_at_ms, attempts_json)
VALUES (?, ?, ?, ?, ?, ?)`),
		dlv.ID, string(ev), dlv.EndpointID, string(dlv.Status), dlv.NextAt.UnixMilli(), string(attempts))
	return err
}

// GetDelivery loads one delivery.
func (s *SQLStore) GetDelivery(ctx context.Context, id string) (Delivery, bool, error) {
	var eventJSON, attemptsJSON, status, endpointID string
	var nextAt int64
	err := s.db.QueryRowContext(ctx, s.q(`
SELECT event_json, endpoint_id, status, next_at_ms, attempts_json FROM deliveries WHERE id = ?`), id).
		Scan(&eventJSON, &endpointID, &status, &nextAt, &attemptsJSON)
	if err == sql.ErrNoRows {
		return Delivery{}, false, nil
	}
	if err != nil {
		return Delivery{}, false, err
	}
	return decodeDelivery(id, endpointID, status, nextAt, eventJSON, attemptsJSON)
}

// DueIDs returns pending delivery ids whose next attempt is at or before now.
func (s *SQLStore) DueIDs(ctx context.Context, now time.Time, limit int) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, s.q(`
SELECT id FROM deliveries WHERE status = ? AND next_at_ms <= ? ORDER BY next_at_ms LIMIT ?`),
		string(StatusPending), now.UnixMilli(), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, rows.Err()
}

// Claim moves next_at to until so a second drainer skips this row. It reports whether this caller won the claim.
func (s *SQLStore) Claim(ctx context.Context, id string, now, until time.Time) (bool, error) {
	res, err := s.db.ExecContext(ctx, s.q(`
UPDATE deliveries SET next_at_ms = ? WHERE id = ? AND status = ? AND next_at_ms <= ?`),
		until.UnixMilli(), id, string(StatusPending), now.UnixMilli())
	if err != nil {
		return false, err
	}
	n, err := res.RowsAffected()
	return n == 1, err
}

// SaveDelivery writes status, schedule, and attempts after an attempt.
func (s *SQLStore) SaveDelivery(ctx context.Context, dlv Delivery) error {
	attempts, err := json.Marshal(dlv.Attempts)
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx, s.q(`
UPDATE deliveries SET status = ?, next_at_ms = ?, attempts_json = ? WHERE id = ?`),
		string(dlv.Status), dlv.NextAt.UnixMilli(), string(attempts), dlv.ID)
	return err
}

func decodeDelivery(id, endpointID, status string, nextAt int64, eventJSON, attemptsJSON string) (Delivery, bool, error) {
	var ev Event
	if err := json.Unmarshal([]byte(eventJSON), &ev); err != nil {
		return Delivery{}, false, err
	}
	var attempts []Attempt
	if attemptsJSON != "" {
		if err := json.Unmarshal([]byte(attemptsJSON), &attempts); err != nil {
			return Delivery{}, false, err
		}
	}
	return Delivery{
		ID:         id,
		Event:      ev,
		EndpointID: endpointID,
		Status:     Status(status),
		NextAt:     time.UnixMilli(nextAt).UTC(),
		Attempts:   attempts,
	}, true, nil
}
