// Package webhook delivers signed event payloads to subscriber endpoints
// with bounded retries and exponential backoff.
package webhook

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	HeaderSignature = "X-GHK-Signature"
	HeaderEventID   = "X-GHK-Event-Id"
	HeaderEventType = "X-GHK-Event-Type"

	DefaultMaxAttempts = 5
)

// Event is the payload delivered to subscribers.
type Event struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	CreatedAt time.Time       `json:"created_at"`
	Data      json.RawMessage `json:"data"`
}

// Endpoint is a subscriber that receives a subset of event types.
// An empty Types list means every event type.
type Endpoint struct {
	ID     string
	URL    string
	Secret string
	Types  []string
}

// Attempt records one HTTP try against an endpoint.
type Attempt struct {
	Number     int
	At         time.Time
	StatusCode int
	Err        string
}

// Delivery is one event destined for one endpoint.
type Delivery struct {
	ID         string
	Event      Event
	EndpointID string
	Attempts   []Attempt
	Status     Status
	NextAt     time.Time
}

// Status is the terminal or in-progress state of a delivery.
type Status string

const (
	StatusPending   Status = "pending"
	StatusDelivered Status = "delivered"
	StatusFailed    Status = "failed"
)

// Clock supplies time so tests can drive backoff without sleeping.
type Clock interface {
	Now() time.Time
}

type systemClock struct{}

func (systemClock) Now() time.Time { return time.Now() }

// Options configures a Dispatcher.
type Options struct {
	MaxAttempts int
	// Backoff returns the delay before attempt n (1-based) is retried after failure.
	// When nil, DefaultBackoff is used.
	Backoff func(attempt int) time.Duration
	Client  *http.Client
	Clock   Clock
}

// DefaultBackoff is 1s, 5s, 30s, 2m, then 10m for later attempts.
func DefaultBackoff(attempt int) time.Duration {
	steps := []time.Duration{
		time.Second,
		5 * time.Second,
		30 * time.Second,
		2 * time.Minute,
		10 * time.Minute,
	}
	if attempt < 1 {
		attempt = 1
	}
	if attempt > len(steps) {
		return steps[len(steps)-1]
	}
	return steps[attempt-1]
}

// Dispatcher stores endpoints and deliveries and posts events.
type Dispatcher struct {
	maxAttempts int
	backoff     func(int) time.Duration
	client      *http.Client
	clock       Clock

	mu          sync.Mutex
	endpoints   map[string]Endpoint
	deliveries  map[string]*Delivery
	nextID      int
}

// NewDispatcher returns a dispatcher ready to accept endpoints and events.
func NewDispatcher(opts Options) *Dispatcher {
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
	return &Dispatcher{
		maxAttempts: maxAttempts,
		backoff:     backoff,
		client:      client,
		clock:       clock,
		endpoints:   map[string]Endpoint{},
		deliveries:  map[string]*Delivery{},
	}
}

