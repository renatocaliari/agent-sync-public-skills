## Phase 3: Plan Critique

### 3a. Analysis via parallel review

Launch a review agent with checklists from `references/plan-critique/`:
- Read: CHECKLISTS.md (flows, states, affordances, data, system, feasibility) and output-format
- Output: Executive Summary + Critical Questions (🚨) + Important (🤔) + Minor (🔎) + Strengths
- **Do not resolve gaps** — only identify and classify
- Save to `.cali-product-workflow/{YYYY-MM-DD}/{_dir}/plans/critique-report.md`

### 3b. Gap Resolution

Ask the user: **Auto-resolve** (applies rules from `auto-resolve-rules.md`) or **Manual** (ask one by one).

- 🔎 is always automatic
- Auto-resolve: save pre-critique version, create new spec with "Resolved Gaps" section, show summary before proceeding
- Manual: ask about each 🚨 and 🤔 individually
- After resolving, create updated spec with documented resolutions
