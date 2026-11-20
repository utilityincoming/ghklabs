# Label truth-up

Status: draft

A physical prototype now exists (captured in [concentrate-sku.md](concentrate-sku.md)),
and its label carries facts the storefront copy has been holding blank. This design
reconciles the **label against the accepted storefront docs** in two parts:

- **Part 2 — Identity & entity.** The label prints a distributor and a consumer
  domain. Decide what that resolves — and what it must not accidentally break.
- **Part 3 — Integrity pass.** Shelf-life, pH, and the line between the
  manufacturer's label words and our own claim discipline.

Not legal advice, not a signed claim sheet. This doc **proposes** edits to accepted
documents; it does not make them. Naming the legal entity and the responsible
person stay out of scope ([pre-publish-checklist.md](pre-publish-checklist.md),
Out of scope; [storefront.md](storefront.md) Q5) — this doc only identifies a
candidate and the decision that would confirm it.

## 0. What the prototype does and does not settle

The prototype is **one product — the concentrate** — not the serum or the powder.
So its lot (GHK2601A), its INCI, and its 10% figure are *that product's* facts and
do **not** fill the serum's or powder's `{lot}` / `{percent}` / `{purity}`. Each
SKU's braces still come from its own label and released COA (Gate A unchanged).

What the prototype settles is **brand-level**, not SKU-level:

| Label fact | What it can resolve | What it must not do |
| --- | --- | --- |
| "Distributed by Cuivelle LLC, 6501 Congress Ave, Boca Raton, FL 33487" | Nothing on our surfaces. Principal answer: this is the supply partner | Be published as `{entity}`, `{address}`, or a display name |
| "CopperPeptides.com" | Nothing on our surfaces. It rides with the partner's label | Be used as our brand or `{support_email}` domain |
| "AFTER OPENING: 6M" | A shelf-life model the storefront has no field for | Be confused with a release date or an expiry date |

## Part 2 — Identity & entity

### 2.1 The fork (principal + counsel decision)

A cosmetic label's "Distributed by" line names the party that markets and sells the
finished good — typically the brand owner / merchant of record. Paired with a
consumer domain (CopperPeptides.com), the natural reading is:

> **Reading A (recommended):** Cuivelle LLC is *our own* selling entity, and
> "Copper Peptides" / CopperPeptides.com is the public brand. This resolves Q9
> (display name), `{entity}` = Cuivelle LLC, `{address}` = 6501 Congress Ave,
> Boca Raton, FL 33487, and gives a `{support_email}` domain.

But the storefront's whole posture is *reseller of a finished cosmetic from a
**supply partner** who is never named* ([fulfillment-model.md](fulfillment-model.md):13,
Gate C). So the other reading is live and dangerous:

> **Reading B (the trap):** Cuivelle LLC is the **supply partner / manufacturer-
> distributor**, and the prototype is *their* product. Then naming Cuivelle or
> CopperPeptides.com on our storefront would **violate the unnamed-partner rule**,
> and our own brand/entity is something else not on this label.

These cannot both be acted on. Which is true is a fact the principals hold (who
owns Cuivelle, who owns the domain, whose product the prototype is), not something
this doc can infer from the label alone.

### 2.2 Resolution (Reading B)

**Decided 2026-11-19.** Cuivelle LLC is the unnamed supply partner. Reading A is
closed. The distributor line, the Boca Raton address, and CopperPeptides.com stay
off every ghklabs surface: pages, metadata, alt text, structured data, and
image files. Our `{entity}`, `{address}`, and `{support_email}` stay blank.
Q9 stays open: the public name is **TBD**. "ghklabs" remains the working name and
"GHK Labs" the provisional public spelling until a name is chosen. Do not treat
the provisional spelling as the trade name.

### 2.3 Product images while the name is open

