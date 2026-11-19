# Storefront copy

Status: accepted

Publish-ready copy for the cosmetic copper-peptide storefront in [storefront.md](storefront.md), written in the voice set in [brand-voice.md](brand-voice.md): *friendly professional biologic hacker salesman* — friend, professional, biologic hacker, salesman, in that order, on every page that asks for money or belief.

This is copy, not legal advice. Counsel signs the claim sheet before a page ships. Claim class is `cosmetic_appearance` for the serum and home pages, and `identity_only` for the raw powder.

GHK Labs resells a finished cosmetic and ships it by dropship (see [fulfillment-model.md](fulfillment-model.md)). So the copy attributes the COA to the manufacturer, never to an in-house lab, and is truthful about third-party fulfillment. The supply partner is not named on any consumer surface.

**Public name is provisional.** This copy uses "GHK Labs" as the seller's display name; internal design docs use the repo slug "ghklabs." The brand name is still open (storefront.md Q9), and `{entity}` / `{address}` are unresolved, so "GHK Labs" is the working public spelling, not a settled trade name. Lock one form before publish.

`{braces}` are facts we do not invent — concentration, lot, purity, method, price, size, pH, date. They come from the label and the released COA before publish, and stay blank on the live page until that number exists. Banned on product pages: *heal, repair, recover, reverse, stimulate/boost collagen, DNA, gene, inflammation, wound, hair growth, clinical proof, proven, doctor-recommended, biohack (verb), research use only, injectable, subcutaneous, reconstitute, dose, protocol.* Science pages may name those words to explain why a paper does not transfer; they do not then attach them to the SKU.

**Prices are commercial decisions**, resolved here from the suggested catalog in [pricing-and-products.md](pricing-and-products.md). They are draft suggestions, gated on the confidential partner wholesale and confirmed raw COGS, and the principals sign off before publish.

**Size is split.** The **powder** pack weights (1 g, 5 g) are genuinely commercial — GHK Labs buys bulk and repacks — so they are resolved here. The **serum's 30 mL is a target, not a settled fact**: under dropship the finished bottle's net contents is the partner's label figure, so it is confirmed against the partner's actual fill before publish, the same way `{percent}` is shown as "target 1%." Treat every "30 mL" below as "target 30 mL, pending partner fill." The remaining label/COA braces (`{percent}`, `{lot}`, `{purity}`, `{method}`, `{pH}`, `{date}`, per-pack `{net weight}`) stay blank until the label and released COA supply them.

---

## Catalog

Three SKUs: one serum, two powder sizes. Nothing else ships at launch.

| SKU | Type | Size | Price | Claim class | Channel |
| --- | --- | --- | --- | --- | --- |
| Copper Tripeptide-1 Serum | cosmetic | 30 mL | $48 | `cosmetic_appearance` | Public |
| GHK-Cu Powder (entry) | raw_material | 1 g | $34 | `identity_only` | Public |
| GHK-Cu Powder (formulator) | raw_material | 5 g | $119 | `identity_only` | Public |

Bulk powder (100 g / 1 kg) is a later B2B/quote tier, not public checkout ([pricing-and-products.md](pricing-and-products.md) §2, §6).

---

## Home

The first screen sells the serum. It names the ingredient, the look and feel, and the lot. It does not open on what the bottle refuses to be, and the secondary link goes to the lot, not to a paper. Research stays on the science page.

**Eyebrow** — Copper Tripeptide-1 serum

**Headline** — Softer-looking skin, with the peptide named and the lot on the glass.

**Deck** — GHK-Cu in a finished serum. The percent is printed on the label. A few drops, morning and evening.

**Primary button** — Shop the serum.

**Secondary link** — See this lot.

**Proof line** — Lot {lot}. Identity by {method}. Purity {purity}. The manufacturer's certificate for that lot — yours to download.

**Three lines**

1. **Named.** Copper Tripeptide-1 at {percent}, the INCI name for GHK-Cu.
2. **For the look.** Fine lines and a softer-looking texture.
3. **For the feel.** Skin that feels moisturized.

**Quiet line, below the fold** — A cosmetic for the look and feel of skin. Notes on the published research sit on the science page.

---

## Serum product page (PDP)

### Above the fold

**Copper Tripeptide-1 serum**

30 mL. $48. Copper Tripeptide-1 at {percent} in a finished, water-based serum.

