# Environment Adaptation

This skill is designed to work in multiple environments. The tools available may vary.

## General approach

1. **Try the most obvious approach first.** If your environment has a built-in way to do it, use it.
2. **If a specific tool is not available**, describe what you need to accomplish and let your environment handle the execution.
3. **The intent matters more than the tool name.** Every step in this workflow can be adapted.

## Common adaptations

| Step in workflow | Generic approach | If your environment has no dedicated tool |
|---|---|---|
| Ask user a question | Present structured options and wait for selection | List the options in your response and ask for a choice |
| Delegate parallel tasks | Use subagent/child agent capability | Execute sequentially or use bash for parallel jobs |
| Visual review of a document | Submit for review using environment's review mechanism | Save the file, ask user to review, wait for feedback |
| Formal approval gate | Require explicit approval with timestamp | Create a receipt file manually, ask user to confirm |
| Create tracked goals | Define ordered steps with completion criteria | Use a todo checklist with verification steps |
| Supervision during execution | Monitor progress against DoD | Re-check after each step, ask user if deviating |
| Impact analysis | Run tests or static analysis before planning | Review manually or skip if no tool available |

## Identical tools (work in all environments)

```
read   — read files
bash   — run commands
write  — write files  
edit   — edit files
grep   — search in files
```
