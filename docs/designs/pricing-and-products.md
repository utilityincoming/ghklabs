# Pricing and product drafts

Status: accepted

Suggested catalog and pricing for the copper-peptide storefront
([storefront.md](storefront.md)), under the reseller/dropship model
([fulfillment-model.md](fulfillment-model.md)) and the voice in
[brand-voice.md](brand-voice.md). Not legal advice, not a payments
determination, not a signed claim sheet. All `{braces}` stay blank until the
label and released COA supply the number.

Prices here are **suggested retail**, chosen to read as premium-but-fair: above
the budget wall, below the clinic brands, with the premium justified by
disclosure (stated %, lot COA, evidence honesty) rather than by romance. They
are proposals for the principals, not committed numbers.

---

## 1. Two cost models (this drives everything)

| Line | Sourcing | COGS basis | Who controls retail |
| --- | --- | --- | --- |
| **Finished serum** (consumer) | Dropship from supply partner | Partner's **wholesale finished-good price** — private, not in repo | Ours **unless** partner sets MAP/fixed (open Q6, fulfillment-model) |
| **Raw GHK-Cu powder** (formulator) | Buy bulk, repack, resell | **Bulk raw-material market price** (table below) | Ours |

The finished serum's wholesale cost is confidential and not recorded here
(fulfillment-model "Confidentiality"). So the serum section below models margin
as a function of that wholesale number `W` and names the retail that keeps a
healthy margin at each `W`. The powder section uses real market pricing.

**Reality check on the active.** A 1% serum at 30 mL carries ~0.3 g of GHK-Cu.
Even at mid cosmetic-grade raw cost (~$2–5/g), that is **$0.60–$1.50 of active
per bottle**. The active is never the cost driver in a finished serum —
packaging, fill, testing, the partner's margin, and payment fees are. Do not let
anyone price the serum "because the peptide is expensive." It isn't. Price it for
position and disclosure.

---

## 2. Raw-material landscape (the table for `/science`)

Bulk/wholesale GHK-Cu (Copper Tripeptide-1, CAS 89030-95-5, 99% claimed). Prices
are supplier listings, Nov 2026, unverified purity. This table feeds a `/science`
explainer — "what you're actually paying for in a copper peptide" — and sets our
powder COGS band.

| Supplier (type) | Grade / claim | 1 g | 100 g | 1 kg | Notes |
| --- | --- | --- | --- | --- | --- |
| Carolina Chemical (US) | ≥99.0%, cosmetic/bulk | $24.80 | $295 ($2.95/g) | $1,600–2,300 | US-based, published COA; realistic cosmetic-grade basis |
| RS Synthesis (US) | 99%, tiered | $75 (1–5 g) | $25/g | on request | Formulator small-batch; steep volume curve |
| Bulk Naturals (US) | 99% powder | $19.56 | — | higher tiers | Aimed at skincare formulators |
| Made-in-China (CN mfr) | 99%, HPLC-MS | $4.34–6.00/g | — | kg ranges unclear | Cosmetic-grade claim; verify COA + CAS |
| China Loong (CN mfr) | 99% | — | — | $40–45/kg (MOQ 1 kg) | Industrial-scale price; grade/documentation must be vetted |
| Etsy marketplace | 99%+, COA | ~$4.14/g (1 kg lot) | — | ~$4,140 | Retail marketplace, not a supply contract |

**Read.** Three price worlds: industrial CN (<$0.05/g, undocumented), bulk
cosmetic-grade (~$1.6–6/g, the serious basis), and repackaged formulator retail
($20–75/g). ghklabs powder COGS planning should assume **~$2–5/g landed,
cosmetic-grade with a usable COA**, not the industrial floor (grade/compliance
risk) and not the formulator retail (that's the price we'd be competing with).

### Size increments the market actually sells

Standard ladder: **1 g → 5 g → 10 g → 50 g → 100 g → 1 kg**. 500 g is
**non-standard** — suppliers jump 100 g straight to 1 kg.

| Supplier | Increments | Per-gram curve |
| --- | --- | --- |
| Carolina Chemical | 1, 5, 10, 100 g, 1 kg | $24.80/g → $2.95/g (100 g) → $1.6–2.3/g (1 kg) |
| RS Synthesis | 1–5, 10, 50, 100 g, kg | $75/g → $25/g (100 g) → quote (kg) |
| Etsy / eBay | mostly 1 g, some to 10 g | $12–35/g |
| Made-in-China | 1 g MOQ → 100 g, then kg | from $3.50/g |

