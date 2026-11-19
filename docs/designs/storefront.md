# Copper peptide storefront

Status: draft

This is the product specification for the public ghklabs shop. It is not an accepted design and it is not legal advice. A cosmetics regulatory counsel and a payments underwriter should review the claims, labeling, and raw-material offer before launch.

The accepted webhook API (`webhook-api.md`) stays the outbound event pipe. This document does not change it. A later design can subscribe the storefront to `lot.released` once lots and COAs are operational.

## Positioning

Default, until the open questions are answered:

- The brand is a premium, science-forward **cosmetic** house. The visual language is modern skincare: calm, precise, ingredient-literate. It is not a research-chemical catalog.
- Finished goods are the primary offer and the homepage story. Raw GHK-Cu powder is a separate catalog for formulators, with its own navigation, PDP template, checkout acknowledgment, and claim rules.
- Education explains what the ingredient is, how cosmetic products are made and tested, and how to use a finished serum on skin. It does not teach injection, ingestion, dosing for treatment, or disease outcomes.

## Tech stack

**Recommendation: Next.js App Router storefront on Vercel, Shopify as the commerce system of record, Shopify Checkout (Shopify Payments) for payment, tax, and fraud.**

| Layer | Choice | Why |
| --- | --- | --- |
| Storefront | Next.js App Router, React, TypeScript | Marketing pages, PDPs, and SEO stay in our repo. App Router fits content routes and product routes in one deploy. |
| Styling | Tailwind CSS | Mobile-first layout without a design-system project in the MVP. |
| Commerce | Shopify, headless (Storefront API for catalog and cart; Admin API for COA files and order reads) | Catalog, inventory, customers, discounts, and order management without building them. |
| Checkout | Shopify hosted checkout | PCI scope stays off this app. Tax, Shop Pay, and fraud tools come with the checkout. |
| Content | MDX in the repo for MVP | Guides and posts are reviewed like code. A CMS can replace MDX later without changing the page model. |
| COA files | Shopify Files, linked by product and lot metafields | Buyers download the lot that matches what they can purchase. |
| Payments | Shopify Payments (Stripe is the acquirer underneath in the US) | A separate Stripe integration duplicates checkout and tax. |
| Hosting | Vercel | Already the deploy target for `api/`. Preview URLs per branch. |
| Events | Existing Go webhook service | `lot.released` and, later, order side effects. Not the storefront runtime. |

### Rejected for the MVP

**Medusa (self-hosted).** Real control over the product model and no Shopify transaction fee. The cost is operating Postgres, payments, tax, and fraud ourselves before there is volume. Revisit if Shopify Payments refuses the powder SKUs or the metafield model cannot express lots.

**Custom commerce on the Go service.** The webhook service is a delivery queue. Cart, tax, PCI, and customer accounts do not belong there.

**Shopify Liquid theme.** Faster to launch, worse for the editorial science pages and the two PDP templates. Headless Next.js is the right split: Shopify owns money and inventory; the repo owns the brand and the claim-controlled copy.

### Pros and cons of the recommendation

Pros:

- Checkout, tax, refunds, and customer accounts are already productized.
- Subscriptions (Shopify subscriptions) and B2B (Shopify wholesale / company accounts) have a later path that does not require a replatform.
- COAs, INCI, and claim class fit Shopify metafields and metaobjects.
- The storefront can stay fast: static science pages, cached product JSON, image CDN.

Cons:

- Shopify is a monthly plan plus payment processing. Checkout UI is Checkout Extensibility, not arbitrary React.
- Powder sold with “research use only” language is a **restricted or prohibited** category for many processors, including Shopify Payments and Stripe. Cosmetic serums are an ordinary skincare SKU. See Compliance.
- Metafields are less strict than a custom schema. The storefront must refuse to render a product whose required metafields are missing.

## Information architecture

