# ghklabs

Outbound webhook delivery for ghklabs.

The long-running server is `cmd/ghklabs`. It stores endpoints and deliveries in SQLite (`GHK_DB`, default `data/ghk.db`) or Postgres (`DATABASE_URL`), and drains due work every second.

```sh
GHK_API_TOKEN=dev go run ./cmd/ghklabs
curl -s -H "Authorization: Bearer dev" -H "Content-Type: application/json" \
  -d '{"id":"lab","url":"https://example.com/hooks","secret":"sek","types":["lot.released"]}' \
  http://localhost:8080/v1/endpoints
curl -s -H "Authorization: Bearer dev" -H "Content-Type: application/json" \
  -d '{"id":"evt_1","type":"lot.released","data":{"lot":"A"}}' \
  http://localhost:8080/v1/events
```

`/v1` routes require that bearer token. Drain also accepts `CRON_SECRET`. `POST /v1/events` attempts the new deliveries before it responds. Anything still pending is retried by the local one-second drain, or on Vercel by the daily `GET /api` cron. Set `DATABASE_URL` (Postgres) and both secrets on Vercel. A SQLite file does not persist across Vercel invocations.

`webhook.Dispatcher` is the in-memory library underneath the same signing and retry rules. Signature header: `X-GHK-Signature: t=<unix>,v1=<hex>`, HMAC-SHA256 of `<unix>.<body>`.

```go
d := webhook.NewDispatcher(webhook.Options{})
_ = d.Register(webhook.Endpoint{
    ID: "billing", URL: "https://example.com/hooks", Secret: os.Getenv("HOOK_SECRET"),
    Types: []string{"invoice.paid"},
})
ids, _ := d.Publish(webhook.Event{ID: "evt_1", Type: "invoice.paid", Data: json.RawMessage(`{"id":"in_1"}`)})
_ = d.DeliverOnce(ctx, ids[0])
```

The accepted API design is `docs/designs/webhook-api.md`.

The product page is a static site in `site/` (`docs/designs/marketing-site.md`). It documents the same calls and does not change the API. Serve it with any static file server from that directory. The Vercel rewrite still sends the API deployment at `/` to `api`.

Greenfield project set up for [Elephant-Goldfish Model](https://github.com/google/elephant-goldfish-model) (EGM) design-driven development.

Design documents live in `docs/designs/`. Project context for an Elephant session lives in `_agents/skills/ghklabs_elephant/SKILL.md`.

## Install the plugin

Install the full plugin so skills can read reviewer prompts and document standards. Copying `skills/` alone is not enough.

Cursor support is experimental. From a clean checkout of `google/elephant-goldfish-model`:

```sh
git diff --exit-code HEAD
egm_candidate="$HOME/.cursor/plugins/local/egm"
mkdir -p "$(dirname "$egm_candidate")"
mkdir "$egm_candidate" && {
  cp -R .cursor-plugin skills agents resources docs LICENSE "$egm_candidate/" &&
  git rev-parse HEAD > "$egm_candidate/.egm-source-sha"
}
```

Reload Cursor and confirm the plugin in Customize. Supported hosts (Antigravity, Claude Code, Codex) are documented in that repository's harness guide.

## Workflow

1. Read `_agents/skills/ghklabs_elephant/SKILL.md` at the start of a design session.
2. `/egm-brainstorm` to stress-test an idea and write a concept brief.
3. `/egm-draft` to turn that brief into a design under `docs/designs/`.
4. `/egm-review` for a fresh Goldfish review.
5. Open a pull request for human review, then implement from the design.

Add each accepted design to `docs/designs/SITEMAP.md`.
