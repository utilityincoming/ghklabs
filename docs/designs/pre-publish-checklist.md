# Pre-publish checklist

Status: draft

The gate a human signs before any storefront page, SKU, or lot goes live. It exists because the house is a **reseller of a cosmetic** (see [fulfillment-model.md](fulfillment-model.md)): we own every claim and every number we publish even though we do not make or test the goods. This is a working checklist, not legal advice. Counsel signs the claim sheet; this list is what the operator confirms before asking counsel and before publish.

Run it once per publishable unit: a page, a SKU, or a newly released lot. A `no` or `unknown` on any required line blocks publish. Keep the signed copy with the release.

## Gate A — Facts are filled, not invented

Every `{brace}` in the copy is a fact from a document, left blank until that document exists. Nothing is estimated, rounded, or carried over from another lot. Source each one:

| Fact | From | Checked |
| --- | --- | --- |
| `{percent}` | The finished-product label / formula sheet | ☐ |
| `{size}` | Label (net contents) | ☐ |
| `{range}` (pH) | Formula sheet | ☐ |
| `{lot}` | The unit being sold | ☐ |
| `{date}` (released) | The lot's COA | ☐ |
| `{method}` (identity) | The lot's COA | ☐ |
| `{purity}` | The lot's COA | ☐ |
| `{price}` | Commercial setup | ☐ |
| `{handling_time}`, `{lead_time}` | The supply partner's shipping terms | ☐ |
| `{return_window}`, `{return_condition}`, `{refund_time}`, `{shipping_refund_rule}` | The supply partner's RMA terms + our policy | ☐ |
| `{support_email}` | Operating setup | ☐ |
| `{entity}`, `{address}` | Legal entity / responsible person | ☐ |

- ☐ No placeholder `{…}` is left rendered on any live page.
- ☐ Every number on the page matches the document it came from, character for character.

## Gate B — Claim discipline

- ☐ Page declares its claim class: `cosmetic_appearance` (serum, home, PDP) or `identity_only` (raw powder).
- ☐ No banned product-page word used as a claim: *heal, repair, recover, reverse, stimulate/boost collagen, DNA, gene, inflammation, wound, hair growth, clinical proof, proven, doctor-recommended, biohack (verb), research use only, injectable, subcutaneous, reconstitute, dose, protocol.* (A scan is in Gate F.)
- ☐ Appearance language only: "the look of," "feels," "softer-looking," "moisturized feel." No structure/function or disease claims.
- ☐ Any study is a labeled citation on a science page, not repeated as what the serum does. The study's setting (cells / animal / people) is named.
- ☐ One inspectable fact (percent, lot, or method) appears before any adjective about skin.
- ☐ One number per sentence where a number matters.

## Gate C — Reseller integrity

- ☐ The COA is attributed to **the manufacturer's certificate**, never "our lab," "we tested," or an implied in-house test.
- ☐ Fulfillment is disclosed: the page/shipping policy says the order is fulfilled and shipped by a third party, with the lead time that implies.
- ☐ Seller of record is stated: GHK Labs sells; returns and support come back through us.
- ☐ The affiliate/endorsement disclosure slot is **not rendered** (we are a reseller, not an affiliate — [fulfillment-model.md](fulfillment-model.md), Decided). If the model ever changes, this flips to required.
- ☐ The supply partner is **not named** anywhere on the page, in metadata, image alt text, or structured data.

## Gate D — COA and lot

- ☐ The sellable variant points at a released lot, and that lot's COA is downloadable from the page.
- ☐ The lot code shown matches the COA; where the partner feed does not guarantee a per-order lot, the lot is labeled as representative.
- ☐ `documentSha256` recorded for the published COA, so a swapped file is detectable.
- ☐ A recalled lot is not attached to any buy button and carries the recall in its first line.
- ☐ A product with no released COA is not `availableForSale`.

## Gate E — Compliance and payments

- ☐ Responsible person (MoCRA) is settled and named on the label; the site's adverse-event contact routes to them. (Open in [fulfillment-model.md](fulfillment-model.md), Q8 — blocking for launch.)
- ☐ Facility registration and product listing are done or counsel has confirmed the plan.
- ☐ The package label (identity, net quantity, full INCI, warnings, responsible person) is signed by counsel. The site does not substitute for it.
- ☐ Any appearance claim that needs evidence ships only with a source counsel accepts. The launch claim set is short and boring.
- ☐ Powder / research-chemical SKUs are **out of this checkout** until the payment processor accepts them. The serum is an ordinary skincare SKU.
- ☐ Ads, email, SEO titles, and alt text are reviewed under the same claim class.

## Gate F — Pages and mechanics

- ☐ Required `/legal` pages are live and linked: shipping, returns, cosmetic-disclaimer, privacy, terms (and raw-material-notice if powder is live).
- ☐ Automated banned-word scan is clean over the page body, SEO fields, and alt text (allowing the negated "what it is not" boundary line and labeled science citations).
- ☐ All links resolve, including the COA download and the adverse-event contact.
- ☐ Mobile layout holds; no horizontal scroll; hero is LCP-friendly.
- ☐ Metadata (title, description) obeys the same claim class.

## Sign-off

A human confirms Gates A–F and records it. Counsel's separate claim-sheet sign-off is referenced here, not replaced by it.

```
Unit (page / SKU / lot): ____________________________
Lot (if applicable):     ____________________________
Reviewer:                ____________________________
Date:                    ____________________________
Counsel claim sheet ref: ____________________________
Gates A–F pass:          ☐ yes   ☐ blocked (reason: __________)
```

## Out of scope

Being the claim sheet itself, naming the legal entity or responsible person, and the payment underwriting decision. Those are counsel's and the operator's; this list only makes sure they were done before publish.