```text
/                           Home
/shop                       All finished cosmetics
/shop/serums                Serum collection
/shop/raw                   Formulator materials (separate IA)
/products/{handle}          PDP (template chosen by product type)
/science                    Ingredient index
/science/ghk-cu             Pillar guide
/journal                    Blog index
/journal/{slug}             Article
/quality                    Testing, lots, how to read a COA
/quality/coas               Public COA index for released lots
/about
/faq
/contact                    Includes a cosmetic adverse-event path
/cart
/account                    Orders, addresses (Shopify customer accounts)
/legal/terms
/legal/privacy
/legal/shipping
/legal/returns
/legal/cosmetic-disclaimer
/legal/raw-material-notice
```

Checkout, payment, and the order-status page after payment are Shopify-hosted URLs. They are not routes in this app.

Primary navigation: Shop, Science, Quality, Journal, Account. Raw materials are a child of Shop, not a top-level “research chemicals” label.

### Journeys

**Cosmetic buyer.** Home or serum collection → serum PDP (INCI, texture, how to apply, lot COA, cosmetic disclaimer) → cart → Shopify checkout → order in the account. No disease claims on any step, including cart cross-sells.

**Formulator.** Raw collection → powder PDP (identity, purity, net weight, COA, storage, “not a finished cosmetic”) → required acknowledgment checkbox in the cart before checkout is enabled → Shopify checkout. The acknowledgment is stored as a cart attribute so it appears on the order.

**Reader.** Journal or `/science/ghk-cu` → cited, non-therapeutic explanation → link to the finished serum and, in a clearly separated aside, the raw material. Articles do not include preparation for injection or ingestion.

**Returning customer.** Account → past order → reorder the same variant. Subscriptions are out of MVP scope; the account page reserves a section for them.

## Data model

Shopify is the system of record for money and stock. The storefront reads a narrower model and fails closed when required fields are absent.

### Product

| Field | Notes |
| --- | --- |
| `id`, `handle`, `title` | Shopify product. |
| `productType` | `cosmetic` or `raw_material`. Drives template, disclaimer, and whether the cart acknowledgment is required. |
| `status` | Draft products are not queryable from the storefront. |
| `claimClass` | `cosmetic_appearance` or `identity_only`. Copy outside the class is not rendered. |
| `inciName` | Required for cosmetics. Example: Copper Tripeptide-1. Full INCI list for finished goods, in descending concentration order as required on the package. |
| `netContents` | e.g. `30 mL` or `1 g`. |
| `shortDescription` | One sentence, appearance or identity only. |
| `description` | Long copy, reviewed. |
| `usage` | Cosmetics: topical application steps. Raw: storage and handling for formulation. Empty for use-on-the-body of raw powder. |
| `heroMedia` | Images. No before/after medical imagery in the MVP. |
| `seo` | Title and description that obey the same claim class. |

### Variant

| Field | Notes |
| --- | --- |
| `id`, `sku`, `title` | One sellable unit. |
| `optionValues` | Size or pack. Potency is not an option on cosmetics (one formula per product). Powder may vary by pack weight only. |
| `price` | Money. |
| `availableForSale` | False when inventory is zero. No backorder in the MVP. |
| `inventoryPolicy` | Deny. |
| `weight` | For shipping profiles. |

A cosmetic and a raw material never share a product. They can share an ingredient metaobject.

### Ingredient

| Field | Notes |
| --- | --- |
| `name`, `inciName` | Display and label name. |
| `cas` | Optional, shown on raw PDPs. |
| `summary` | What it is, in cosmetic language. |
| `allowedClaims` | List of appearance phrases approved for this ingredient. |
| `bannedContexts` | Injection, ingestion, disease names. Used by editorial review, not as a runtime filter that “sanitizes” unsafe copy into safe copy. Unsafe copy is rejected in review. |

### Lot and COA

| Field | Notes |
| --- | --- |
| `lotCode` | Printed on the unit the buyer receives. |
| `productId` | Parent product. |
| `releasedOn` | Date. |
| `purity` | Raw materials. Method named beside the number (e.g. HPLC). |
| `identityMethod` | e.g. HPLC, MS. |
| `appearance`, `peptideContent` | As reported. Absent fields are omitted, not invented. |
| `documentUrl` | PDF in Shopify Files. |
| `documentSha256` | So a replaced file is detectable. |
| `status` | `released` or `recalled`. Recalled lots stay visible with a recall banner and are not attached to sellable variants. |