// Register stores or replaces an endpoint. ID and URL are required.
func (d *Dispatcher) Register(ep Endpoint) error {
	if ep.ID == "" || ep.URL == "" {
		return errors.New("webhook: endpoint id and url are required")
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	d.endpoints[ep.ID] = ep
	return nil
}

// Publish enqueues a delivery for every endpoint subscribed to ev.Type.
// The event ID is required. CreatedAt defaults to now when zero.
func (d *Dispatcher) Publish(ev Event) ([]string, error) {
	if ev.ID == "" || ev.Type == "" {
		return nil, errors.New("webhook: event id and type are required")
	}
	if ev.CreatedAt.IsZero() {
		ev.CreatedAt = d.clock.Now().UTC()
	}
	if ev.Data == nil {
		ev.Data = json.RawMessage(`null`)
	}

	d.mu.Lock()
	defer d.mu.Unlock()
	var ids []string
	for _, ep := range d.endpoints {
		if !wants(ep, ev.Type) {
			continue
		}
		d.nextID++
		id := fmt.Sprintf("dlv_%d", d.nextID)
		d.deliveries[id] = &Delivery{
			ID:         id,
			Event:      ev,
			EndpointID: ep.ID,
			Status:     StatusPending,
			NextAt:     d.clock.Now(),
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func wants(ep Endpoint, eventType string) bool {
	if len(ep.Types) == 0 {
		return true
	}
	for _, t := range ep.Types {
		if t == eventType {
			return true
		}
	}
	return false
}

// Get returns a copy of a delivery.
func (d *Dispatcher) Get(id string) (Delivery, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	dlv, ok := d.deliveries[id]
	if !ok {
		return Delivery{}, false
	}
	return clone(*dlv), true
}

// Due returns pending deliveries whose next attempt time is at or before now.
func (d *Dispatcher) Due() []Delivery {
	now := d.clock.Now()
	d.mu.Lock()
	defer d.mu.Unlock()
	var out []Delivery
	for _, dlv := range d.deliveries {
		if dlv.Status == StatusPending && !dlv.NextAt.After(now) {
			out = append(out, clone(*dlv))
		}
	}
	return out
}

// DeliverOnce performs a single HTTP attempt for a pending delivery.
// Success marks it delivered. Failure schedules another attempt or marks it failed
// once MaxAttempts is reached. A non-2xx response is a failure.
func (d *Dispatcher) DeliverOnce(ctx context.Context, id string) error {
	d.mu.Lock()
	dlv, ok := d.deliveries[id]
	if !ok {
		d.mu.Unlock()
		return fmt.Errorf("webhook: unknown delivery %s", id)
	}
	if dlv.Status != StatusPending {
		d.mu.Unlock()
		return nil
	}
	ep, ok := d.endpoints[dlv.EndpointID]
	if !ok {
		dlv.Status = StatusFailed
		dlv.Attempts = append(dlv.Attempts, Attempt{
			Number: len(dlv.Attempts) + 1,
			At:     d.clock.Now(),
			Err:    "endpoint removed",
		})
		d.mu.Unlock()
		return errors.New("webhook: endpoint removed")
	}
	body, err := json.Marshal(dlv.Event)
	if err != nil {
		d.mu.Unlock()
		return err
	}
	event := dlv.Event
	attemptN := len(dlv.Attempts) + 1
	d.mu.Unlock()

	status, postErr := d.post(ctx, ep, event, body)

	d.mu.Lock()
	defer d.mu.Unlock()
	dlv = d.deliveries[id]
	att := Attempt{Number: attemptN, At: d.clock.Now(), StatusCode: status}
	if postErr != nil {
		att.Err = postErr.Error()
	}
	dlv.Attempts = append(dlv.Attempts, att)
	if postErr == nil {
		dlv.Status = StatusDelivered
		return nil
	}
	if attemptN >= d.maxAttempts {
		dlv.Status = StatusFailed
		return postErr
	}
	dlv.NextAt = d.clock.Now().Add(d.backoff(attemptN))
	return postErr
}

func (d *Dispatcher) post(ctx context.Context, ep Endpoint, ev Event, body []byte) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep.URL, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(HeaderEventID, ev.ID)
	req.Header.Set(HeaderEventType, ev.Type)
	req.Header.Set(HeaderSignature, Sign(ep.Secret, d.clock.Now(), body))

	resp, err := d.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 1<<16))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp.StatusCode, fmt.Errorf("webhook: endpoint returned %d", resp.StatusCode)
	}
	return resp.StatusCode, nil
}

// Sign returns a Stripe-style signature header value: t=<unix>,v1=<hex hmac>.
// The MAC covers "<unix>.<body>" with the endpoint secret.
func Sign(secret string, at time.Time, body []byte) string {
	ts := at.Unix()
	mac := hmac.New(sha256.New, []byte(secret))
	fmt.Fprintf(mac, "%d.", ts)
	mac.Write(body)
	return fmt.Sprintf("t=%d,v1=%s", ts, hex.EncodeToString(mac.Sum(nil)))
}

// Verify checks a signature header against body and secret.
// Tolerance rejects timestamps further than tolerance from now. Zero tolerance disables the check.
func Verify(secret, header string, body []byte, now time.Time, tolerance time.Duration) error {
	var ts int64
	var sig string
	_, err := fmt.Sscanf(header, "t=%d,v1=%s", &ts, &sig)
	if err != nil || sig == "" {
		return errors.New("webhook: malformed signature")
	}
	if tolerance > 0 {
		delta := now.Sub(time.Unix(ts, 0))
		if delta < 0 {
			delta = -delta
		}
		if delta > tolerance {
			return errors.New("webhook: signature timestamp outside tolerance")
		}
	}
	mac := hmac.New(sha256.New, []byte(secret))
	fmt.Fprintf(mac, "%d.", ts)
	mac.Write(body)
	want := hex.EncodeToString(mac.Sum(nil))
	if !hmac.Equal([]byte(sig), []byte(want)) {
		return errors.New("webhook: signature mismatch")
	}
	return nil
}

func clone(d Delivery) Delivery {
	d.Attempts = append([]Attempt(nil), d.Attempts...)
	d.Event.Data = append(json.RawMessage(nil), d.Event.Data...)
	return d
}