The steep per-gram taper (~$25–75/g at 1 g → ~$3/g at 100 g → ~$2/g at 1 kg) is
the category norm and reads as fairness; a flat per-gram price would look
predatory. **Compliance line:** 100 g and up is wholesale/distributor volume,
not a consumer buyer. A kilogram of peptide with research-ish disclaimers is the
exact SKU that closes a Shopify/Stripe merchant account (storefront.md payment
risk). Those tiers belong on a **B2B/quote channel (invoice or gated wholesale),
never public checkout** beside the serum.

---

## 3. Competitive finished-goods map (where we sit)

| Tier | Examples | Price band (30 mL eq.) | States GHK-Cu %? | Shows lot COA? |
| --- | --- | --- | --- | --- |
| Budget volume | The Ordinary, Good Molecules, COSRX, Asterwood | $12–32 | Rarely | No |
| **ghklabs (proposed)** | — | **$42–58** | **Yes** | **Yes** |
| Prestige / clinic | NIOD CAIS3, koppā, Allies of Skin | $80–95 | NIOD/koppā yes | Supplier spec, not per-lot |

The $42–58 lane is open and defensible: nobody in the budget tier discloses, and
the prestige tier charges 2–3× for it without showing a per-lot certificate. Our
premium is the disclosure, priced fairly.

---

## 4. Margin model

**Fee assumptions (per order, US, Shopify Payments).** Processing ~2.9% + $0.30.
Platform/app amortized ~$1.00/order. Pick/pack/ship is the partner's and is
inside `W` for the serum (dropship) or inside our fulfillment for powder.

Gross margin = Retail − COGS − processing − platform.

### 4a. Serum (COGS = partner wholesale `W`, unknown)

At a suggested hero retail of **$48 (30 mL)**:

| Partner wholesale `W` | Processing | Platform | Gross margin | GM % |
| --- | --- | --- | --- | --- |
| $16 | $1.69 | $1.00 | $29.31 | 61% |
| $19 | $1.69 | $1.00 | $26.31 | 55% |
| $22 | $1.69 | $1.00 | $23.31 | 49% |
| $25 | $1.69 | $1.00 | $20.31 | 42% |

**Target: keep gross margin ≥ 50%**, standard-healthy for DTC skincare that
carries support, returns, and claim-review overhead. That means a **negotiated
wholesale ceiling of ~$21 for a $48 retail.** If the partner's wholesale or MAP
forces a higher floor, move retail up the ladder in §5 rather than eroding
margin below 50%. If MAP *caps* retail below what margin needs, that's a
sourcing conversation, not a copy one.

### 4b. Powder (COGS = our landed raw cost + repack)

Per unit COGS = raw active + vial/label/desiccant + COA handling + inbound.
Assume raw $2–5/g and ~$4/unit repack+COA+inbound.

| SKU | Raw (at $3.5/g) | Repack/COA | Unit COGS | Suggested retail | Processing+plat | Gross margin | GM % |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 g (launch) | $3.50 | $4.00 | $7.50 | **$34** | $2.29 | $24.21 | 71% |
| 5 g (launch) | $17.50 | $5.00 | $22.50 | **$119** | $4.75 | $91.75 | 77% |
| 10 g (later) | $35.00 | $6.00 | $41.00 | $199 | $7.07 | $150.93 | 76% |

Powder margins are high because repackaging is the value-add and we carry the
compliance (COA surfacing, acknowledgment flow). Retail stays **inside the
market band** ($20–75/g at 1 g scale → our $34/g undercuts RS Synthesis and sits
near Carolina/Etsy while beating them on disclosure). Fair, not gouging; the
volume discount (1 g $34 → 10 g $19.90/g) rewards formulators honestly.

---

## 5. Retail ladder (if wholesale/MAP moves the serum)

| Scenario | Serum 30 mL retail | Needed wholesale ceiling for ≥50% GM |
| --- | --- | --- |
| Aggressive (undercut prestige hard) | $42 | ≤ $18 |
| **Base (recommended)** | **$48** | **≤ $21** |
| If partner MAP floors higher | $54 | ≤ $24 |
| Premium ceiling (still < NIOD) | $58 | ≤ $26 |

Do not cross $58 for the 30 mL hero: above it we're arguing with NIOD/koppā on
their turf (named founders, pH bands, routine engineering) without that baggage.

---

## 6. Drafted products

**Focused catalog: one serum + two raw powders. Three SKUs total.** Everything
else (trial size, subscription, a second serum) is deliberately deferred so the
storefront says one clear thing: a disclosed Copper Tripeptide-1 serum, and the
same peptide as a formulator powder in two sizes.

**Supersedes storefront.md MVP (one powder variant → two).** storefront.md MVP
specifies "one raw-material product, one variant." This catalog keeps one powder
*product* but launches **two pack-weight variants** (1 g, 5 g). This is
data-model-legal — storefront.md allows the powder to "vary by pack weight only"
— and storefront.md's MVP line is updated to match in the same change.