- **What it is for** — the look of fine lines and the feel of moisturized skin. Apply to clean, intact skin, morning or night.
- **What it is not** — a drug, a wound treatment, or a lesson in how to inject anything.
- **On this lot** — {lot}, released {date}. {method} identity. {purity}. [Download the manufacturer's COA.]
- **Shipping** — Sold by GHK Labs, fulfilled and shipped by our fulfillment partner. Lead time {lead_time}. Returns through us.

**Add to bag.**

### At a glance

| | |
| --- | --- |
| Targets | The look of fine lines. Dry-feeling skin. |
| Suited for | Skin that already uses a serum. Patch-test first. |
| Format | Serum, water-based. 30 mL. |
| Copper Tripeptide-1 | {percent}, printed on the label. |

pH {range}, once the formula sheet says so. "Water-based," alcohol-free, oil-free, and silicone-free each render only if the partner's INCI / formula sheet confirms it — the formula is the partner's, so none of these is asserted unconfirmed.

### About

Copper Tripeptide-1 is the label name for GHK-Cu, a copper complex of three amino acids. This bottle is that ingredient in a finished cosmetic, at {percent}, in 30 mL. It is meant to soften the look of texture and leave skin feeling moisturized. Studies of the molecule are a different page. They are not a test of this serum.

### How to use

- After cleansing, apply a few drops to the face, morning and evening.
- Follow with a moisturizer if you want one.
- Use on unbroken skin only.
- Patch-test on the inner forearm before the first full use. If redness, burning, itching, or blistering shows up, rinse off and stop.
- If irritation occurs, rinse off, stop, and talk to a physician.
- Keep out of reach of children.

### Specs

30 mL. $48. Lot {lot}, released {date}. Identity by {method}. Full INCI on the package, in descending order. Current COA on the lot line.

---

## Science pillar

**Eyebrow** — What the papers measured, in their own rooms.

**Opening** — GHK-Cu is glycyl-L-histidyl-L-lysine bound to copper. Three amino acids. On a cosmetic label the INCI name is Copper Tripeptide-1.

A famous early paper (Maquart et al., 1988) reported more collagen synthesis in cultured fibroblasts. That sentence is about cells in a dish. A later human study people cite is small, set after laser resurfacing, and mixed: satisfaction moved; several blinded measurements did not. We keep those results in their own rooms.

If you want the serum, it is here for appearance and feel. If you want the powder, it is a cosmetic ingredient for formulators, sold separately, with its own notice.

**Evidence classes, said in our voice**

| Class | How we read it |
| --- | --- |
| Laboratory | Measured in cells. Useful for a mechanism question. Not a result on a face. |
| Animal | Measured in an animal model. It does not transfer to people on its own. |
| Human | Measured in people. We say how many, what was compared, and whether anyone was blinded. |
| Review | A summary of earlier papers. Only as strong as those papers. |
| This bottle | What we claim for the serum is appearance and feel. The papers are linked beside it, labeled as papers. |

We have not run a finished-product test on this serum. When we do, it gets its own line here, separate from the ingredient research above.

**Close** — Shop the serum. (Formulating instead? The GHK-Cu powder has its own spec page, sold separately.)

---

## Raw GHK-Cu powder (formulator)

*The salesman who steps back. Identity, pack, storage, notice — no personal protocol, no pitch.*

**GHK-Cu, cosmetic ingredient**

GHK-Cu powder for people who formulate, in two pack weights. Not a finished serum. Not a vial.

**Pack weight**

| Size | Price | Per gram |
| --- | --- | --- |
| 1 g | $34 | $34.00 |
| 5 g | $119 | $23.80 |

One product, one formula, priced by pack weight only — the 5 g is the formulator buy, cheaper per gram. Larger volume (100 g, 1 kg) is a B2B/quote request, not a cart item.

| | |
| --- | --- |
| Net weight | {net weight} (per selected pack) |
| Identity | GHK-Cu, method {method} |
| Purity | {purity}, same figure as the COA |
| Lot | {lot} |
| COA | Download. A replaced file changes the recorded hash. |

### About

You are buying the ingredient, identified and weighed. The number on this page matches the certificate for this lot. Storage is whatever that certificate says. There is no mix rate for skin, and no instruction for swallowing or injecting. Confirm the formulator notice before checkout.

### Specs

1 g — $34 · 5 g — $119. Net weight {net weight} for the selected pack. Lot {lot}. {method}. {purity}.

**Read the powder spec.**

---

## Quality

Each released lot carries the manufacturer's certificate. If a number is not on that COA, it is not on this page. The method sits beside every purity figure. A recalled lot stays listed, with the recall in the first line, and it is not attached to a buy button.

## Shipping & who you buy from

GHK Labs is the seller. Your order is fulfilled and shipped by our fulfillment partner, so delivery runs on their lead time, not a same-day promise. Returns and support go through us, at {support_email}. {return_window} to start a return.

---

## Microcopy

- **COA link** — Download the manufacturer's COA · lot {lot}
- **Patch-test reminder** — New to copper peptides? Patch-test the inner forearm first.
- **Out-of-stock** — This lot sold through. The next lot ships with its own COA.
- **Recall banner** — Recalled: lot {lot}. Do not use. [What happened.]
- **Cart note** — Cosmetic for appearance and feel. Not a drug.
- **Ships-from note** — Sold by GHK Labs. Fulfilled and shipped by our fulfillment partner.
- **Disclosure slot (reserved)** — Unused under the decided reseller / merchant-of-record posture ([fulfillment-model.md](fulfillment-model.md), Decided). Kept in the component only so an affiliate surface, if one ever exists, has a generic, partner-free line ("We may earn a commission on this purchase"). Not rendered on the store.
- **Formulator checkbox** — I am buying GHK-Cu as a cosmetic ingredient to formulate. It is not a finished product to swallow, inject, or apply as sold.

## Policy pages

Draft copy for the `/legal` routes, in the brand voice. Not legal advice; counsel confirms terms, the entity, the responsible person, and the windows before publish. `{braces}` are facts to fill from the operating setup (entity, address, support channel, lead time, return window) and the supply partner's actual RMA terms. The supply partner is not named.

### Shipping — `/legal/shipping`

**Who you buy from.** GHK Labs is the seller of record. {entity}, {address}.

**How it ships.** Your order is fulfilled and shipped by our fulfillment partner. We hand the order to them after payment clears; they pick, pack, and ship it.

**When.** Orders usually leave within {handling_time} and arrive within {lead_time} to US addresses. These are the partner's timelines, not a same-day promise. We ship to the United States only for now.

**Tracking.** You get a tracking link when the order ships. If a parcel stalls or arrives damaged, write {support_email} and we sort it out.

**What's on the box.** The parcel may show the fulfillment partner as the shipper. The product, lot, and the manufacturer's COA are unchanged by who hands it to the carrier.

### Returns — `/legal/returns`

**The window.** You have {return_window} from delivery to start a return. Email {support_email} first; we send the return instructions and the address.

**Condition.** For a hygiene product, we can accept {return_condition} (for example, unopened, with the lot code intact). We say exactly what qualifies before you ship anything back.

**Refunds.** Once the return reaches our fulfillment partner and checks out, we refund the item price to the original payment method within {refund_time}. Original shipping is {shipping_refund_rule}.

**A problem with the product.** If an item arrives wrong, damaged, or past its date, that is on us — email {support_email} with the lot code and a photo, and we replace or refund it without the return dance.

**Adverse reaction.** If the serum irritates your skin, rinse off and stop. For a reaction you want on record, use the adverse-event contact on [Contact]; that is separate from a return and it reaches the responsible person for the product.

### Cosmetic disclaimer — `/legal/cosmetic-disclaimer`

This is a cosmetic, sold for the look and feel of skin: softer-looking texture and a moisturized feel. It is not a drug. It is not meant to treat, prevent, or change any disease or condition, and nothing on this site should be read that way.

Studies of GHK-Cu as a molecule are linked on the science pages, labeled as studies. They describe what a paper measured, in cells, in animals, or in people. They are not a test of this serum and not a claim about what it does on your skin.

Use on clean, intact skin. Patch-test first. If redness, burning, itching, or blistering shows up, rinse off and stop, and talk to a physician if it persists. Keep out of reach of children. The package label is the controlling document for ingredients, warnings, and the responsible person; this page does not replace it.

## Out of scope

Naming a legal entity, writing the claim sheet, inventing concentrations, and translating this voice into ads. Ads later reuse the same banned-word list. This copy does not accept the storefront; it fills the voice that storefront question 9 points at.

## Acceptance

**Accepted 2026-11-19**, alongside accepted [storefront.md](storefront.md), [fulfillment-model.md](fulfillment-model.md), and [brand-voice.md](brand-voice.md). The copy is publish-ready: prices resolve from the accepted [pricing-and-products.md](pricing-and-products.md) and are draft suggestions gated on the confidential partner wholesale and confirmed raw COGS, with principal sign-off before publish. Every remaining `{brace}` (`{percent}`, `{lot}`, `{purity}`, `{method}`, `{pH}`, `{date}`, per-pack `{net weight}`, serum net contents, `{entity}`, `{address}`, `{support_email}`, and the partner's `{lead_time}` / `{return_window}` / RMA figures) is a **launch gate, not a design gate**, and stays blank on the live page until its source document exists — enforced by [pre-publish-checklist.md](pre-publish-checklist.md) Gate A. Public display name "GHK Labs" is provisional pending [storefront.md](storefront.md) Q9.
