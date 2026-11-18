# Webhook publish API

Status: accepted

## Problem

`webhook.Dispatcher` signs and retries in memory. Callers must invoke `DeliverOnce`, and pending work disappears when the process exits. The public site needs a way to enqueue an event and have delivery continue after a restart.

## Behavior

- `POST /v1/endpoints` creates or replaces a subscriber: `id`, `url`, `secret`, and optional `types`. An empty `types` list receives every event.
- `POST /v1/events` stores one pending delivery per matching endpoint and returns their ids. The event `id` and `type` are required. `data` is JSON.
- `POST /v1/drain` and `GET /v1/drain` attempt each delivery that is due. Success marks it delivered. Failure schedules the existing backoff. The attempt limit is still five, then the delivery is failed.
- `GET /v1/deliveries/{id}` returns status and attempts. It does not return the endpoint secret.
- `GET /healthz` is unauthenticated.
- Every `/v1` route requires `Authorization: Bearer <token>`. The token is `GHK_API_TOKEN`. Drain also accepts `CRON_SECRET`, which is what Vercel Cron sends.
- Deliveries and endpoints are stored in SQL. `DATABASE_URL` starting with `postgres://` or `postgresql://` selects Postgres. Any other `GHK_DB` value is a SQLite file, default `data/ghk.db`.
- A drain claims a row by pushing `next_at` forward 30 seconds before the HTTP call, so two drainers do not both send it.
- The long-running server drains on an interval (default 1s). On Vercel, Cron calls `GET /v1/drain` each minute because the function does not stay running.

## Out of scope

Inbound webhook verification for payments, a subscriber UI, and automatic endpoint registration from the catalog.
