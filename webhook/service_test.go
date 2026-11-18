package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func testService(t *testing.T, client *http.Client) *Service {
	t.Helper()
	store, err := OpenSQL(context.Background(), "file:"+strings.ReplaceAll(t.Name(), "/", "_")+"?mode=memory&cache=shared")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = store.Close() })
	return NewService(store, Options{
		Client:      client,
		MaxAttempts: 2,
		Backoff:     func(int) time.Duration { return time.Hour },
		Clock:       systemClock{},
	})
}

func TestPublishAndDrain(t *testing.T) {
	var bodies int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodies++
		body, _ := io.ReadAll(r.Body)
		if err := Verify("sek", r.Header.Get(HeaderSignature), body, time.Now(), time.Minute); err != nil {
			t.Error(err)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	svc := testService(t, srv.Client())
	ctx := context.Background()
	if err := svc.Register(ctx, Endpoint{ID: "lab", URL: srv.URL, Secret: "sek", Types: []string{"lot.released"}}); err != nil {
		t.Fatal(err)
	}
	ids, err := svc.Publish(ctx, Event{ID: "evt_9", Type: "lot.released", Data: json.RawMessage(`{"lot":"A"}`)})
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) != 1 {
		t.Fatalf("ids: %v", ids)
	}
	result, err := svc.Drain(ctx, 10)
	if err != nil {
		t.Fatal(err)
	}
	if result.Delivered != 1 || bodies != 1 {
		t.Fatalf("result=%+v bodies=%d", result, bodies)
	}
	again, err := svc.Drain(ctx, 10)
	if err != nil || again.Attempted != 0 {
		t.Fatalf("second drain: %+v %v", again, err)
	}
}

func TestHTTPPublishDeliversImmediately(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()
	svc := testService(t, srv.Client())
	h := Handler{Service: svc, Tokens: []string{"api"}}
	body := fmt.Sprintf(`{"id":"lab","url":%q,"secret":"sek","types":["lot.released"]}`, srv.URL)
	req := httptest.NewRequest(http.MethodPost, "/v1/endpoints", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer api")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("register: %d %s", rec.Code, rec.Body.String())
	}
	req = httptest.NewRequest(http.MethodPost, "/v1/events", strings.NewReader(`{"id":"evt_1","type":"lot.released","data":{"lot":"A"}}`))
	req.Header.Set("Authorization", "Bearer api")
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusAccepted || !strings.Contains(rec.Body.String(), `"delivered":1`) {
		t.Fatalf("publish: %d %s", rec.Code, rec.Body.String())
	}
}

func TestHTTPAuthAndCronToken(t *testing.T) {
	svc := testService(t, http.DefaultClient)
	h := Handler{Service: svc, Tokens: []string{"api", "cron"}}
	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	req.Header.Set("Authorization", "Bearer cron")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("drain with cron token: %d %s", rec.Code, rec.Body.String())
	}
	req = httptest.NewRequest(http.MethodPost, "/v1/events", strings.NewReader(`{"id":"e","type":"ping"}`))
	req.Header.Set("Authorization", "Bearer cron")
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("cron token on publish: %d", rec.Code)
	}
}
