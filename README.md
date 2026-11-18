# ghklabs

Webhook delivery for outbound event notifications.

`webhook.Dispatcher` registers subscriber endpoints, publishes events, and POSTs a JSON body with an HMAC-SHA256 signature. Failed attempts stay pending and become due again after backoff. A delivery is marked failed after the attempt limit.

Signature header: `X-GHK-Signature: t=<unix>,v1=<hex>`, where `v1` is HMAC-SHA256 of `<unix>.<body>` using the endpoint secret. Receivers check it with `webhook.Verify`.

```go
d := webhook.NewDispatcher(webhook.Options{})
_ = d.Register(webhook.Endpoint{
    ID: "billing", URL: "https://example.com/hooks", Secret: os.Getenv("HOOK_SECRET"),
    Types: []string{"invoice.paid"},
})
ids, _ := d.Publish(webhook.Event{ID: "evt_1", Type: "invoice.paid", Data: json.RawMessage(`{"id":"in_1"}`)})
_ = d.DeliverOnce(ctx, ids[0])
```
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
