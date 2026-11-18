---
name: ghklabs-elephant
description: Load ghklabs project context before Elephant-Goldfish brainstorming, drafting, or implementation. Use when designing or changing ghklabs.
---

# ghklabs Elephant

ghklabs is a greenfield repository. There is no application code yet. Treat `docs/designs/` as the source of truth for intended behavior.

## Context

- Design index: `docs/designs/SITEMAP.md`
- Designs: `docs/designs/`
- This skill: `_agents/skills/ghklabs_elephant/SKILL.md`

## Rules

- Do not invent product requirements. Capture them in a design doc first.
- New designs go in `docs/designs/` and get a row in `SITEMAP.md`.
- Implement from an accepted design. If the code and the design diverge, update one of them in the same change and say which.
