# Marketing site

Status: accepted

## Problem

The repository documents an outbound webhook API and ships a Go server. There is no public page that states what the product does, how a delivery moves, or which calls to make first.

## Template

Start from the shape of a shadcn SaaS product site (21st’s Saazio is the closest complete template; the free SaaS landing on the same shelf is the open stand-in). Keep Magic UI / React Bits motion on one artifact — a delivery trace — and keep Mobbin-style hierarchy: one action, then the three calls. Do not adopt agency, portfolio, or storefront templates (Kora, Glass portfolio, YukoHealth). Those fit a studio or a catalog, not a delivery API.

The page is original markup. It does not vendor a paid template.

## Behavior

- A static site at `site/` explains outbound delivery: register an endpoint, publish an event, drain what is still pending.
- The hero states the product in one line and shows the signature header used by the server: `X-GHK-Signature: t=<unix>,v1=<hex>`, HMAC-SHA256 of `<unix>.<body>`.
- A delivery trace walks the states the API actually uses: pending, attempt, delivered, retry, failed after five attempts. It is a walkthrough in the browser. It does not call the API.
- The API section shows the three calls from the README: create an endpoint, publish an event, and the bearer requirement. Copy buttons copy those samples.
- Runtime notes stay factual: SQLite via `GHK_DB`, Postgres via `DATABASE_URL`, long-running drain about once a second, Vercel cron once a day through `GET /api`. A SQLite file does not persist on Vercel.
- No pricing, logos, testimonials, or subscriber console. Those are out of scope in the webhook API design.
- The Go API and `vercel.json` rewrite stay the API deployment. This site is a separate static project.

## Out of scope

Accounts, a live dispatcher in the browser, and changing `/v1` behavior.
