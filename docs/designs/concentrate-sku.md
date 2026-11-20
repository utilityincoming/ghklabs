# GHK-Cu concentrate SKU

Status: draft

A fourth SKU for the storefront ([storefront.md](storefront.md)): a finished
cosmetic **concentrate** meant to be diluted into the user's own serum or cream,
distinct from the ready-to-use serum (P1) and the formulator powder (R1/R2). It
sits under the same reseller/dropship model ([fulfillment-model.md](fulfillment-model.md)),
the same voice ([brand-voice.md](brand-voice.md)), and the same authority order
and banned-word discipline as [storefront-copy.md](storefront-copy.md). Not legal
advice, not a signed claim sheet, not a committed price.

This document **captures a physical prototype already in hand**; it does not
invent the product. Facts below in `{braces}` are still unpublished (lot, purity,
method, pH, date, price, entity). Facts taken from the prototype label are stated
plainly and marked *(from label)*.

## 1. What the prototype is

A 5 mL amber-glass vial, copper flip-top orifice-reducer cap, front panel
"**Cu. CONCENTRATE · GHK-Cu 10%**", described on-label as "GHK-Cu (Copper
Tripeptide-1) concentrate for precision dosing and controlled copper intensity."

| Field | Value *(from label)* |
| --- | --- |
| Front name | GHK-Cu 10% Concentrate |
| Net contents | 5 mL (0.17 fl oz) |
| Strength | 10% Copper Tripeptide-1 (printed on label) |
| Format | Water-based concentrate, dropper/orifice-reducer vial |
| INCI | Water (Aqua), Propanediol, Copper Tripeptide-1, Glycerin, 1,2-Hexanediol, Phenoxyethanol, Ethylhexylglycerin, Lactic Acid, Sodium Hydroxide |
| Directions | "Do not use undiluted. Mix 1–2 drops into your serum or cream in the palm, blend, and apply immediately. If stored cold, allow to return to room temperature and invert gently." |
| Caution | "For external use only. Must be diluted before use. Avoid contact with eyes. Discontinue use if irritation occurs. Keep out of reach of children. Patch test before first use." |
| Period-after-opening | 6M |
| Lot (this unit) | GHK2601A |
| Distributor on label | Cuivelle LLC, 6501 Congress Ave, Boca Raton, FL 33487 · CopperPeptides.com |

Two label facts the storefront does not yet carry, and which this SKU forces into
the open: the **public entity is printed** (Cuivelle LLC / CopperPeptides.com),
which touches the unresolved display-name question ([storefront.md](storefront.md)
Q9 — "GHK Labs" is provisional); and shelf life is a **period-after-opening (6M)**,
not the `{date}`/expiry model the other pages assume. Both are deferred to the
separate "label truth-up" work (SITEMAP: this design does not resolve Q9 or the
PAO/expiry model; it only notes the collision).

## 2. The central tension (resolve before accept)

The concentrate is the product the catalog was shaped to avoid.

- `pricing-and-products.md` §6 P1 keeps the serum to **one strength, one size,
  one formula**, explicitly rejecting a "2%/3% variant" because it "invites the
  concentration-vs-blend confusion we're positioned to beat."
- A 10% dilute-it-yourself concentrate **is** blend-your-own, at a higher number
  than the hero serum, with the mixing decision handed to the customer.
- The prototype's own front copy ("precision **dosing**", "controlled copper
  intensity") uses the register `storefront-copy.md` bans on product pages
  (`dose`, `protocol`, and the implied potency-chasing the serum positioning
  disowns).

This is not a copy nit; it is a positioning decision. Three ways out, for the
principals to choose (Q1):

1. **Ship it, re-voiced.** Keep the SKU, drop the "dosing/intensity" frame.
   Market it as a *booster you add to products you already own*, not a
   strength you dial. Appearance-and-feel claims only, same as the serum. The
   10% is stated as the label fact, never as "more is better."
2. **Ship it, walled off.** Treat it like the powder: its own IA lane, a
   required "I understand this must be diluted" acknowledgment at cart, no
   placement beside the serum hero. Lower confusion risk, lower volume.
3. **Don't ship it.** The prototype stays a prototype; the catalog holds at
   three SKUs. Record the reason so it isn't re-proposed.

This design drafts option 1 as the recommended path (below) because the product
physically exists and the booster framing is honest and claim-safe; options 2–3
are live until a principal signs.

## 3. Catalog placement (if shipped)

New row, `productType: cosmetic`, claimClass `cosmetic_appearance` (same class as
the serum — it is a finished cosmetic for appearance and feel, not a raw
material).

| SKU | Type | Size | Price | Claim class | Channel |
| --- | --- | --- | --- | --- | --- |
| Copper Tripeptide-1 Concentrate | cosmetic | 5 mL | ${price} | `cosmetic_appearance` | Public |

