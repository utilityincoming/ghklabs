package webhook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const claimLease = 30 * time.Second

// Service publishes events into SQL and drains due deliveries.
type Service struct {
	store       *SQLStore
	maxAttempts int
	backoff     func(int) time.Duration
	client      *http.Client
	clock       Clock

	mu     sync.Mutex
	nextID int
}

// NewService returns a service. Nil options use the same defaults as Dispatcher.
func NewService(store *SQLStore, opts Options) *Service {
	maxAttempts := opts.MaxAttempts
	if maxAttempts < 1 {
		maxAttempts = DefaultMaxAttempts
	}
	backoff := opts.Backoff
	if backoff == nil {
		backoff = DefaultBackoff
	}
	client := opts.Client
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}
	clock := opts.Clock
	if clock == nil {
		clock = systemClock{}
	}
	return &Service{
		store:       store,
		maxAttempts: maxAttempts,
		backoff:     backoff,
		client:      client,
		clock:       clock,
	}
}

// Register stores an endpoint.
func (s *Service) Register(ctx context.Context, ep Endpoint) error {
	if ep.ID == "" || ep.URL == "" {
		return errors.New("webhook: endpoint id and url are required")
	}
	return s.store.UpsertEndpoint(ctx, ep)
}

// Publish enqueues a delivery for every matching endpoint.
func (s *Service) Publish(ctx context.Context, ev Event) ([]string, error) {
	if ev.ID == "" || ev.Type == "" {
		return nil, errors.New("webhook: event id and type are required")
	}
	if ev.CreatedAt.IsZero() {
		ev.CreatedAt = s.clock.Now().UTC()
	}
	if ev.Data == nil {
		ev.Data = json.RawMessage(`null`)
	}
	eps, err := s.store.ListEndpoints(ctx)
	if err != nil {
		return nil, err
	}
	now := s.clock.Now()
	var ids []string
	for _, ep := range eps {
		if !wants(ep, ev.Type) {
			continue
		}
		id := s.newID(now)
		if err := s.store.InsertDelivery(ctx, Delivery{
			ID:         id,
			Event:      ev,
			EndpointID: ep.ID,
			Status:     StatusPending,
			NextAt:     now,
		}); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (s *Service) newID(now time.Time) string {
	s.mu.Lock()
	s.nextID++
	n := s.nextID
	s.mu.Unlock()
	return fmt.Sprintf("dlv_%d_%d", now.UnixNano(), n)
}

// Get returns a delivery.
func (s *Service) Get(ctx context.Context, id string) (Delivery, bool, error) {
	return s.store.GetDelivery(ctx, id)
}

// DrainResult counts one drain pass.
type DrainResult struct {
	Attempted int `json:"attempted"`
	Delivered int `json:"delivered"`
	Failed    int `json:"failed"`
	Pending   int `json:"pending"`
}

// Drain attempts up to limit due deliveries. Limit defaults to 50.
func (s *Service) Drain(ctx context.Context, limit int) (DrainResult, error) {
	if limit < 1 {
		limit = 50
	}
	now := s.clock.Now()
	ids, err := s.store.DueIDs(ctx, now, limit)
	if err != nil {
		return DrainResult{}, err
	}
	var result DrainResult
	for _, id := range ids {
		ok, err := s.store.Claim(ctx, id, now, now.Add(claimLease))
		if err != nil {
			return result, err
		}
		if !ok {
			continue
		}
		status, err := s.deliverClaimed(ctx, id)
		if err != nil && status == "" {
			return result, err
		}
		result.Attempted++
		switch status {
		case StatusDelivered:
			result.Delivered++
		case StatusFailed:
			result.Failed++
		default:
			result.Pending++
		}
	}
	return result, nil
}

func (s *Service) deliverClaimed(ctx context.Context, id string) (Status, error) {
	dlv, ok, err := s.store.GetDelivery(ctx, id)
	if err != nil || !ok {
		if !ok {
			return "", fmt.Errorf("webhook: claimed delivery %s missing", id)
		}
		return "", err
	}
	ep, found, err := s.store.Endpoint(ctx, dlv.EndpointID)
	if err != nil {
		return "", err
	}
	body, err := json.Marshal(dlv.Event)
	if err != nil {
		return "", err
	}
	attemptN := len(dlv.Attempts) + 1
	var statusCode int
	var postErr error
	if !found {
		postErr = errors.New("endpoint removed")
	} else {
		statusCode, postErr = postEvent(ctx, s.client, s.clock, ep, dlv.Event, body)
	}
	att := Attempt{Number: attemptN, At: s.clock.Now(), StatusCode: statusCode}
	if postErr != nil {
		att.Err = postErr.Error()
	}
	dlv.Attempts = append(dlv.Attempts, att)
	switch {
	case postErr == nil:
		dlv.Status = StatusDelivered
	case attemptN >= s.maxAttempts:
		dlv.Status = StatusFailed
	default:
		dlv.Status = StatusPending
		dlv.NextAt = s.clock.Now().Add(s.backoff(attemptN))
	}
	if err := s.store.SaveDelivery(ctx, dlv); err != nil {
		return "", err
	}
	return dlv.Status, nil
}
