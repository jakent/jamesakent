# AGENTS.md

Personal site for jamesakent.com, built spec-first. Start with
`docs/roadmap.md` (phases) and `docs/adr/` (decisions, especially ADR 0001 on
where specs and docs live).

## Layout
- `specs/features/` — Gherkin features. **Human-owned.** Agents may propose new or
  changed features in a PR of their own (spec-writer work); CODEOWNERS requires
  the owner's approval before they merge. Implementation PRs must not touch them.
- `specs/model/` — entity model and glossary (Markdown)
- `specs/changes/<date>-<name>/` — proposal, tasks, optional design
- `docs/roadmap.md`, `docs/adr/` — living roadmap and append-only decisions
- `cmd/site` — HTTP server; `cmd/jak` — CLI (planned)
- `internal/` — shared packages used by the site, CLI and (later) MCP server
- `acceptance/` — godog suite: steps → DSL → drivers

## Workflow
One change at a time, in this order, each a skill in `.claude/skills/`:
`change` -> `write-feature` (owner approves) -> `implement-steps` -> `implement`.
Spec PRs and implementation PRs are always separate.

## Rules
- Behavior lives only in `.feature` files. Markdown holds everything else.
- Step definitions are one line and only call the DSL.
- Testability is a design rule: inject the clock, put external calls behind
  interfaces with fakes, scope config per request.
- Do not add `Co-Authored-By` or "Generated with" attribution lines to commits
  or PR descriptions. Agent involvement is already visible through the issue,
  the change folder and the PR trail.
- Secrets are never committed. Deployment belongs to the droplet-infra repo.

## Commands
- `make run` — start the site on :8080
- `make test` — unit tests
- `make vet` — go vet
- `make acceptance` — scenarios against the site in-process (skips `@stack`)
- `make up` / `make down` — local stack (Caddy, site, biff-repl, Jaeger) on
  http://localhost:8000 (override with `PUBLIC_PORT`); Jaeger UI at
  http://traces.localhost:8000
- `make acceptance-stack` — all scenarios, including `@stack`, against the running stack
