## Phase 0: Initial Questions

### 0x. Auto-Discovery Check (before anything else)

**BEFORE asking anything to the user**, verify the directory structure exists. Create `.cali-product-workflow/` if it doesn't exist.

Then scan for existing workflows:

```bash
count=0
for f in .cali-product-workflow/*/*/index.json; do
  if [ -f "$f" ] && grep -q '"workflow_status"[[:space:]]*:[[:space:]]*"in-progress"' "$f" 2>/dev/null; then
    echo "ACTIVE_WORKFLOW_FOUND:$f"
    cat "$f"
    count=$((count + 1))
  fi
done
if [ "$count" -eq 0 ]; then
  echo "NEW_WORKFLOW"
fi
```

**If 1 or more in-progress workflows exist**:
1. Read the found `index.json` files
2. If **only 1**: ask the user if they want to continue, view status, or cancel
3. If **multiple**: list them and recommend cleaning up stale ones

**If new workflow**: continue to 0a.

### Resume Mechanics (when resuming a workflow)

When resuming an existing workflow (the user explicitly asked to continue), follow this flow:

1. **Read the full `index.json`** — extract `name`, `current_phase_index`, `current_phase`, `artifacts`, `workflow_status`
2. **Read session checkpoints** if they exist — extract phase, step, pending decisions, user choices
3. **Survey existing artifacts** — check what spec, interface, critique, and plan files exist
4. **Map artifacts to completed phases**:
   - Approval in `.plannotator/approvals/` → that phase's gate has passed
   - `spec-product.md` exists → Phase 1 (Shape) completed
   - `interfaces.md` exists → Phase 2 (Interface) completed
   - `critique-report.md` exists → Phase 3 (Critique) completed
   - `spec-tech.md` exists and approved → Phase 5 (Planning) completed
5. **Determine resume point**:
   - If checkpoint with `phase == current_phase_index` → jump to that step
   - If `current_phase_index >= 5` and spec-tech approved → skip to Phase 6
   - Otherwise → start current phase from its first step
6. **DO NOT re-ask answered questions.** Use checkpoint data.
7. **Jump to the determined phase** and execute normally. Do not recreate existing artifacts.

### 0a. Workflow Steps

Ask the user which workflow stages to activate and whether to run impact analysis on existing code:

Present options for stages:
- **Shape Up Planning** (Recommended) — Understand problem, define IN/OUT scope. Generates spec-product.md. Automatically activates Plan Critique + Review Gate.
- **Interface Brainstorming** — Explore 5 interface directions. Automatically activates Plan Critique + Review Gate.
- **Tech Planning Sequencing** — Break into scopes with DoD + acceptance criteria.

Then ask about **impact analysis**: run regression/impact check on existing code before planning, or proceed directly.

If user selects no workflow stages: proceed to implementation directly (impact analysis still offered).

### 0b. Strategic Exploration (always ask)

**ALWAYS ask** if the user wants to explore strategic directions before planning.

Present options (multi-select):
- **Jobs To Be Done (JTBD)** — Map functional, emotional and social jobs
- **Evolutionary Principles** — Explore innovation via stepping-stones and optionality
- **Opportunity Mapping** — Map opportunities with ranked solutions
- **Market Analysis** — PESTLE, Foresight, Wardley Maps
- **Short-Cycle Product** — Quick idea validation with short learning cycles

**If user selects one or more approaches:**
1. Read `references/strategic-exploration.md` for each approach's details
2. Execute the selected ones **in parallel** using your environment's parallel task capabilities
3. Consolidate into `strategic-insights.md`
4. Incorporate outputs as Shape Up input

**If nothing selected:** proceed directly to Phase 1.

### Auto-chaining rules

| User selection | Phases that run automatically |
|---|---|
| Shape Up only | Shape Up → **Plan Critique** → **Review Gate** → Tech Planning (no gate) → Execution |
| Interface only | Interface Brain. → **Plan Critique** → **Review Gate** → Tech Planning (no gate) → Execution |
| Shape Up + Interface | Shape Up → Interface Brain. → **Plan Critique** → **Review Gate** → Tech Planning (no gate) → Execution |
| Tech Planning only | Tech Planning (with own **Review Gate**) → Execution |
| Shape Up + Tech Planning | Shape Up → **Plan Critique** → **Review Gate** → Tech Planning (no gate) → Execution |
| All | Shape Up → Interface Brain. → **Plan Critique** → **Review Gate** → Tech Planning (no gate) → Execution |

**Plan Critique** and **Review Gate** never appear as options — they are automatic.
**Review Gate** never duplicates.
