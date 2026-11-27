# Marketing site

Status: accepted

The live pages lead with the serum story from [storefront-copy.md](storefront-copy.md): `/` (generic unbranded bottle, [label-truth-up.md](label-truth-up.md) §2.3), `/science`, `/quality`, and `/raw`. Unpublished facts are omitted, not printed as braces. Powder prices and checkout stay off until commercial sign-off. The sections below still describe the original webhook page. That walkthrough is not what these files render.

## Storefront layout

Open with the serum introduction and product render, then give lot verification its own section. Keep ingredient research, label details, and use guidance in distinct sections. Use warm cream for the page, ivory for cards and navigation, and pale blue for lot verification and status. Use deep slate for body copy and a darker blue for links. Carry the same header and dark footer across the storefront.

## Mobile storefront

Keep Serum, Science, Quality, and Powder visible in the narrow header. Use a non-sticky header on narrow screens to leave room for the page content. Put the homepage actions immediately after the product introduction and link to Quality from the lot verification section. Describe the current lot status beside the specimen while the manufacturer file is pending. Provide a keyboard skip link on every page. Serve a smaller hero source when viewport width and pixel density allow.

## Problem

The repository documents an outbound webhook API and ships a Go server. There is no public page that states what the product does, how a delivery moves, or which calls to make first.

## Template

Start from the shape of a shadcn SaaS product site (21st’s Saazio is the closest complete template; the free SaaS landing on the same shelf is the open stand-in). Keep Magic UI / React Bits motion on one artifact — a delivery trace — and keep Mobbin-style hierarchy: one action, then the three calls. Do not adopt agency, portfolio, or storefront templates (Kora, Glass portfolio, YukoHealth). Those fit a studio or a catalog, not a delivery API.

The page is original markup. It does not vendor a paid template.

## Color

The page is light. The blue it attributes is the color of light that has passed through GHK-Cu.

Cu(II)GHK at pH 7.4 has a d–d maximum at 606 nm (Ufnalska et al., Inorg. Chem. 2021). A broad Gaussian on that center, 70 nm wide, under D65 transmits sRGB `#6c9ef9`, hue about 226°. That hex is `--signal` and `--solution`. It is the fill for buttons, the delivered state, and the trace glow. Unbound aqua copper absorbs near 816 nm and looks cyan. That cyan is not this blue.

`#6c9ef9` on the paper is about 2.5:1, so it is not used for small type. Kickers, the wordmark, and the signature label use `#1f4eae` (`--ink-blue`), the same hue darkened until it clears about 7:1 on `#f5f8fc`. Body text is `#102033`. Warning and failure are `#c56a12` and `#c2412d` so they stay distinct on the light ground.

## Behavior

- A static site at `site/` explains outbound delivery: register an endpoint, publish an event, drain what is still pending.
- The hero states the product in one line and shows the signature header used by the server: `X-GHK-Signature: t=<unix>,v1=<hex>`, HMAC-SHA256 of `<unix>.<body>`.
- A delivery trace walks the states the API actually uses: pending, attempt, delivered, retry, failed after five attempts. It is a walkthrough in the browser. It does not call the API.
- The API section shows the three calls from the README: create an endpoint, publish an event, and the bearer requirement. Copy buttons copy those samples.
- Runtime notes stay factual: SQLite via `GHK_DB`, Postgres via `DATABASE_URL`, long-running drain about once a second, Vercel cron once a day through `GET /api`. A SQLite file does not persist on Vercel.
- No pricing, logos, testimonials, or subscriber console. Those are out of scope in the webhook API design.
- On Vercel, `/` serves this page. `/styles.css`, `/trace.js`, and `/favicon.svg` are the files in `site/`. `/v1/*` and `/healthz` still reach the Go handler. Cron still calls `GET /api`. The long-running server is unchanged: `GET /` there remains the health JSON.

## Out of scope

Accounts, a live dispatcher in the browser, and changing `/v1` behavior.
