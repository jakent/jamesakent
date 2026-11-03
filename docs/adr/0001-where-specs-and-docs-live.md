# 0001. Where specs and docs live

Status: accepted (2026-11-03)

## Context
This repo is built spec-first by humans and agents, possibly from several
machines. Plans and decisions kept outside the repo (for example in a local
`~/.claude/plans/` folder) are invisible to other machines and to agents.

## Decision
Everything an agent needs to work lives in the repo, in fixed places:

| Kind | Location | Lifetime |
|---|---|---|
| Behavior (features, rules, acceptance criteria) | `specs/features/*.feature` | permanent, human-owned |
| Entity model and glossary | `specs/model/` | permanent |
| Change proposals (why, scope, tasks, testability) | `specs/changes/<date>-<name>/`, moved to `specs/changes/archive/` after merge | per change |
| Architecture decisions | `docs/adr/NNNN-title.md` | permanent, append-only |
| Roadmap | `docs/roadmap.md` | living; updated in the PR that changes it |
| Agent instructions | `AGENTS.md` (`CLAUDE.md` imports it), `.claude/skills/`, `.claude/agents/` | permanent |

Rule: behavior lives only in `.feature` files. Markdown never restates
scenarios; it points at them.

Tool-local plan files (such as `~/.claude/plans/`) are scratch space. Anything
worth keeping is committed here before work starts from it.

## Consequences
- A fresh clone on any machine has the full context.
- Specs and code change in the same PR, so history shows why.
