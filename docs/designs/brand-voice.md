# Brand voice

Status: draft

This is the voice for the first public copy on the cosmetic storefront in [storefront.md](storefront.md). It does not change the webhook API or the delivery marketing page. It is not legal advice. Counsel still signs the claim sheet before a page ships.

Pages were read on 18 November 2026 (America/Chicago). Quotes below are short samples of public marketing language, used to name a voice, not to copy it.

## What the field sounds like

| Voice | Where it shows up | What it buys | What it costs |
| --- | --- | --- | --- |
| Evidence desk | [CopperPeptide.com](https://copperpeptide.com/) | Trust. Every item names its format and evidence class. A dish, a mouse, and a person are three different sentences. The explainer ([what is GHK-Cu](https://copperpeptide.com/what-is-ghk-cu/)) keeps the Maquart 1988 fibroblast line attached to “cultured cells,” and the [route piece](https://copperpeptide.com/why-route-and-quality-matter/) refuses to treat a cosmetic, a research material, and an injectable as the same object. | No product. Authority with nowhere to go. A reader who is ready to buy has to leave. |
| Protocol label | [koppā](https://koppabeauty.com/) | A narrow catalog, a printed percent (2% serum, 0.5% cream and balm), and an explicit split: cited studies are about the ingredient, not finished koppā products. Supplier COA called out as a minimum 99% purity spec. | “Collagen support” and “firmer-looking skin” sit in the same breath. The hacker reader still has to supply the evidence boundary. |
| Clinic engine | [NIOD CAIS3](https://niod.com/en-us/copper-amino-isolate-serum-3-11-cais3-serum-100368.html) | States 1% GHK-Cu plus 1% GHK, names the blue, gives a pH band, and tells you the order of the routine. | “Promotes natural collagen,” “visibly repaired” wrinkles, and week-count outcome lines. The jargon performs expertise and also steps outside a cosmetic appearance claim. |
| Plain bench | The Ordinary Multi-Peptide + Copper Peptides 1% | Short, cheap, clinical, no romance. The 1% is a blend name; GHK-Cu’s share is not the headline. | Authority is the house brand. The peptide itself is one line in a stack. |
| Founder asterisk | Skin Biology CP Serum (Dr. Loren Pickart), sold at reverseskinaging.com | A named biochemist, a how-to-apply, and a long paper trail. | “Restore youthful,” “proven,” and healing-adjacent lines with asterisks. The copper source on the classic CP Serum is often a hydrolyzed-protein complex, which a careful reader should not treat as “GHK-Cu, stated percent.” |
| Luxury recovery | Allies of Skin Copper Tripeptide & Ectoin | Texture, barrier, percent stack (copper tripeptide complex plus ectoin and a copper salt), clinical-sounding recovery stats. | “Recover,” “repair,” and damage language. The bottle is doing medical work in the adjectives. |

The open lane is the CopperPeptide.com habit of naming the evidence, worn by someone who actually sells a finished cosmetic and will show the lot. koppā is the closest neighbor on disclosure. NIOD is the neighbor on “I read the bottle.” Neither sounds like a person who will walk you through the paper and then hand you one serum.

## The voice

**Friendly professional biologic hacker salesman.**

One person. Four jobs, in this order, on every page that asks for money or belief:

1. **Friend.** Second person. Short sentences. Assume the reader is intelligent and new to this ingredient. No club passwords (“stack,” “pin,” “research chem,” “not for human consumption” with a wink).
2. **Professional.** INCI name on cosmetics (Copper Tripeptide-1). Common name GHK-Cu when identity is the point. Numbers only when they match the label or the linked COA. Method named beside a purity figure.
3. **Biologic hacker.** The hack is reading, not dosing. Say what the molecule is (three amino acids plus copper), what a cited study actually measured, and which setting it was in. Keep route, formula, and product as separate objects, the way CopperPeptide.com does.
4. **Salesman.** One ask. On consumer pages the ask is the finished serum. Quality and science pages earn the ask; they do not bury it. Raw powder stays a separate, quieter voice: identity, pack, storage, and the formulator notice. It does not pitch a personal protocol.

Working name remains ghklabs until question 9 in the storefront design is closed.

### Sentence shapes

- Name the thing, then the boundary, then the next step.
- Prefer “the look of” and “feels” for the serum. Prefer “the paper measured” for the science page.
- One number per sentence when a number matters (size, stated percent, lot purity, sample size).
- A citation is a clause, not a badge that upgrades a claim. “Maquart and colleagues, 1988, reported more collagen synthesis in cultured fibroblasts. That is a cell result. It is not what this serum claims.”

### Words we use

look of texture, softer-looking fine lines, moisturized feel, comfortable skin, Copper Tripeptide-1, GHK-Cu, stated concentration, lot, HPLC (or the method on that COA), identity, purity, intact skin, cosmetic, formulator, vehicle, sample size, cultured cells, animal model, human study.

### Words we do not use on product pages

heal, repair, recover, reverse, stimulate collagen, boost collagen, DNA, gene, inflammation, wound, hair growth, clinical proof, proven, doctor-recommended, biohack (as a verb on the serum), research use only, injectable, subcutaneous, reconstitute, dose, protocol (as a personal treatment plan).

Science pages may name those words when explaining why a paper does not transfer. They do not then attach them to the SKU.

### Evidence class, said in our voice

Steal the CopperPeptide.com badge, drop the desk tone:

| Class | How we say it |
| --- | --- |
| Laboratory | “Measured in cells. Useful for a mechanism question. Not a result on a face.” |
| Animal | “Measured in an animal model. It does not transfer to people on its own.” |
| Human | “Measured in people. We say how many, what was compared, and whether anyone was blinded.” |
| Review | “A summary of earlier papers. Only as strong as those papers.” |
| This bottle | “What we claim for the serum is appearance and feel. The papers are linked beside it, labeled as papers.” |

If we have not run a finished-product test, we say that in one line, the way koppā separates ingredient research from tests of its own products.

## Initial copy

Placeholders in braces are facts we do not invent: concentration, lot, purity, method, price, size. They come from the label and the released COA before publish. Claim class for all of these samples is `cosmetic_appearance` except the raw aside, which is `identity_only`.

### Home hero

Eyebrow: Copper Tripeptide-1, for skin.

Headline: A blue serum with the peptide named, the percent printed, and the lot on the glass.

Deck: GHK-Cu is a small copper complex people have studied for decades. Those studies are not the same thing as a bottle. Ours is a cosmetic for the look and feel of skin: softer-looking texture, a moisturized feel, nothing that pretends to treat a condition.

Primary: Shop the serum.

Secondary: Read what the papers actually measured.

Proof line under the fold: Current lot {lot}. Identity by {method}. Purity {purity} on the COA. Same file you can download.

### Serum PDP, above the fold

{Size}. {Price}. Copper Tripeptide-1 at {percent} in a finished serum.

What it is for: the look of fine lines and the feel of moisturized skin. Apply to clean, intact skin, morning or night, then a moisturizer if you want one.

What it is not: a drug, a wound treatment, or a lesson in how to inject anything.

On this lot: {lot}, released {date}. {method} identity. {purity}. [Download the COA.]

### Science pillar, opening

GHK-Cu is glycyl-L-histidyl-L-lysine bound to copper. Three amino acids. On a cosmetic label the INCI name is Copper Tripeptide-1.

A famous early paper (Maquart et al., 1988) reported more collagen synthesis in cultured fibroblasts. That sentence is about cells in a dish. A later human study people cite is small, set after laser resurfacing, and mixed: satisfaction moved; several blinded measurements did not (see our note on that paper). We keep those results in their own rooms.

If you want the serum, it is here for appearance and feel. If you want the powder, it is a cosmetic ingredient for formulators, sold separately, with its own notice.

### Raw PDP, the salesman who steps back

{Net weight} of GHK-Cu for formulation. Not a finished cosmetic. Not sold to swallow or inject. No mixing instructions for the body.

Identity {method}. Purity {purity}. Lot {lot}. Store as the COA says. You confirm the formulator notice before checkout.

### Quality

We publish the lot that is on the unit you can buy. If a number is not on that COA, it is not on this page. A recalled lot stays listed, with the recall in the first line, and it is not attached to a buy button.

## How a page earns trust before it asks

Use this order on home, PDP, and the pillar. Skip a step only when the page is not selling (legal, contact).

1. Say the category in the first screen: cosmetic serum, or formulator ingredient.
2. Give one inspectable fact (stated percent, lot, or method) before any adjective about skin.
3. Put the evidence boundary in the same view as any mention of research. One sentence is enough on a PDP. The pillar can go long.
4. Make the COA a link with the lot code visible, not a badge that says “lab tested.”
5. Ask once, with a verb that matches the product: “Shop the serum” or “Read the powder spec.” Do not offer both as equal heroes.

## Two page drafts

Read 18 November 2026 (America/Chicago). These follow the shape of two sellers. They do not copy their claims, prices, or formulas.

**Finished serum: The Ordinary.** Their Multi-Peptide + Copper Peptides 1% page is a white-label pattern we can use: name the technology in the title, one overview line, a targets / suited-for / format row, free-from chips, a short about, then how to use (a few drops, morning and evening, unbroken skin, patch test, stop if irritation, keep from children). We do not import their stack of peptide trademarks, their “1%” blend name, or lines about elasticity, regeneration, or collagen. A white-label serum has one hero ingredient until a real INCI list exists.

**Raw powder: Peptaura.** Their GHK-Cu pages are a spec table: pack, purity, batch, identity, mass, and a COA you can open. That table is the pattern. We do not import the rest of the offer. Peptaura lists lyophilized vials in milligrams, price per milligram, and research-use language on the certificates it hosts. This catalog sells a cosmetic ingredient to formulators, by net weight, with no vial, no per-milligram price, and no research-only wink. Storefront default stands.

Braces are unpublished facts. Leave them blank on the live page until the label or the released COA has the number.

### Finished serum

**Copper Tripeptide-1 serum**

A water-based serum for the look of fine lines and the feel of moisturized skin.

| | |
| --- | --- |
| Targets | The look of fine lines. Dry-feeling skin. |
| Suited for | Skin that already uses a serum. Patch-test first. |
| Format | Serum, water-based. {size}. |
| Copper Tripeptide-1 | {percent}, printed on the label. |

pH {range}, once the formula sheet says so. Alcohol-free, oil-free, and silicone-free only if the INCI confirms each one.

**About**

Copper Tripeptide-1 is the label name for GHK-Cu, a copper complex of three amino acids. This bottle is that ingredient in a finished cosmetic, at {percent}, in {size}. It is meant to soften the look of texture and leave skin feeling moisturized. Studies of the molecule are a different page. They are not a test of this serum.

**How to use**

- After cleansing, apply a few drops to the face, morning and evening.
- Follow with a moisturizer if you want one.
- Use on unbroken skin only.
- Patch-test on the inner forearm before the first full use. If redness, burning, itching, or blistering shows up, rinse off and stop.
- If irritation occurs, rinse off, stop, and talk to a physician.
- Keep out of reach of children.

**Specs**

{size}. {price}. Lot {lot}, released {date}. Identity by {method}. Full INCI on the package, in descending order. Current COA on the lot line.

Add to bag.

### Raw GHK-Cu

**GHK-Cu, cosmetic ingredient**

{net weight} of GHK-Cu powder for people who formulate. Not a finished serum. Not a vial.

| | |
| --- | --- |
| Net weight | {net weight} |
| Identity | GHK-Cu, method {method} |
| Purity | {purity}, same figure as the COA |
| Lot | {lot} |
| COA | Download. A replaced file changes the recorded hash. |

**About**

You are buying the ingredient, identified and weighed. The number on this page matches the certificate for this lot. Storage is whatever that certificate says. There is no mix rate for skin, and no instruction for swallowing or injecting. Confirm the formulator notice before checkout.

**Specs**

{net weight}. {price}. Lot {lot}. {method}. {purity}.

Read the powder spec.

## Out of scope

Naming a legal entity, writing the claim sheet, inventing concentrations, and translating this voice into ads. Ads later use the same banned-word list. This draft does not accept the storefront; question 9 there (brand name and visual reference) can point at this voice without closing the other open questions.
