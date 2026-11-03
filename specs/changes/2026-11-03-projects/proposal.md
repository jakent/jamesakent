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

## Open questions for the owner
- Initial project list and wording. Draft: Clojure REPL, Game of Three,
  Game of Three-D, Game of Biff, Game of Clojure.
- Demo URLs are placeholders until the subdomains exist (`life.`, `life3d.`).
  Until then should demo scenarios run against fixture content, or should the
  live demos be hosted first?
- Is `kind` worth keeping, or should "has a demo" just mean `demo_url` is set?

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
