# Fulfillment model: dropship reseller

Status: draft

This document records how ghklabs sources and ships the goods it sells, and what that reality forces to be true in copy, data, and compliance. It is the groundwork layer under [storefront.md](storefront.md); where the two disagree, the difference is called out here and `storefront.md` is updated when this is accepted. It is not legal advice. Counsel and a payments underwriter review the reseller posture before launch.

## Confidentiality

ghklabs resells a finished cosmetic sourced from a **supply partner** and ships it by dropship. The commercial relationship between the parties is private. This repository does not name the partner, does not quote the contract, and does not record margins or terms. "Supply partner" is the only handle used here. The partner's identity and the agreement live with the principals, outside version control.

This privacy covers the *terms*. It does not remove any truthful statement the buyer is owed (who they are buying from, that it ships from a third party, what testing stands behind the lot). Those are separate obligations handled below.

## Role

ghklabs is the **merchant of record**. The supply partner manufactures, holds inventory, tests, and ships.

| Responsibility | Owner |
| --- | --- |
| Storefront, brand, copy, claim review | ghklabs |
| Taking payment, tax, invoices, chargebacks | ghklabs (via Shopify checkout) |
| Every claim on the site and in ads | ghklabs — liability does not transfer to the supplier because they made the goods |
| Customer support, returns intake, refunds | ghklabs |
| Cosmetic adverse-event door (MoCRA contact) | ghklabs as the seller; routed to the responsible person on the label |
| Product formula, manufacturing, batch testing, COA issuance | Supply partner |
| Inventory, pick/pack, carrier, shipping label | Supply partner (dropship) |
| Lot code on the unit, COA document | Supply partner; ghklabs surfaces it, does not author it |

The one line that governs the copy: **we sell and stand behind the sale; we do not make or test the goods.** Copy may not imply in-house manufacturing or in-house testing.

## What changes versus storefront.md

`storefront.md` assumes ghklabs owns lots and COAs, runs a `lot.released` pipeline, and treats Shopify as the inventory system of record. Under dropship:

1. **COA provenance.** The certificate is issued by the supply partner, for their batch. ghklabs republishes or deep-links it and attributes it as the supplier's/manufacturer's certificate — never as "our lab," "the file we hold," or "we tested." The `documentSha256` still lets us detect a swapped file for a given lot; it is integrity, not authorship.
2. **Inventory.** ghklabs holds no stock. `availableForSale` tracks the partner's availability (feed, API, or a manual cadence in the MVP), not a warehouse we run. No claim that a specific physical unit was inspected by us.
3. **Lot binding.** The sellable variant still points at a current released lot, but "the lot on the unit you receive" is set by the partner at ship time and may not be knowable pre-purchase. Copy promises *a* released COA for the product, not a pre-selected lot number per order, unless the feed guarantees it.
4. **Fulfillment + shipping.** Orders are handed to the partner for dropship. Shipping lead time, carrier, and origin are theirs. Buyers are told the order ships from a third-party fulfiller. Blind/unbranded shipping (so the partner is not disclosed on the parcel) is a requirement to confirm with the partner, not an assumption.
5. **Returns.** Returns route through ghklabs support to a return address the partner supplies. The storefront's returns policy reflects the partner's RMA window, not an invented one.
6. **`lot.released` events.** Still future. The trigger is the partner publishing a lot; ghklabs ingests, it does not originate. MVP can surface the current COA manually.
7. **Pricing.** Whether retail price is ours to set or constrained by the partner is an open question (below). The data model does not assume either.

The tech stack recommendation in `storefront.md` (Next.js + Shopify headless + hosted checkout) still holds — dropship does not need Medusa or custom commerce. Shopify stays the commerce system of record for *money and orders*; the product/inventory truth is mirrored from the partner.

## Data model deltas

Against the Lot and COA table in `storefront.md`:

| Field | Change under dropship |
| --- | --- |
| `lotCode` | Supplier-assigned. Shown when the feed provides the lot that will ship; otherwise the current released lot is shown as representative, labeled as such. |
| `documentUrl` | Supplier-issued COA. Re-hosted in Shopify Files or deep-linked (open question). Labeled "certificate issued by the manufacturer." |
| `documentSha256` | Kept, for swap detection. Does not imply authorship. |
| `source` (new) | `supplier`. Marks that identity/purity figures originate from the partner's certificate, not an in-house test. |
| `availabilitySource` (new) | `partner_feed` \| `manual`. Drives `availableForSale` since we hold no inventory. |

