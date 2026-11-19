# Webhook publish API

Status: accepted

## Problem

`webhook.Dispatcher` signs and retries in memory. Callers must invoke `DeliverOnce`, and pending work disappears when the process exits. The public site needs a way to enqueue an event and have delivery continue after a restart.

## Behavior

- `POST /v1/endpoints` creates or replaces a subscriber: `id`, `url`, `secret`, and optional `types`. An empty `types` list receives every event.
- `POST /v1/events` stores one pending delivery per matching endpoint, attempts those deliveries before responding, and returns their ids plus the drain counts. The event `id` and `type` are required. `data` is JSON. Deliveries that fail this attempt stay pending for a later drain.
- `POST /v1/drain` and `GET /v1/drain` attempt each delivery that is due. Success marks it delivered. Failure schedules the existing backoff. The attempt limit is still five, then the delivery is failed.
- `GET /v1/deliveries/{id}` returns status and attempts. It does not return the endpoint secret.
- `GET /healthz` and `GET /` are unauthenticated and return `{"status":"ok"}` from the Go handler. On Vercel, `/` is the product page, so the public health check there is `GET /healthz`.
- A Vercel instance retries database setup on the next request when the first attempt fails. A successful setup is kept for the life of that instance.
- Every `/v1` route requires `Authorization: Bearer <token>`. The token is `GHK_API_TOKEN`. Drain also accepts `CRON_SECRET`, which is what Vercel Cron sends.
- Deliveries and endpoints are stored in SQL. `DATABASE_URL` starting with `postgres://` or `postgresql://` selects Postgres. Any other `GHK_DB` value is a SQLite file, default `data/ghk.db`.
- A drain claims a row by pushing `next_at` forward 30 seconds before the HTTP call, so two drainers do not both send it.
- The long-running server drains on an interval (default 1s). On Vercel, Cron calls `GET /api` once a day, which runs the same drain. Hobby plans reject schedules that run more often than daily.

## Out of scope

Inbound webhook verification for payments, a subscriber UI, and automatic endpoint registration from the catalog.
