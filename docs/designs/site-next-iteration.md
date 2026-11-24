# Next iteration for the ghklabs site

Status: draft for review

Production baseline: image rollout in commit `587199e`. The homepage and Science page serve all six WebP assets. The homepage captions identify the product render and specimen certificate as examples.

## 1. Make the verification path explicit

The homepage link “See this lot” currently leads to a specimen, while the Quality page says no lot has been released. Rename the link to “How lot verification works” and add a direct link to Quality beside the specimen. State the current release status near the image so readers can tell at a glance whether a file is available. Keep the specimen caption visible.

**Ready when** a first-time visitor can reach Quality from the homepage and tell whether there is a manufacturer PDF to download. When the first lot is released, follow the existing [publication checklist](pre-publish-checklist.md) for the lot code, manufacturer source, file hash, and wording about which lot ships.

## 2. Restore primary navigation on narrow screens

The stylesheet hides the primary navigation at 800 px and below. The footer still links to Science, Quality, and Powder, but readers have to scroll past the page content to find those links. Try a wrapping row or a keyboard-accessible menu. Keep the main serum link easy to find.

**Ready when** the primary destinations are reachable near the top at 320 and 390 px, keyboard focus is visible, and the header fits without horizontal scrolling.

## 3. Measure the hero before generating more variants

All six images are 1536 × 1024. The hero is about 45 KiB and already uses high fetch priority. Capture an uncached mobile Lighthouse run with network throttling, record the LCP element and timing, and compare against a 768 px WebP candidate with matching framing. Use responsive sources if the saved transfer improves the result while keeping the bottle legible. Check below-the-fold images for lazy loading and confirm that image boxes reserve space before decoding.

**Ready when** both runs are documented and a 390 px screenshot confirms the chosen hero framing and readable caption.

## 4. Replace the render when final packaging is available

Use a photo of the supplied bottle once the public packaging is approved. Match the current cream surface, diffuse lighting, and framing so the rest of the image set can still sit beside it. Add an in-use photo only if it clarifies how to dispense the serum. Before using the molecular illustration in additional placements, have its copper coordination and stereochemistry checked against a cited structure reference.

**Ready when** the product photo matches the approved bottle and label, and any accompanying certificate comes from the manufacturer for the lot named on the page.

Start with the verification wording and mobile navigation. Both can be previewed together without waiting for packaging or a released lot.