The story still needs a picture. Use **generic product images**: a plain bottle,
dropper, or powder jar on a light editorial ground, in the calm skincare look
from [storefront.md](storefront.md). No wordmark, no partner name, no
CopperPeptides.com, no distributor address, no "precision dosing" line, and no
before/after. Alt text names the INCI and the format ("Copper Tripeptide-1
serum bottle"), not a brand. These images are stand-ins. They are replaced when
Q9 locks and a real pack shot exists that does not show the partner.

## Part 3 — Integrity pass

Findings from reading the label against the accepted docs. Each is a proposed
change to an accepted doc, listed precisely for review.

### 3.1 Shelf life has no home on the site (real gap)

The label carries **period-after-opening: 6M** and a lot code, with no printed
expiry. The storefront model has **no field for PAO**. Gate A sources `{date}` only
as the lot's *released* date (from the COA), and the copy says "released {date}" —
that is a release date, not shelf life. So the one durability fact a buyer acts on
(how long after opening to use it) has nowhere to appear.

**Proposed:**
- Add `{pao}` to Gate A, sourced from **the finished-product label**
  ([pre-publish-checklist.md](pre-publish-checklist.md) Gate A table).
- Add a spec row "Use within" / PAO to the serum and concentrate PDPs
  ([storefront-copy.md](storefront-copy.md) §Serum "At a glance" and the
  concentrate's), rendered only when the label states it.
- Add microcopy: **"After opening, use within {pao}."**
- Keep `{date}` meaning the COA release date, explicitly distinct from `{pao}`
  (one line in Gate A so they are never conflated).

### 3.2 pH is a real, available fact (confirm, no gap)

The concentrate INCI carries **Lactic Acid** and **Sodium Hydroxide** — a pH-adjusted
acid system — so pH is a genuine formula-sheet number, exactly what Gate A already
sources `{range}` from. No change; annotate that pH figures are available for any
acid-system SKU and are a formula-sheet fact, not an invented one. The existing
"pH {range} renders once the formula sheet says so" copy is correct as written.

### 3.3 Manufacturer label words are not our claims (reinforce the wall)

The prototype's **front label** uses "precision dosing" and "controlled copper
intensity" — register the storefront **bans** on product pages (`dose`, `protocol`,
potency-chasing; [storefront-copy.md](storefront-copy.md), Gate B). The package
label is the manufacturer's and is the controlling document for ingredients and
warnings ([storefront-copy.md](storefront-copy.md) cosmetic-disclaimer), but our
**marketing copy is ours** and must obey the banned list even when transcribing
from a bottle we resell.

**Proposed:** add one explicit line to Gates C and F
([pre-publish-checklist.md](pre-publish-checklist.md)):
- Gate C: "Marketing phrases printed on the manufacturer's front label (e.g.
  'precision dosing', 'controlled intensity', a bare '99%') are **not** carried
  onto our pages, metadata, or alt text; our claim class governs our copy."
- Gate F: the banned-word scan explicitly covers **text copied from supplier
  labels or listings**, not only text we authored.

This is already handled correctly for the concentrate in
[concentrate-sku.md](concentrate-sku.md) §2/§5 (it re-voices those words); the
checklist lines make the rule general so the next resold label gets the same
treatment.

## 4. Out of scope

Naming the legal entity or responsible person (counsel/principal; Q5, Gate E),
filling any specific SKU's `{lot}`/`{percent}`/`{purity}` (each needs its own
label + COA), and the concentrate's own positioning/price (that is
[concentrate-sku.md](concentrate-sku.md)). This doc proposes no change to `/v1`.

## 5. Open questions

1. **Q-ID1 is answered (§2.2).** Reading B. The public name (Q9) is still TBD.
2. **Does every SKU label carry a PAO**, and is 6M the house default or
   per-product? Confirms the `{pao}` field's source and whether it is one value or
   per-label (§3.1).

## 6. Acceptance

Not accepted. Draft for review. On acceptance:
- §3.1 and §3.3 are edits to [pre-publish-checklist.md](pre-publish-checklist.md)
  and [storefront-copy.md](storefront-copy.md), made in that accepting change.
- §2.2 is already decided: no accepted doc gains the partner's name, domain, or
  address. Q9 and `{entity}` stay open.
