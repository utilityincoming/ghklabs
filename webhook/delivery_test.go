package webhook

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

type fakeClock struct {
	mu  sync.Mutex
	now time.Time
}

func (c *fakeClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *fakeClock) Advance(d time.Duration) {
	c.mu.Lock()
	c.now = c.now.Add(d)
	c.mu.Unlock()
}

func TestPublishFiltersTypesAndSigns(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1_700_000_000, 0).UTC()}
	var gotBody []byte
	var gotSig, gotType string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotSig = r.Header.Get(HeaderSignature)
		gotType = r.Header.Get(HeaderEventType)
		gotBody, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	d := NewDispatcher(Options{Clock: clock, Client: srv.Client()})
	if err := d.Register(Endpoint{ID: "orders", URL: srv.URL, Secret: "sek", Types: []string{"order.created"}}); err != nil {
		t.Fatal(err)
	}
	if err := d.Register(Endpoint{ID: "billing", URL: srv.URL, Secret: "sek", Types: []string{"invoice.paid"}}); err != nil {
		t.Fatal(err)
	}

	ids, err := d.Publish(Event{ID: "evt_1", Type: "order.created", Data: json.RawMessage(`{"n":1}`)})
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) != 1 {
		t.Fatalf("deliveries: got %d, want 1", len(ids))
	}
	if err := d.DeliverOnce(context.Background(), ids[0]); err != nil {
		t.Fatal(err)
	}
	dlv, ok := d.Get(ids[0])
	if !ok || dlv.Status != StatusDelivered {
		t.Fatalf("status: %+v ok=%v", dlv.Status, ok)
	}
	if gotType != "order.created" {
		t.Fatalf("type header: %s", gotType)
	}
	if err := Verify("sek", gotSig, gotBody, clock.Now(), time.Minute); err != nil {
		t.Fatal(err)
	}
	if err := Verify("other", gotSig, gotBody, clock.Now(), time.Minute); err == nil {
		t.Fatal("expected mismatch")
	}
}

func TestRetryThenFail(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1_700_000_000, 0).UTC()}
	var hits int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits++
		w.WriteHeader(http.StatusBadGateway)
	}))
	defer srv.Close()

	d := NewDispatcher(Options{
		Clock:       clock,
		Client:      srv.Client(),
		MaxAttempts: 3,
		Backoff:     func(int) time.Duration { return time.Second },
	})
	if err := d.Register(Endpoint{ID: "all", URL: srv.URL, Secret: "s"}); err != nil {
		t.Fatal(err)
	}
	ids, err := d.Publish(Event{ID: "evt_2", Type: "ping"})
	if err != nil {
		t.Fatal(err)
	}

	if err := d.DeliverOnce(context.Background(), ids[0]); err == nil {
		t.Fatal("expected failure")
	}
	dlv, _ := d.Get(ids[0])
	if dlv.Status != StatusPending {
		t.Fatalf("status after first failure: %s", dlv.Status)
	}
	if len(d.Due()) != 0 {
		t.Fatal("should not be due until backoff elapses")
	}
	clock.Advance(time.Second)
	if len(d.Due()) != 1 {
		t.Fatal("expected due delivery")
	}
	_ = d.DeliverOnce(context.Background(), ids[0])
	_ = d.DeliverOnce(context.Background(), ids[0])
	dlv, _ = d.Get(ids[0])
	if dlv.Status != StatusFailed {
		t.Fatalf("status: %s", dlv.Status)
	}
	if hits != 3 || len(dlv.Attempts) != 3 {
		t.Fatalf("hits=%d attempts=%d", hits, len(dlv.Attempts))
	}
}

func TestSignRoundTripTolerance(t *testing.T) {
	body := []byte(`{"ok":true}`)
	at := time.Unix(1_700_000_000, 0)
	header := Sign("secret", at, body)
	if err := Verify("secret", header, body, at.Add(30*time.Second), time.Minute); err != nil {
		t.Fatal(err)
	}
	if err := Verify("secret", header, body, at.Add(2*time.Minute), time.Minute); err == nil {
		t.Fatal("expected tolerance failure")
	}
}
