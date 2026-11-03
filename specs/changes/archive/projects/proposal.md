# Projects slice

## Why
The site needs its first real content: a list of projects with links to source
and live demos. It is also the first change run through the full spec-first
loop (feature -> steps -> DSL -> drivers -> code).

## Scope
- `specs/features/projects.feature` (new; needs owner approval)
- `content/projects.yaml` embedded in the binary, validated against the
  invariants in `specs/model/entities.md`
- `GET /api/projects` (JSON, optional `?tag=`) and an HTML projects page at `/`
- Acceptance: api and http drivers both run the scenarios

Out of scope: hosting the demos on subdomains (separate change); the REPL and
`jak` CLI listing projects (later drivers).

## Decisions (owner, 2026-11-03)
- Initial project list and wording as drafted: Clojure REPL, Game of Three,
  Game of Three-D, Game of Biff, Game of Clojure.
- Fixtures first: acceptance scenarios run in-process against fixture content
  (`acceptance/fixtures/projects.yaml`). Demo URLs in the real content may point
  at subdomains that do not exist yet.
- No `kind` field. "Has a live demo" means `demo_url` is set.

## Testability
- Content is data, loaded through an interface, so scenarios can use a fixture set.
- No time, randomness or external calls involved.
- Scenarios run through the api and http drivers on every PR; the ui driver
  only checks layout later.

## Tasks
1. Owner reviews and approves `projects.feature` (this PR).
2. `implement-steps`: steps, DSL (`ListProjects`, `ListProjectsTagged`, ...), api and http drivers.
3. `implement`: content package, validation tests, `/api/projects`, projects page.
4. Archive this folder after merge.
