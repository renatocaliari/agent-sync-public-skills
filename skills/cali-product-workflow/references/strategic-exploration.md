# Strategic Exploration (Phase 0b)

⚠️ **Read this file only when the user shows interest in strategic exploration.**
Otherwise, skip straight to Phase 1.

---

## Detect Signals

Look for user mentions of:
- Strategic direction: "how to evolve", "new features", "opportunities", "strategy"
- Methods: JTBD, jobs-to-be-done, evolutionary, opportunity mapping, short-cycle
- Exploration: "what to build", "ideas for the product"

## Ask

Ask the user which strategic approaches they want to explore before Shape Up. Present these options as multi-select:

- **Job-to-Be-Done Framework** — Analysis of functional, emotional and social jobs. Undeclared needs.
- **Evolutionary Product Thinking** — Stepping-stones, evolutionary forces, optionality, avoiding premature convergence.
- **Opportunity Mapping** — Opportunities ranked by impact and effort. Quick wins + strategic bets.
- **Short-Cycle Product Method** — Validation with quick experiments. Metrics, channels, pricing, business model.
- **Multi-Method Market Analysis** — PESTLE, Wardley Maps, Delphi, Foresight — deep market analysis.

## Execution

If user selects one or more:

1. Run each selected skill in parallel using your environment's parallel task capabilities. Each task should:
   - Execute the analysis using the corresponding skill
   - Save output to `.cali-product-workflow/{YYYY-MM-DD}/{_dir}/strategic/{skill-name}.md`

2. Consolidate into `strategic-insights.md` with:
   - Executive summary (10-15 bullets)
   - Links to full analyses
   - Top opportunities consolidated
   - Recommended focus areas

3. Show summary in chat with file links

4. For each skill, ask the user which specific insights to incorporate into Shape Up (multi-select).

5. Integrate selected insights into Shape Up:
   - Inject as context
   - Add sections to spec-product.md (e.g. "Considered Jobs")

If user selects nothing → skip to Phase 1.
