# Global Agent Instructions

## Skill Triggers

### New Features / Planning
When the request involves planning a new feature, the scope is unclear, or delivery boundaries are ambiguous, load the `shape-up-planning` skill to structure the problem and define scope before implementation.

### UI/UX / Interaction Design
When the request involves UI, UX, interaction flows, layouts, screens, wireframes, or frontend decisions, load the `interface-brainstorming` skill to generate strategically distinct proposals.

### Implementation Sequencing
After shaping converges and engineering decomposition or rollout order is needed, load the `tech-planning-sequencing` skill to produce a structured development plan with clear dependencies.

### Testing Uncommitted Changes
After making frontend changes, load the `agent-browser` skill to verify the UI in the browser. Use the `dogfood` skill to systematically explore and find issues before committing.


## Living Documentation
This project keeps a living spec in `docs/current_system.md`. After completing any feature implementation or making a significant change, create or update this file to reflect the current system state. Remove or update rules that no longer apply.

## Testing
Every feature must include tests. Detect the project's testing framework from its configuration files (go.mod, package.json, Cargo.toml, etc.) and use the standard conventions. Place tests alongside the code they test.

## Server Management (server.renatocaliari.com)
Before making any changes to services on server.renatocaliari.com, read the server guide:
- **Local**: `/Users/cali/Development/SERVER_GUIDE.md`
- **On server** (any LLM with SSH root): `/root/SERVER_GUIDE.md` or `srv-guide` alias

This guide covers all services (n8n, spacebot, paperclip, hermes-cali, hermes-lara, treinador), Cloudflare Tunnel + Zero Trust, Docker networks, Caddy config, and all known pitfalls.

<!-- codebase-memory-mcp:start -->
# Codebase Knowledge Graph (codebase-memory-mcp)

This project uses codebase-memory-mcp to maintain a knowledge graph of the codebase. ALWAYS prefer MCP graph tools over grep/glob/file-search for code discovery.

## Priority Order
1. `search_graph` — find functions, classes, routes, variables by pattern
2. `trace_path` — trace who calls a function or what it calls
3. `get_code_snippet` — read specific function/class source code
4. `query_graph` — run Cypher queries for complex patterns
5. `get_architecture` — high-level project summary

## When to fall back to grep/glob
- Searching for string literals, error messages, config values
- Searching non-code files (Dockerfiles, shell scripts, configs)
- When MCP tools return insufficient results

## Examples
- Find a handler: `search_graph(name_pattern=".*OrderHandler.*")`
- Who calls it: `trace_path(function_name="OrderHandler", direction="inbound")`
- Read source: `get_code_snippet(qualified_name="pkg/orders.OrderHandler")`
<!-- codebase-memory-mcp:end -->
