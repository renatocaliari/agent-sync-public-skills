## Phase 5: Tech Planning Sequencing

### 5a. Scope Generation

Read `references/tech-planning/` (TECH-CONTEXT.md, SCOPES-AND-SEQUENCING.md, TECH-OUTPUT.md, generation-principles.md) and launch a planning agent to:

1. Check strategic stability (Step 0)
2. Codebase awareness check (Step 1)
3. Technical risk analysis (Step 2)
4. Identify spikes (Step 3)
5. Define typed scopes: feature | optimization | spike (Step 4)
6. Sequence (riskiest-first or ui-first) (Step 5)
7. Detail each scope with DoD + acceptance criteria (Step 6)
8. Format per output-format.md (Step 7)

Output: `.cali-product-workflow/{YYYY-MM-DD}/{_dir}/plans/spec-tech_{v}.md`
Input: `.cali-product-workflow/{YYYY-MM-DD}/{_dir}/plans/spec-product_{v}.md`

⚠️ **Security check:** Read the YAML frontmatter of spec-product.md:
```bash
head -10 ...spec-product_{v}.md | grep "approved:"
```
- ✅ `approved: true` → proceed
- ❌ No `approved: true` → **GO BACK to Phase 4. Do not proceed.** This is deterministic — read the file, don't rely on memory.

### 5b. Conditional Review Gate

**If standalone (no Shape Up/Interface):** execute the Review Gate on `spec-tech.md` using your environment's approval mechanism.

After approval, stamp spec-tech.md (same procedure as Phase 4):
1. Add `approved: true` to YAML frontmatter
2. Create receipt
3. Freeze the file

**If post-Shape-Up:** the gate already ran in Phase 4 — skip this step.

### 5c. Goal Generation (Step 9)

After tech plan approval, convert each scope into a tracked goal with completion criteria:

**For each scope in the approved spec-tech.md:**
- Define ordered steps with completion criteria
- Document: DoD, acceptance criteria, dependencies
- Optimization/spike scopes with metrics become experiment loops instead
- Scopes with dependencies: start AFTER the dependency is complete
- If blocked: document the blockage reason
- Adjust scopes during execution if needed using available goal adjustment tools