Price is **not set here**. Modeled like the serum (dropship finished good, COGS =
confidential partner wholesale `W`), per `pricing-and-products.md` §4a — a small
high-strength add-on typically prices between a trial size and the hero; a band of
**$36–52** is the suggestion to resolve against `W` with ≥50% GM, gated on partner
wholesale and principal sign-off. It must not read as "cheaper per mL than the
serum therefore better value" — different product, different job.

## 4. Usage model (new to the catalog)

Neither the serum's "apply a few drops" nor the powder's "no mix rate." This is
**dilute-before-use**, and the directions are the label's, verbatim in substance:

- Do not use undiluted.
- Add 1–2 drops into a serum or cream in the palm, blend, apply immediately.
- If stored cold, return to room temperature and invert gently before use.
- Unbroken skin only. Patch-test the inner forearm first. If redness, burning,
  itching, or blistering shows up, rinse off and stop; see a physician if it
  persists. Keep out of reach of children.

The page states the dilution instruction as a **safety and use fact**, not as a
"customize your dose" selling point — that is the line that keeps option 1 inside
the banned-word rule.

## 5. Copy (draft, re-voiced per option 1)

Same five authority lines as every selling page in `storefront-copy.md` (who
sells, which lot, which number, which claim, one ask). `{braces}` stay blank until
the label and released COA supply them; `10%` is a product fact, printed on the
label, and is stated, not rounded.

**Eyebrow** — Copper Tripeptide-1 concentrate · 5 mL

**Name** — Copper Tripeptide-1 Concentrate

**Deck** — A small, water-based Copper Tripeptide-1 concentrate at 10%, the figure
on the label. Add a drop or two into the serum or cream you already use. Not a
stand-alone serum — it is made to be blended in.

- **What it is for** — adding Copper Tripeptide-1 to a routine you already like,
  for the look of fine lines and a moisturized feel.
- **What it is not** — a serum to use on its own, a drug, or a strength to chase.
  Do not use it undiluted.
- **How to use** — 1–2 drops blended into a serum or cream in the palm, then
  apply. Patch-test first.
- **On this lot** — Representative certificate: lot {lot}, released {date}.
  {method} identity. {purity}, copied from that file. [Download the manufacturer's
  COA.] The bottle that ships carries its own lot code.
- **The percent** — 10% is the figure printed on the label. This page repeats it.
  It does not round it, and it does not tell you more is better.
- **Shipping** — Sold by {display name}. We do not make it. Fulfilled and shipped
  by our fulfillment partner.

**Add to bag.**

**At a glance**

| | |
| --- | --- |
| For | Blending Copper Tripeptide-1 into products you already own. |
| Suited for | People already using a serum who want to add the peptide. Patch-test first. |
| Format | Water-based concentrate, 5 mL. Must be diluted. |
| Copper Tripeptide-1 | 10%, printed on the label. |

pH {range} renders once the formula sheet gives it — the INCI carries Lactic Acid
and Sodium Hydroxide, so pH is a knowable, not-yet-published number. "Water-based"
is confirmed from the INCI; no alcohol-free / oil-free / silicone-free claim
unless the partner INCI confirms each.

**Metadata** — Title: Copper Tripeptide-1 concentrate, 10%, lot on the page.
Description: A 5 mL water-based concentrate sold by {display name}. Blend 1–2 drops
into your serum or cream. For the look of fine lines and a moisturized feel.
Manufacturer certificate for the current released lot.

**Microcopy additions**

- **Dilution reminder** — Concentrate. Do not use undiluted — blend 1–2 drops into
  a serum or cream.
- **Cart note** — Cosmetic for appearance and feel, meant to be diluted. Not a drug.

## 6. Out of scope

Setting the price, writing the claim sheet, resolving the public display name
(Q9), and reconciling the PAO/expiry model across pages — the last two are the
"label truth-up" work tracked separately in SITEMAP. This design does not change
`/v1`, the serum, or the powder SKUs.

## 7. Open questions

1. **Positioning (§2).** Ship re-voiced (option 1, drafted), ship walled-off
   (option 2), or don't ship (option 3). Principal decision; blocks everything
   below.
2. **Price.** Band $36–52 suggested, gated on confidential partner wholesale `W`
   and ≥50% GM (`pricing-and-products.md` §4a), principal sign-off before publish.
3. **Is the concentrate dropshipped by the same partner** as the serum, and does
   its label percent/fill come from the partner sheet the same way? Confirm before
   the 10% and 5 mL are asserted as settled.
4. **Claim sheet.** A finished cosmetic asking for money needs counsel's signed
   claim sheet, same gate as the serum, before publish.

## 8. Acceptance

Not accepted. Draft for review. On acceptance, add the SKU row to
`storefront-copy.md` §Catalog and `pricing-and-products.md` §6, and update
`storefront.md`'s catalog count, in the same change.
