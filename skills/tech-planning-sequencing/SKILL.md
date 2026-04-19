---
name: tech-planning-sequencing
description: >
  Generate structured development plans with risk-based or UI-first sequencing. Analyzes tasks,
  identifies critical spikes, and produces detailed implementation sequences with DoD and acceptance
  criteria. Use when planning new features, PRDs, or complex task breakdowns. Activate this skill
  automatically or when any of these signals are detected:
  - Keywords / intents: "sequence tasks", "prioritize", "riskiest first", "dependencies", "group scopes", "technical planning", "plan spikes", "order implementation"
  - Workflow stage: after high-level brainstorming/pitch, during or before detailed task breakdown, when dependencies or risk order are unclear
  - Always before executing complex or multi-step technical work
---

# When to use this skill

Use this skill when you need to:
- Plan the implementation sequence for a new feature or PRD
- Decompose tasks into a risk-aware development order
- Identify critical spikes before implementation
- Structure work into scopes with clear Definition of Done

# When NOT to use this skill

- For simple, well-understood tasks that don't need sequencing
- When the user asks for code implementation directly (switch to Code mode)
- For architectural decisions without task breakdowns (use Architect mode directly)

## Activation Triggers (Universal)
Activate this skill automatically or when any of these signals are detected:
- Keywords / intents: "sequence tasks", "prioritize", "riskiest first", "dependencies", "group scopes", "technical planning", "plan spikes", "order implementation"
- Workflow stage: after high-level brainstorming/pitch, during or before detailed task breakdown, when dependencies or risk order are unclear
- Always before executing complex or multi-step technical work

## Output Guidelines for Any Orchestrator
- Default mode: riskiest-first (identify & front-load spikes)
- Alternative modes: UI-first, suggestive (infer from context)
- Group into scopes with explicit DoD (Definition of Done) and acceptance criteria per scope
- Use tags: [SCOPE-1] [DOD] [RISK-ORDER] [SPIKES] [SEQUENCE]
- After output, suggest: "Approve sequence? Adjust mode? Proceed to execution?"

## GENERATED CODE AND PRODUCT FEATURES
All generated code and product FEATURES must follow the Generation Principles (see GENERATION_PRINCIPLES.md or linked section).

## Portability Note
Self-contained and framework-agnostic. Works in any orchestration that supports intent detection or explicit skill calls.

# Inputs required

The user must provide:

1. **sequencing_strategy** (optional): `riskiest-first` (default) or `ui-first`
   - `riskiest-first`: Prioritize high-risk tasks and spikes early
   - `ui-first`: Prioritize UI/UX deliverables to get early feedback

2. **analysis_mode** (optional): `suggestive` (default) or `strict`
   - `suggestive`: Identify and suggest risks, spikes, and enablers not explicitly marked
   - `strict`: Only work with explicitly marked risks and tasks

3. **tasks** (required): PRD, solution proposal, or task list
   - May include markers: `nice-to-have` or `risky`
   - Format: bullet list or structured document

# Workflow

1. **Read and analyze the task list**
   - Identify any markers ("nice-to-have", "risky")
   - Understand the domain and context

2. **Identify critical spikes**
   - Look for tasks marked as "risky"
   - If `analysis_mode` = `suggestive`, identify additional fundamental uncertainties
   - Define "initial critical spikes scope" if applicable

3. **Define main functional scopes**
   - Identify and name the main functional scopes
   - First scope must be "domain core functionality"
   - Group related tasks logically

4. **Sequence scopes at high level**
   - Apply `sequencing_strategy` to order scopes
   - Justify the sequence based on the chosen strategy

5. **Detail task sequence per scope**
   - Use principles 0-6 (see [references/SEQUENCE_PRINCIPLES.md](references/SEQUENCE_PRINCIPLES.md))
   - For each task: name, objective, components, justification, acceptance criteria
   - Mark suggested items with `💡 [suggestion]` when `analysis_mode` = `suggestive`

6. **Output in specified format**
   - Follow the complete template in [references/OUTPUT_FORMAT.md](references/OUTPUT_FORMAT.md)

# References

- **[references/PRINCIPLES.md](references/PRINCIPLES.md)** - The 6 sequencing principles for task ordering
- **[references/OUTPUT_FORMAT.md](references/OUTPUT_FORMAT.md)** - Complete output template with all sections