A new order-side concern: a `fulfillment_partner_order_ref` on the order once an integration exists, so a dropshipped shipment maps back to our order. Out of MVP if hand-off is manual.

## Consumer-facing truthfulness (what the buyer is owed)

Private terms, yes. But the page must not mislead:

- **Who sells.** ghklabs is the seller; say so in terms and checkout.
- **Third-party fulfillment.** Shipping/returns pages state the order is fulfilled and shipped by a third party, with the lead time that implies. Naming the partner is not required; hiding that it ships from someone else is not an option if it affects delivery, returns, or the parcel's origin.
- **Testing provenance.** The COA is presented as the manufacturer's certificate for the batch. No "our lab," no implied in-house testing.
- **Material connection / affiliate.** If ghklabs is a plain reseller (buys at wholesale, sells at retail, is the merchant of record), this is ordinary retail and needs no endorsement disclosure. If any surface instead runs affiliate links or endorsements where ghklabs earns a commission on someone else's checkout, FTC 16 CFR Part 255 requires a clear, conspicuous disclosure there. The model decides which: **pick one posture and keep the copy consistent with it.** A generic disclosure slot is reserved near the buy path so adding it later is a config change, not a rebuild.

None of the above names the partner. All of it is compatible with a private agreement.

## Compliance deltas

- **Claims are still ours.** As merchant of record we are responsible for every appearance claim, regardless of who manufactured the serum. The claim-class discipline in `storefront.md` and the banned-word list in [brand-voice.md](brand-voice.md) apply unchanged.
- **MoCRA.** The responsible person is a function of who manufactures and distributes. Confirm whether ghklabs, the partner, or both carry responsible-person duties for the finished cosmetic, and whose name and address go on the label and the adverse-event contact. Blocking for launch copy; not for this draft.
- **Payment risk.** Unchanged and favorable: a Copper Tripeptide-1 cosmetic serum is an ordinary skincare SKU. Dropship does not add processor risk as long as the powder/research-chemical SKUs stay out of this checkout. Keep the catalog split.
- **Substantiation.** If the partner supplies claim substantiation, counsel still accepts it on our behalf before it ships as our claim.

## Phasing

1. **Partner site launches first** (the partner's own storefront). ghklabs builds nothing consumer-facing on top of it yet.
2. **ghklabs offering + architecture thereafter.** Then: confirm the sourcing posture (reseller vs. affiliate), the fulfillment hand-off (feed/API vs. manual), COA hosting, and MoCRA responsibility; then implement the storefront per `storefront.md` as amended by this document.

## Open questions

1. **Posture.** Reseller/merchant-of-record (default, per this doc) or affiliate earning commission on the partner's checkout? Changes disclosure and who takes payment. Default: reseller.
2. **Hand-off.** Partner feed/API for inventory and order routing, or manual in the MVP? Default: manual MVP, feed in phase 2.
3. **COA hosting.** Re-host the supplier's COA PDFs in our Shopify Files (stable URL, our integrity hash) or deep-link to the partner (reveals origin)? Default: re-host, attributed to the manufacturer.
4. **Lot visibility.** Does the feed tell us the exact lot that will ship per order, or only the current released lot? Default: show current released lot as representative until the feed guarantees per-order lots.
5. **Blind shipping.** Will the partner ship unbranded so the parcel does not disclose them? Confirm. Default: assume not guaranteed; shipping copy says "ships from our fulfillment partner" without a name.
6. **Pricing control.** Is retail price ours to set, or MAP/fixed by the partner? Default: unknown, do not assume.
7. **Returns window.** The partner's RMA terms set our returns policy. Unknown. Blocking for the returns page, not this draft.
8. **Responsible person (MoCRA).** Who is it for the finished cosmetic? Unknown. Blocking for label and launch copy.

## Acceptance

Accept after questions 1–3 and 8 are answered or defaults explicitly kept. On acceptance, fold the COA-provenance, inventory, and fulfillment deltas into `storefront.md` and update the data model there.