The sellable variant points at the current released lot. Older released COAs stay on `/quality/coas`. The webhook `lot.released` event is the future trigger to publish a lot; the MVP can publish lots in Shopify admin.

### Content

| Field | Notes |
| --- | --- |
| `slug`, `title`, `description` | Routes under `/journal` or `/science`. |
| `kind` | `pillar`, `article`, or `faq`. |
| `claimClass` | Same enum as products. Pillar pages about GHK-Cu are `cosmetic_appearance` plus identity. They may summarize published research with citations and an explicit statement that the study outcome is not a product claim. |
| `relatedProducts` | Handles. |
| `reviewedOn` | Date of the last claim review. |
| `body` | MDX. |

### Cart and order

The cart is a Shopify cart: lines, merchandise id, quantity, cost, buyer identity.

Cart attributes:

- `cosmetic_disclaimer_version` when any line is `cosmetic`.
- `raw_acknowledgment` = `accepted` and `raw_notice_version` when any line is `raw_material`. Checkout stays disabled in our UI until this is set. Shopify checkout itself must repeat the notice (checkout UI extension or required product-page policy), because a headless cart attribute is not a legal control by itself.

Orders, refunds, taxes, and tenders are Shopify orders. The MVP does not copy them into the webhook database. A later design can publish `order.paid` if fulfillment needs it.

Customer accounts use Shopify customer accounts. The app does not store passwords.

## Core pages

| Page | Goal | Required content |
| --- | --- | --- |
| Home | Establish the cosmetic brand and route to a serum. | One hero serum, three appearance benefits that match approved claims, a quality proof (test method, not a cure), a link to the GHK-Cu guide. Raw materials are a text link in the footer and shop nav, not the hero. |
| Shop / serums | Scan finished goods. | Cards: name, price, size, one appearance line. No powder in this grid. |
| Raw shop | Sell to formulators without looking like a consumer cure. | Identity, pack weight, purity of the current lot, COA link, raw-material notice. |
| Cosmetic PDP | Convert. | Gallery, price, size, add to cart, INCI, full ingredient list, usage, what the formula is for (look and feel of skin), current lot COA, links to the pillar and the cosmetic disclaimer. |
| Raw PDP | Convert a formulator. | Spec table, COA download, storage, net weight, acknowledgment, statement that the powder is not a finished cosmetic and is not sold for ingestion or injection. No reconstitution protocol. |
| Science pillar | Rank and educate. | Identity, cosmetic role, how to read a label, citations separated from product claims, links to SKUs. |
| Journal | Support SEO with compliant articles. | Index of reviewed posts. |
| Quality / COA index | Make testing inspectable. | Lot, date, product, purity or identity, PDF. |
| Cart | Confirm lines and the correct legal notice. | Line items, quantities, the disclaimer block for the product types present, checkout button. |
| Account | Reorder and read order status. | Shopify-hosted order history. |
| Policies and FAQ | Set expectations. | Shipping, returns, privacy, terms, both product-type notices, adverse-event contact. |
| Contact | Human path. | Support email and a distinct adverse-event note for finished cosmetics (MoCRA). |

## Compliance

This section is a product constraint list for counsel to confirm. It is not a determination of lawful status.

**Finished serums are cosmetics.** Labels and pages may talk about cleansing, beautifying, or altering appearance: moisturize, soften the look of skin, improve the look of skin texture. They may not claim to treat, prevent, or affect disease, and they may not make structure/function claims (collagen production, wound healing, hair regrowth, “repairs DNA”, hormone effects, pain, inflammation as a condition). “Stimulates collagen” and “heals skin” are treated as out of class for this catalog. Citations to wound-healing or anti-aging **studies** stay in the science section, labeled as literature, and are not repeated as what the serum does.

**MoCRA (US cosmetics).** Finished products need a responsible person, facility registration, product listing, adverse-event recording, and labeling that counsel signs off (identity, net quantity, ingredients, warnings, responsible person). The contact page is the public adverse-event door. The website does not replace the package label.

