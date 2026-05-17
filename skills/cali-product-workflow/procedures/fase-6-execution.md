## Phase 6: Supervision + Execution

### ⚠️ Activate supervision ONLY during execution
**Never activate during Phases 1-5.** It could interfere with earlier phases.

### 6a. Git Worktree Check (before executing scopes)

**Check if the workflow is in an isolated directory.** If multiple workflows modify the same code in parallel, use a worktree to avoid conflicts:

```bash
if git rev-parse --git-dir | grep -q '.git/worktrees/'; then
  echo "INSIDE_WORKTREE"
else
  echo "MAIN_REPO"
fi
```

**If in the main repository and modifying code:** Ask the user if they want to create an isolated branch + worktree:
- Yes: create `.worktrees/pw-{name}-{date}`, branch `pw/{name}/{YYYY-MM-DD}`
- No: execute in current directory (only for 1 workflow at a time)

If creating a worktree:
```bash
BASE_BRANCH=$(git remote show origin 2>/dev/null | grep "HEAD branch" | cut -d" " -f5 || echo "main")
git fetch origin 2>/dev/null || true
git worktree add .worktrees/pw-{name}-{date} -b pw/{name}/{YYYY-MM-DD} "$BASE_BRANCH"
cd .worktrees/pw-{name}-{date}
```

- Copy the approved plan (`.cali-product-workflow/`) to the worktree if needed
- Execute all scopes inside the worktree
- At the end, ask: "Commit + push?" and "Remove worktree?"

> The worktree is optional. Workflows with 1 scope or no code changes can skip.

### 6b. Scope Executor Routing

| Scope Type | Has Metric? | Executor Approach | Supervision |
|---|---|---|---|
| Optimization | Yes | Experiment loop driven by metrics | Self-supervising (metric-driven) |
| Scope with autoresearch | Yes | Experiment loop | Self-supervising |
| Spike with metric | Yes | Experiment loop | Self-supervising |
| Feature | No | Tracked goal with ordered steps | Supervised with outcome = DoD |
| Refactoring (no metric) | No | Tracked goal | Supervised with outcome = DoD |
| Investigative spike | No | Tracked goal | Supervised with outcome = DoD |
| Interface brainstorming | No | Tracked goal (5 proposals) | Supervised with outcome = DoD |

### When starting execution of each scope:

1. **Feature/refactor/spike without metric:** Create a tracked goal with ordered steps and DoD as completion criteria. Activate supervision with: "Execute scope with DoD and acceptance criteria, do not deviate from approved scope."

2. **Optimization/spike with metric:** Set up an experiment loop driven by metrics. No external supervision needed — the metric loop is self-supervising.

3. **If blocked:** Document the blockage reason.

4. **Scope adjustment:** Adjust scope during execution if needed.

> **Supervision is especially useful for long scopes where focus may drift. Activate WHEN STARTING each scope, not before.**

After tech planning, route each scope to its appropriate executor based on its type.
