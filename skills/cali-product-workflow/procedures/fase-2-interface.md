## Phase 2: Interface Brainstorming

Read `references/interface/archetypes.md` (5 archetypes) and `references/interface/output-format.md`.

**Generate each proposal individually in parallel** — launch 5 independent workers:

1. **Proposal A** — Archetype A with ASCII wireframes, breadboarding, trade-offs
2. **Proposal B** — Archetype B, full format
3. **Proposal C** — Archetype C, full format
4. **Proposal D** — Archetype D, full format
5. **Proposal E + Hybrid** — Archetype E plus hybrid recommendation combining strong elements

Each worker should be independent with no cross-contamination.
Combine outputs into `.cali-product-workflow/{YYYY-MM-DD}/{_dir}/interfaces/interfaces_{v}.md`.
**Do not ask for input** — generate everything at once.

**After generating the 5 proposals, submit for visual review.**

After visual review, ask the user which proposal to follow (Hybrid recommended, or A-E).

After selection, create `spec-product_{v+1}.md` incorporating the chosen interface (ASCII sketches).
