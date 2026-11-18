package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/utilityincoming/ghklabs/webhook"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("GHK_DB")
	}
	if dsn == "" {
		dsn = "data/ghk.db"
	}
	if dsn == "data/ghk.db" {
		if err := os.MkdirAll("data", 0o755); err != nil {
			log.Fatal(err)
		}
	}
	store, err := webhook.OpenSQL(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	svc := webhook.NewService(store, webhook.Options{})
	handler := webhook.Handler{Service: svc, Tokens: []string{os.Getenv("GHK_API_TOKEN"), os.Getenv("CRON_SECRET")}}

	interval := time.Second
	if v := os.Getenv("GHK_DRAIN_INTERVAL"); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			log.Fatal(err)
		}
		interval = d
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			if _, err := svc.Drain(context.Background(), 50); err != nil {
				log.Printf("drain: %v", err)
			}
		}
	}()

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
