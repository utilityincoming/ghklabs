// Package api is the Vercel Go entrypoint. Cron should call GET /v1/drain.
package api

import (
	"context"
	"net/http"
	"os"
	"sync"

	"github.com/utilityincoming/ghklabs/webhook"
)

var (
	mu      sync.Mutex
	handler http.Handler
)

func Handler(w http.ResponseWriter, r *http.Request) {
	h, err := currentHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	h.ServeHTTP(w, r)
}

func currentHandler() (http.Handler, error) {
	mu.Lock()
	defer mu.Unlock()
	if handler != nil {
		return handler, nil
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("GHK_DB")
	}
	if dsn == "" {
		return nil, errNoDatabase
	}
	store, err := webhook.OpenSQL(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	svc := webhook.NewService(store, webhook.Options{})
	handler = webhook.Handler{Service: svc, Tokens: []string{os.Getenv("GHK_API_TOKEN"), os.Getenv("CRON_SECRET")}}
	return handler, nil
}

var errNoDatabase = errString("DATABASE_URL or GHK_DB is required")

type errString string

func (e errString) Error() string { return string(e) }