**Raw powder.** Default offer: sold as a cosmetic ingredient to formulators, not as a food, drug, or research chemical for human administration. Pages do not give oral doses, injection volumes, reconstitution for injection, or “not for human consumption” winks that still explain human use. If counsel later requires a research-use-only posture, that is a different product type with its own payment review; it is not the MVP default, because that posture fights the premium-skincare brand and is the phrase processors flag.

**INCI and COA.** Cosmetics show the INCI name Copper Tripeptide-1 where that is the ingredient. Raw PDPs and COAs may also show the common name GHK-Cu. Purity numbers on the site must match the linked COA. A product with no released COA can be drafted in Shopify and must not be `availableForSale` on the storefront.

**Substantiation.** Appearance claims that need evidence (for example a specific texture or hydration claim) ship only with a source counsel accepts. The MVP claim set should be short and boring.

**Payment risk.** Shopify Payments and Stripe restrict many peptide and “research chemical” businesses. A cosmetic serum with Copper Tripeptide-1 is a normal skincare ingredient (widely sold). A jar of bulk peptide with research-only disclaimers can get the whole merchant account closed. The catalog split exists so the powder can be removed from checkout without taking the serums down. Do not add needles, bacteriostatic water, or injection content.

**Marketing surfaces.** Ads, email, SEO titles, image alt text, and reviews are in scope for the same claim class. User reviews that assert disease treatment are not published.

**Age.** No youth-oriented creative. Checkout does not target minors. A numeric age gate is an open question.

## MVP and later phases

**MVP**

- Next.js storefront for the routes above, minus account UI beyond linking to Shopify customer accounts.
- One serum product and one raw-material product, each with one variant and one released COA.
- Cart with the two disclaimer behaviors, Shopify checkout, test-mode purchase.
- Five pages of fixed legal copy, one science pillar, two journal posts, quality index.
- Claim review checklist in the repo (a markdown list a human signs before publish).
- Mobile layout, basic metadata, and a performance budget: LCP-friendly hero, no client-side catalog waterfall.

**Phase 2**

- More SKUs, subscription selling plans, email via Shopify.
- Lot publish flow from `lot.released` into the COA metaobject.
- Reviews with a claim filter.
- Richer account pages if Shopify customer accounts are not enough.

**Phase 3**

- Bulk price breaks.
- Wholesale company accounts and a separate price list.
- International pricing, duties, and translated claim sets. No translated medical claims.

Out of MVP: custom subscriptions, a formulation calculator, a wholesale portal UI, loyalty, and any content that tells a person how to inject or ingest the powder.

## Open questions

Decisions below change the spec. Defaults in this draft are what implementation will follow if they are not changed before the design is accepted.

1. **Buyer.** Is the first revenue cosmetic consumers, formulators, or both? Default: consumers first, formulators as a secondary catalog in the same store.
2. **Powder posture.** Cosmetic ingredient for formulators, or research-use-only? Default: cosmetic ingredient, no research-only branding. Counsel can override.
3. **Catalog at launch.** One hero serum plus one powder pack, or serums only until payments approve powder? Default: build both templates; enable powder checkout only after the processor accepts that SKU.
4. **Geography.** US only at launch? Default: United States shipping only.
5. **Entity.** Legal name, responsible person, and the address that must appear on cosmetic labels and the site. Unknown. Blocking for launch copy, not for this draft.
6. **Accounts.** Shopify customer accounts (default) versus a later custom account on top of Shopify.
7. **Subscriptions.** Confirm they wait until phase 2. Default: yes.
8. **Wholesale.** Same. Default: phase 3, not a second site.
9. **Brand name and visual reference.** Working name is ghklabs. Default look: light, editorial skincare, not a lab catalog. A reference site would lock the next design (visual UI). Draft voice and the competitor read live in [brand-voice.md](brand-voice.md).
10. **Age gate.** Default: no popup gate; no marketing to minors; terms say products are for adults.
11. **Counsel.** Name who signs the claim sheet before a PDP or article can be published. Unknown.

## Acceptance

Accept this document only after questions 1–5 are answered or the defaults are explicitly kept. Implementation of the storefront waits on an accepted revision.
