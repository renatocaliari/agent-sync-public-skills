## Phase 4: Review Gate

### 4x. Claim Verification (before the Gate)

**BEFORE submitting for approval**, verify all file:line references in the spec:

```bash
grep -E '\`[^\`]+:[0-9]+\`' .cali-product-workflow/{YYYY-MM-DD}/{_dir}/plans/spec-product_{v}.md | \
  sed 's/.*\`\([^\`]*:[0-9]*\).*/\1/' | \
  sort -u > /tmp/refs_to_verify.txt
```

**For each reference**, reopen the file and verify:
1. Does the mentioned code/class/function exist?
2. Does the line match what the spec claims?
3. Are there discrepancies?

Generate a claim verification report with: Verified ✅, Discrepancies ⚠️, Not Found ❌.

**If there are discrepancies**: fix the spec before the Gate, document corrections, add note.

### Review Gate

⚠️ **SAFETY RULES — DO NOT SKIP:**
1. **Verbal approval in chat does NOT replace the gate.** Require a formal approval mechanism.
2. **Formal approval is MANDATORY.** Only proceed after explicit approval is recorded.
3. If changes are requested, adjust and re-submit.
4. After approval, the spec is frozen. Future changes require a new version and new gate.

**Submit the spec for formal review and approval using whatever review mechanism your environment provides.**

**IMPORTANT — After approval, stamp the spec:**

1. Add to the YAML frontmatter:
   ```yaml
   approved: true
   approved_at: "<timestamp>"
   approved_via: "<review mechanism used>"
   ```

2. Create an approval receipt at `.plannotator/approvals/{_dir}/spec-product_{v}.approved.md`:
   ```bash
   mkdir -p .plannotator/approvals/{_dir}
   ```
   Save a receipt with: approval timestamp, spec hash, verdict.

3. **Frozen file:** Spec can NOT be changed after stamping. Future revisions create `spec-product_{v+1}.md`.

4. Subsequent phases check for `approved: true` in the frontmatter.

> **If only Tech Planning was selected (standalone):** the Review Gate runs at the end of Tech Planning, not here.
