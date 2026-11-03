# AGENTS.md

Personal site for jamesakent.com, built spec-first. Start with
`docs/roadmap.md` (phases) and `docs/adr/` (decisions, especially ADR 0001 on
where specs and docs live).

## Layout
- `specs/features/` — Gherkin features. **Human-owned. Agents never edit these.**
- `specs/model/` — entity model and glossary (Markdown)
- `specs/changes/<date>-<name>/` — proposal, tasks, optional design
- `docs/roadmap.md`, `docs/adr/` — living roadmap and append-only decisions
- `cmd/site` — HTTP server; `cmd/jak` — CLI (planned)
- `internal/` — shared packages used by the site, CLI and (later) MCP server
- `acceptance/` — godog suite: steps → DSL → drivers

## Rules
- Behavior lives only in `.feature` files. Markdown holds everything else.
- Step definitions are one line and only call the DSL.
- Testability is a design rule: inject the clock, put external calls behind
  interfaces with fakes, scope config per request.
- Secrets are never committed. Deployment belongs to the droplet-infra repo.

## Commands
- `make run` — start the site on :8080
- `make test` — unit tests
- `make vet` — go vet