### The serum — 1 SKU (`productType: cosmetic`, claimClass `cosmetic_appearance`)

**P1 — Copper Tripeptide-1 Serum 1%  ·  30 mL (target)  ·  $48  (the hero, the whole cosmetic line)**
- INCI headline: Copper Tripeptide-1 ({percent} on label, target 1%).
- Size note: under dropship the 30 mL net contents is the partner's label fact,
  not our pick — confirm the partner's actual fill before publish, same as the %.
- One sentence: "GHK-Cu in a finished, water-based serum — the percent printed
  on the label, the lot's certificate yours to download."
- Appearance claims only: look of fine lines, softer-looking texture, skin that
  feels moisturized. No collagen/heal/repair.
- Variant: 30 mL, `availableForSale` tracks partner. Ships from fulfillment
  partner. Current released lot COA shown as representative.
- Margin: ≥50% GM if negotiated wholesale ≤ $21 (see §4a).

*One serum, one size, one formula.* No 2%/3% variant (invites the
concentration-vs-blend confusion we're positioned to beat), no trial size, no
subscription at launch. Each of those is a later lever, not a launch SKU — keep
the page singular.

### The powders — 2 SKUs (`productType: raw_material`, claimClass `identity_only`)

Separate IA, required cart acknowledgment, no reconstitution/dosing copy. Enable
checkout only after the processor accepts the SKU (storefront.md compliance).

**R1 — GHK-Cu (Copper Tripeptide-1) Powder  ·  1 g  ·  ≥99%  ·  $34**  (entry / try)
**R2 — GHK-Cu (Copper Tripeptide-1) Powder  ·  5 g  ·  ≥99%  ·  $119**  ($23.80/g, the formulator buy)
- Two sizes cover both buyers: 1 g to try the identity and COA, 5 g for someone
  actually formulating. The 5 g undercuts the 1 g per-gram ($23.80 vs $34),
  rewarding the real customer honestly.
- Spec table: identity method, purity of current lot (method named beside the
  number), net weight, storage, "not a finished cosmetic; not for ingestion or
  injection." COA download. No protocol.
- Margins 71–77% (§4b); pricing undercuts formulator retail while beating it on
  COA transparency.

**Bulk tiers — B2B/quote channel, NOT public checkout (later).** If demand
warrants larger volume, offer it off the public storefront to protect the
merchant account (see §2 compliance line). Suggested per-gram framework, margin
compressing fairly with volume:

| Size | Suggested | Per-gram | Channel |
| --- | --- | --- | --- |
| 1 g | $34 | $34.00 | Public (launch) |
| 5 g | $119 | $23.80 | Public (launch) |
| 100 g | ~$790 | $7.90 | B2B / quote |
| 500 g | ~$2,950 | $5.90 | B2B / quote (non-standard size; offer on request only) |
| 1 kg | ~$4,900 | $4.90 | B2B / quote |

These undercut RS Synthesis ($25/g at 100 g) sharply, sit above Carolina's
near-cost ($2.95/g), and keep a healthy but tapering margin. Numbers are a
framework pending confirmed landed COGS (§7 Q2).

---

## 7. Open questions (gate real prices)

1. **Partner wholesale `W` / MAP (Q6, fulfillment-model).** Unknown. The §4a
   table is parametric until this lands. If MAP fixes retail, §5 is moot and the
   only lever is the negotiated wholesale.
2. **Landed raw cost + cosmetic-grade COA source for powder.** Assumed $2–5/g;
   confirm a supplier that ships a usable cosmetic-grade COA with CAS 89030-95-5
   before committing R1–R2 prices.
3. **Repack/compliance cost per powder unit.** Assumed ~$4–6; vial, desiccant,
   label, COA handling, inbound. Confirm.
4. **Processing rate.** Assumed Shopify Payments 2.9% + $0.30; confirm plan.
5. **Powder legal posture.** Prices assume the cosmetic-ingredient posture
   (storefront.md), not research-use-only. Counsel can override and that changes
   the whole SKU.

## 8. Acceptance

**Accepted** by principal direction: the catalog shape (one serum, two powder
pack-weights), the positioning, the two-cost-model margin method, and the
raw-material landscape are the agreed basis for the storefront.

Acceptance of the *structure* does not commit the *numbers*. Open questions Q1
(wholesale/MAP) and Q2 (raw COGS) are still open, so the suggested prices remain
suggestions until those land and the principals sign them off, and the P1 serum
copy still needs a signed claim sheet before publish. Nothing here quotes the
partner's terms.
