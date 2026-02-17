# Expanded Leaf Vocabulary & New Plant Variants

## Problem

The current favicon-farm has 7 leaf path definitions (pointed, wide, narrow, medium, round, slim, fat-pointed), all using basic symmetric cubic Bézier curves. We want to expand the leaf vocabulary with more realistic and varied shapes using advanced SVG path techniques (smooth curves, arcs, chained segments), and create new plant icon variants that showcase them.

## Approach

Add 6 new leaf shape definitions using techniques from MDN/CSS-Tricks SVG path documentation, improve existing shapes, and compose 5 new plant icon variants.

## New Leaf Shapes

| ID | Name | Technique | Box | Description |
|---|---|---|---|---|
| LEAF_HEART | Heart-shaped | C curves, inward-notched base | 30×40 | Philodendron — two lobes at base curving inward |
| LEAF_SERRATED | Serrated edge | Chained C/S segments | 40×50 | Oak/maple — zigzag edges via multiple small curves |
| LEAF_LOBED | Lobed/split | C curves with cutouts | 40×50 | Monstera — wide shape with indentations |
| LEAF_GRASS | Grass blade | Q curves, very tight | 30×40 | Bamboo — extremely narrow, slightly curved |
| LEAF_SUCCULENT | Succulent teardrop | A arcs for roundness | 30×40 | Fleshy — fat bottom tapering to blunt tip |
| LEAF_FAN | Fan-shaped | C curves fanning wide | 40×50 | Ginkgo/palm — narrow base spreading into fan |

## Existing Shape Improvements

- Refine LEAF_WIDE and LEAF_ROUND control points for more natural silhouettes
- Add slight asymmetry options (1-2 unit CP offset)

## New Plant Variants (5)

1. **Tropical Mix** — heart + lobed + wide
2. **Herb Garden** — serrated + pointed + grass
3. **Succulent Pot** — succulent + round + fat-pointed
4. **Fan Palm** — fan + grass + slim
5. **Wild Mix** — one of each new type, asymmetric

## Deliverables

- Update `gen_variants.py` with new leaf constants + 5 new plant variants
- Regenerate `variants_block.html`
- Create `leaf-showcase.html` showing all leaf shapes individually with labels
