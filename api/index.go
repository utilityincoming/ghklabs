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
	once    sync.Once
	handler http.Handler
	initErr error
)

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			dsn = os.Getenv("GHK_DB")
		}
		if dsn == "" {
			initErr = errNoDatabase
			return
		}
		store, err := webhook.OpenSQL(context.Background(), dsn)
		if err != nil {
			initErr = err
			return
		}
		svc := webhook.NewService(store, webhook.Options{})
		handler = webhook.Handler{Service: svc, Tokens: []string{os.Getenv("GHK_API_TOKEN"), os.Getenv("CRON_SECRET")}}
	})
	if initErr != nil {
		http.Error(w, initErr.Error(), http.StatusServiceUnavailable)
		return
	}
	handler.ServeHTTP(w, r)
}

var errNoDatabase = errString("DATABASE_URL or GHK_DB is required")

type errString string

func (e errString) Error() string { return string(e) }
