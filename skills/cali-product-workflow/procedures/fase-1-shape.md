## Phase 1: Shape Up Planning

### 1a. Parallel Recon (optional — recommended for complex features)

Before shaping, explore the codebase in parallel:
- Launch an exploration agent to map current code state, relevant files, existing flows, and impact points
- Launch another to map technical risks, external dependencies, and constraints

Read both outputs before proceeding.

### 1b. Shaping

Read the shape-up section references to guide the process:

- **`references/shape-up/SHAPING-COMPLETE.md`** — context, clarification, responsibilities
- **`references/shape-up/SHAPING-PRINCIPLES.md`** — core and shaping principles
- **`references/shape-up/RISK-ANALYSIS.md`** — risk analysis and strategic alternatives
- **`references/shape-up/EXECUTION-GUIDE.md`** — sequencing, persistence, cross-domain
- **`references/shape-up/proposal-structure.md`** — shaping output structure
- **`references/shape-up/output-expectations.md`** — strong vs weak output criteria

Ask the user strategic questions when needed.

After shaping:
- Save to `.cali-product-workflow/{YYYY-MM-DD}/{_dir}/plans/spec-product_{v}.md`
- Do not ask about Interface Brainstorming — already decided in Phase 0

### 1c. Scope Adjustment (after Shape Up)

Show the IN/OUT scope table. Ask the user:
1. Remove any IN scopes?
2. Include any OUT scopes?

If user selects nothing → proceed with original Shape Up.
If changes exist → create `spec-product_{v+1}.md` with adjusted scopes and document what changed.
