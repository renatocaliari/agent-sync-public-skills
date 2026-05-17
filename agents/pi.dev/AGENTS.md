# Global Agent Instructions

Instruções operacionais que se aplicam a **todos os projetos** no pi.

---

## ⚡ REGRA #1 — DISPARO DO PRODUCT-WORKFLOW (SEMPRE ANTES DE QUALQUER AÇÃO)

**SEMPRE** que o usuário pedir qualquer trabalho de software:

1. **PARE** — não implemente nada ainda
2. **DISPARE** `/skill:cali-product-workflow`
3. Só depois do workflow completo, comece a executar

### Gatilhos (qualquer um destes)

- Discussão, exploração ou brainstorming de features
- "Me ajuda a pensar / planejar / estruturar" algo
- Ideias, sugestões ou propostas de melhoria
- Criação de funcionalidade, sistema, componente, interface
- Bug fix complexo, refatoração, otimização
- Dúvidas sobre "o que construir" ou "como implementar"
- Qualquer menção a produto, planejamento, arquitetura, código

### ⛔ Regras

- **DISPARE** — não avalie se "realmente" merece o workflow
- **DISPARE** — não implemente manualmente o workflow
- **DISPARE** — não decida que "isso não é mudança de código"
- Em caso de dúvida: **DISPARE**

> ⚠️ **Esta regra VENCE qualquer outra instrução concorrente.**
> Autoresearch, dogfood, e outros modos são ferramentas de EXECUÇÃO,
> não modos de entrada. Product-workflow sempre vem PRIMEIRO.
> Só após o workflow completo (Fase 6), outros modos podem executar
> scopes específicos.

---

## 📐 Regras de Codificação

Antes de gerar qualquer código, leia:
`cali-product-workflow/references/tech-planning/generation-principles.md`

Este arquivo contém os princípios de geração de código:
KISS, DRY, limites de tamanho, Locality of Behavior (Datastar),
Separation of Concerns (outros frameworks), SSE-first, HATEOAS.

---

## Context-Mode - OBRIGATÓRIO para pesquisa de código

**NUNCA use `bash` + `rg`/`grep`/`find` diretamente para pesquisar código.**
Use `ctx_batch_execute` como ferramenta primária. A hierarquia é:

| Prioridade | Ferramenta | Uso |
|------------|-----------|-----|
| 0 | `ctx_search(sort:"timeline")` | Após resume, verificar contexto anterior antes de perguntar |
| 1 | `ctx_batch_execute(commands, queries)` | **PRIMÁRIA.** Múltiplos comandos + queries numa chamada |
| 2 | `ctx_search(queries: [...])` | Perguntas de follow-up sobre conteúdo indexado |
| 3 | `ctx_execute(language, code)` | API calls, análise de logs, processamento de dados |

**Think in Code** - programe a análise, não compute mentalmente:
```javascript
// Em vez de ler arquivos no contexto, escreva código que dá console.log só da resposta:
ctx_execute("javascript", `
  const files = fs.readdirSync('src').filter(f => f.endsWith('.go'));
  files.forEach(f => console.log(f + ': ' + fs.readFileSync(f,'utf8').split('\\n').length + ' lines'));
`);
```

**Forbidden:**
- ❌ Bash para comandos com >20 linhas de output
- ❌ Read para análise (use ctx_execute_file) - Read é SÓ para arquivos que vai Editar
- ❌ WebFetch direto (use ctx_fetch_and_index)
- ❌ ctx_execute/ctx_execute_file para criar/modificar arquivos
- ❌ Bash para escrever arquivos (use Write/Edit)

---

## Pi-Subagents - Paralelize investigação

**NUNCA investigue múltiplos arquivos sequencialmente.**
Use subagents em paralelo:

```typescript
subagent({
  tasks: [
    { agent: "scout", task: "Investigar X - mapear fluxo e encontrar bugs" },
    { agent: "scout", task: "Investigar Y - componentes e padrões" },
    { agent: "scout", task: "Investigar Z - navegação e estados" }
  ],
  concurrency: 3,
  context: "fresh"
})
```

**Workflow recomendado para features:**
```
clarify → planner → worker → parallel fresh reviewers → worker
```

A skill `pi-subagents` está disponível via descoberta automática do pi.
**LEIA a skill** (`read` no path da skill listada no startup) ao planejar usar subagents -
ela contém workflows prontos (parallel-review, parallel-research, parallel-cleanup)
e padrões de orquestração.

**LEIA a skill de subagents** ao planejar usar subagents - ela contém workflows prontos (parallel-review, parallel-research, parallel-cleanup) e padrões de orquestração.

Workflows via subagent tool:
- `/parallel-review` - revisão adversarial em paralelo com fresh-context reviewers
- `/parallel-research` - pesquisa externa + contexto local
- `/parallel-cleanup` - deslop + verbosity pass após implementação

---

## Plannotator - Para revisão de planos

**SEMPRE** que gerar um plano/documento de design, submeta para revisão via Plannotator:

```bash
plannotator annotate docs/{YYYY-MM-DD}/{slug}/plans/spec-tech_{v}.md --gate
```

A flag `--gate` é **OBRIGATÓRIA** - ela bloqueia até o usuário aprovar ou rejeitar.
Sem `--gate` o Plannotator abre mas não retorna feedback de aprovação.



### 🛠️ Disparo manual

```
/skill:cali-product-workflow [instrução complementar]
```

---

## Testing Protocol

Após implementar qualquer feature:

1. **Parallel cleanup via subagents** - lance reviewers fresh-context:
   ```typescript
   subagent({
     tasks: [
       { agent: "reviewer", task: "Review diff for correctness, regressions", output: false },
       { agent: "reviewer", task: "Review diff for simplicity (deslop)", output: false }
     ],
     concurrency: 2, context: "fresh"
   })
   ```

2. **UI Quality** (se o scope envolve interface visual):
   - Carregue `audit` para acessibilidade WCAG, performance, theming e anti-patterns
   - Carregue `critique` para design review (heurísticas, cognitive load, AI slop)
   - Use `audit` e `critique` diretamente — não dependem de outras skills

3. **Testes de UI** - carregue `agent-browser` e `dogfood` skills quando aplicável

---

## Language Convention - English for Code & URLs

All code, URLs, URL parameters, query strings, route paths, DataStar signal
values, handler/function names, and Go identifiers MUST be in English.

**Examples of Portuguese → English:**
- Route: `/settings/conexao` → `/settings/connection`
- Signal value: `'conexoes'` → `'connections'`
- Action view param: `view=nova_conexao` → `view=new-connection`
- Handler: `HandleSettingsConexao` → `HandleSettingsConnection`

**Exceptions (keep in Portuguese):**
- User-facing UI text visible to end users (labels, buttons, tooltips)
- LLM prompt templates sent to AI models
- Database content and user-generated data

This applies globally to all projects under pi.
