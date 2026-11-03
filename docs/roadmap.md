# Plan: jamesakent.com, an agent-native, observable, SDD-first personal site in Go

## Context
You want a hosted site at **jamesakent.com** that does three things:
- showcases your toy apps (the Game of Life variants and the Clojure REPL at `/repl`)
- shows OpenTelemetry tracing **publicly**: visitors browse Jaeger to see how requests flow
- shows the work itself: a public ticket board, public specs, and agent PRs

It is also where you learn to build a **software factory** with spec-driven development (SDD), reusing the KAOS approach (from the author's private notes; the relevant ideas are restated here):
- the chain is Gherkin features → functional DSL → protocol drivers → system
- humans own the features, as holdout scenarios
- agents implement; skills and subagents do the work
- work is tracked in change folders

**Decisions so far:**

| Area | Decision |
|---|---|
| Language | Go: site, CLI, (later) MCP server, OTel and the godog acceptance suite in one module |
| Repo | New public repo that owns `/` |
| Demos | On subdomains |
| Tracing | Public Jaeger v2 |
| Tickets | GitHub Issues + Projects, plus a `/work` page on the site |

**CLI vs MCP:**
- **`jak` CLI**: the factory's tool, for you, Claude Code skills and CI.
- **MCP server**: an optional, later public showcase for agents without shell access (visitors, claude.ai).
- Both are thin wrappers over the same `internal/` packages.

Current state (from `droplet-infra/` and `biff-repl/`):
- Caddy sends **all** of `jamesakent.com` to `biff-repl:8080`.
- biff-repl's `/` redirects to `/repl` (`biff-repl/src/biff_repl/home.clj:8`), and it serves `/css/*`, `/js/*` and `/img/*` at the root (`biff-repl/src/biff_repl/ui.clj:28`).
- Deploy pipeline: app repo → GHCR → `repository_dispatch` → `droplet-infra/.github/workflows/deploy.yml` → SSH. Reuse it.
- Demos:
  - `game-of-three` and `game-of-three-d`: CLJS, built as static files
  - `game-of-biff`: a Biff server app
  - `game-of-clojure`: desktop JVM, so a link + GIF

## Architecture

```
Caddy  (tracing directive: root span + W3C traceparent → OTLP to jaeger:4317)
├─ jamesakent.com
│    /repl*  /css/*  /js/*  /img/*   → biff-repl:8080   (OTel Java agent, phase 3)
│    everything else                  → site:8080        (Go, otelhttp)
├─ www.jamesakent.com                 → redirect (unchanged)
├─ traces.jamesakent.com              → jaeger:16686     (query UI only, public)
├─ life.jamesakent.com / life3d.…     → file_server /srv/demos/<name>
└─ mcp.jamesakent.com                 → site:8080/mcp    (phase 5, optional)

Jaeger v2 (single container, internal OTLP 4317/4318 never exposed)
  processors: attributes/redaction (drop client IP, cookies, auth headers; truncate eval bodies),
              probabilistic or tail sampling; storage: badger with 72h TTL on a named volume
```

### Caddy and compose changes (droplet-infra)
- `Caddyfile`:
  - Replace the single `reverse_proxy biff-repl:8080` with the path split above, using a `@repl` matcher and `handle` blocks.
  - Add `traces.`, `life.` and `life3d.` blocks, and later `mcp.`.
  - Add the `tracing` directive.
  - Delete the stale "TODO: Replace domain" comment.
  - Keep the `www` redirect.
- `docker-compose.yml`:
  - Add the `site` and `jaeger` services on the `web` network, with a jaeger volume and memory limits.
  - Set `OTEL_EXPORTER_OTLP_ENDPOINT=http://jaeger:4317` and `OTEL_SERVICE_NAME` on caddy, site and biff-repl.
- `jaeger/config.yaml`: the Jaeger v2 config (receivers, the redaction processor, badger TTL).
- Hetzner DNS: A records for `traces`, `life` and `life3d`, or one wildcard.
- `deploy.yml`: add a `static` payload type that downloads a demo artifact and rsyncs it to `/srv/demos/<name>`.
- `CLAUDE.md`: update the project table.
- biff-repl follow-up: move its assets under `/repl/static/*` and remove the `/` redirect, so the matcher shrinks to `/repl*`.

### Repo: `jakent/jamesakent` (public, Go module)
```
AGENTS.md                         # main agent instructions; CLAUDE.md → @AGENTS.md
specs/
  features/                       # HUMAN-OWNED holdouts (CODEOWNERS + CI guard)
    routing.feature projects.feature tracing.feature work.feature cli.feature demos.feature [mcp.feature]
  model/  entities.md  glossary.md   # Project, Demo, Ticket (GitHub Issue view), Change
  changes/<date>-<name>/  proposal.md  tasks.md  [design.md]   → changes/archive/
content/  projects.yaml  about.md     # go:embed; validated by unit tests
cmd/
  site/   # HTTP server (templ + htmx); mounts /mcp in phase 5
  jak/    # CLI (cobra): projects list|add, smoke, trace <id>, work list, deploy status
internal/
  content/    # Project/Demo types, loading, validation
  web/        # handlers, templ components
  api/        # /api/projects, /api/work JSON
  telemetry/  # OTel setup (traces, metrics, logs), otelhttp, trace-id in page footer
  work/       # GitHub client (Issues + Projects v2 GraphQL), cached; links issue → change folder, @req scenarios, PR, CI
  jaeger/     # query-API client (used by `jak trace` and the tracing specs)
  mcpserver/  # phase 5: list_projects, get_project, list_work, get_trace, eval_clojure (→ biff-repl)
acceptance/   # godog, run with `go test ./acceptance`
  steps/      # one-line steps → DSL
  dsl/        # plain funcs, ctx first: dsl.ListProjects(ctx)
  drivers/    # http.go api.go ui.go (playwright-go) cli.go [mcp.go] — one interface
  fakes/      # fake GitHub API (alias-aware), so work.feature needs no real network
  support/    # world/ctx, profiles (BASE_URL, DRIVER, tags)
scripts/traceability/  # every @req has a scenario; no undefined/unused steps
.claude/skills/  .claude/agents/
Dockerfile (distroless)  docker-compose.yml (site, biff-repl, caddy, jaeger, github-fake)  Makefile
.github/workflows/  ci.yml  release.yml  agent.yml
```

### Executable-spec chain
Entity model → `.feature` files (behavior; `Rule:` = invariants) → DSL (verbs) → drivers → system. A single driver interface means the same scenario runs through http, api, ui, cli and (later) mcp.

```gherkin
@req:SITE-002
Feature: Projects
  Rule: Every project links to its source, and to a live demo when one exists
    Scenario: A visitor finds the Game of Life demos
      When a visitor lists the projects
      Then "Game of Three-D" has a live demo
      And every project links to its source repository
```
```go
sc.When(`^a visitor lists the projects$`, func(ctx context.Context) error { return dsl.ListProjects(ctx) })
```
Other features:
- **`tracing.feature`**: "When a visitor evaluates `(+ 1 2)` in the REPL, then that request's trace has spans from caddy, site and biff-repl, and has no client IP attribute." The DSL sends a known `traceparent` and polls Jaeger's query API. This tests both the trace flow and the redaction.
- **`work.feature`**: "An issue labeled `agent` with a linked change folder shows on /work in the Spec'd column, with links to its scenarios." It runs against the GitHub fake on every PR and as `@smoke` against real GitHub.

**When each suite runs:**

| Suite | When | Target |
|---|---|---|
| All scenarios | Every PR | http/api/cli drivers + fakes |
| `@journey` | Merge and nightly | ui driver (playwright-go, plus axe) |
| `@smoke` | After each deploy | Prod: `/`, `/repl`, `traces.`, `/work`, demo subdomains |

## Phases: build the factory while building the site

**Phase 0: walking skeleton (L0)**
1. Create the public repo with branch protection, CODEOWNERS on `specs/features/**`, CI (`go vet`, `golangci-lint`, `go test`, godog) and a Makefile. `make up` runs the compose stack with prod-identical routing on `*.localhost`; `make acceptance DRIVER=http`.
2. `routing.feature` (`@smoke`). OTel goes in from day one: otelhttp, Caddy `tracing`, and Jaeger v2 with redaction and TTL config. The footer shows "view this page's trace" → `traces.…/trace/{id}`.
3. Release → GHCR → dispatch. Make the droplet-infra Caddy and compose changes above. Deploy and run `jak smoke --env prod`.

**Phase 1: projects slice**
4. Write `entities.md` and `glossary.md`. You write `projects.feature`. Then build `projects.yaml`, the validation tests, `/api/projects`, the templ page, and the http/api/ui drivers.
5. CI traceability gates: undefined steps, unused steps, and an `@req` without a scenario all fail the build.

**Phase 2: work tracking, CLI and skills (L1)**
6. GitHub setup: a public Project board with columns Backlog → Spec'd → Agent → Review → Done (plus Blocked), and labels `agent-eligible` and `human-only`. An issue template requires a link to a change folder and the `@req` IDs it covers.
7. `internal/work` + `/work` page + `work.feature` against the GitHub fake. Cache GitHub API responses so visitor traffic can't exhaust the rate limit.
8. `jak` CLI with a cli driver (`cli.feature` reruns the projects and work scenarios through the CLI).
9. `AGENTS.md`:
   - commands and layout
   - "behavior lives only in `.feature` files"
   - agents never edit `specs/features/**`
   - testability rules: injectable clock, external calls behind interfaces with fakes, scoped config
10. Skills in `.claude/skills/` (they call `jak` and make wherever possible):

    | Skill | What it does |
    |---|---|
    | `change` | Change folder + GitHub issue, linked |
    | `write-feature` | Draft scenarios for your review |
    | `implement-steps` | Steps, DSL and drivers only |
    | `implement` | App code until the change's scenarios pass; archives the folder |
    | `add-project` | Add a project |
    | `add-demo` | Caddy block, DNS, demo release workflow, `@smoke` scenario |
    | `debug-trace` | `jak trace` / Jaeger API to diagnose a request |
    | `review-pr` | Holdout check, traceability, testability |

**Phase 3: tracing showcase**
11. Add the OTel Java agent to biff-repl's Dockerfile so REPL evals join the trace (Go + JVM). Write `tracing.feature`, including the redaction assertions.
12. A "How requests flow" page that explains the system and links Jaeger's System Architecture graph and some example traces.

**Phase 4: demos on subdomains**
13. game-of-three and game-of-three-d: a release workflow builds static files, uploads an artifact and dispatches a `static` payload; deploy rsyncs it to `/srv/demos/<name>`. Add the Caddy blocks and DNS.
14. `demos.feature` `@smoke`: each demo returns 200 and renders a canvas (ui driver).
15. game-of-biff: optional server app. game-of-clojure: link + GIF, and a later factory exercise to port it.

**Phase 5: public MCP server (optional showcase)**
16. Use the official `modelcontextprotocol/go-sdk` over Streamable HTTP at `mcp.jamesakent.com`.
    - Read-only tools that reuse `internal/` packages.
    - `eval_clojure` proxies to biff-repl's sandbox and inherits its rate limit and timeout.
    - Rate-limit per IP.
17. Add an mcp driver so the projects and work scenarios also run through MCP. Add a "Use this site from your agent" page.

**Phase 6: delegation (L2), then L3**
18. Subagents in `.claude/agents/`:

    | Subagent | Scope |
    |---|---|
    | `spec-writer` | Proposals and draft features only |
    | `implementer` | Everything except `specs/features/**` |
    | `reviewer` | Read-only; runs the `review-pr` checks |
    | `infra` | droplet-infra, via `--add-dir` |

19. `agent.yml`: when an issue is labeled `agent` and has a linked change folder, headless Claude Code moves the card to Agent, creates a branch, opens a PR and moves the card to Review. CI runs and you review. Adapt droplet-infra's `claude.yml`. CODEOWNERS requires the owner's approval for any change under `specs/features/**`, so agents can propose specs but not merge them. A CI guard also fails an `implementer` PR that touches `specs/features/**`; spec changes go in a separate `spec-writer` PR. Agents never get prod credentials.
20. L3 (later):
    - a nightly `@smoke` that files an issue with a trace link when it fails
    - dependency-bump agents
    - worktrees for parallel changes
    - factory metrics (lead time, agent PRs merged without rework) shown on `/work`

## What to highlight on the site
1. **"How this site is built"**: the factory itself. The living Cucumber report at `/specs`, the public `/work` board, the change folders, the four-layer diagram, and agent PR stats.
2. **Observability, live**: "view this page's trace", the public Jaeger UI with its service graph, and the cross-language REPL trace (Caddy → Go → JVM).
3. **Playground**: the REPL and the Game of Life variants side by side, the same idea in Clojure, CLJS/Reagent, 3D and Biff/HTMX.
4. **Agent-native**: the `jak` CLI and, later, a public MCP server.
5. **BDD/open-source lineage**: clj-cucumber, burpless, funcade.
6. **Writing**: posts on the SDD and factory journey. Describe patterns only; keep client names (KAOS, Fanatics) out.
7. **About/Now**: your role at Chariot, what you're learning, contact and resume.

## Verification
- `make up`:
  - `localhost/` shows the site
  - `localhost/repl` evaluates `(+ 1 2)`
  - the footer trace link opens the trace in Jaeger at `traces.localhost`, with caddy/site/biff-repl spans and no client IP
  - `/work` shows issues from the GitHub fake
  - the demo subdomains resolve
- `make acceptance` passes with `DRIVER=http|api|cli` and `ui` for `@journey`. The traceability script reports no gaps.
- After deploy, `jak smoke --env prod` passes against `/`, `/repl`, `traces.jamesakent.com`, `/work` and each demo subdomain.
- Check that ports 4317/4318 aren't reachable from outside: `nc -z jamesakent.com 4317` fails.
- Factory check: open an `agent-eligible` issue (e.g. "add game-of-biff to projects") with a change folder and label it `agent`. Confirm that:
  - the card moves across `/work`
  - the agent opens a PR and CI is green
  - no `specs/features/**` file is touched (spec changes arrive in their own PR for your approval)
  - you only had to review

## Files touched outside the new repo
- `droplet-infra/`: `Caddyfile`, `docker-compose.yml`, new `jaeger/config.yaml`, `.github/workflows/deploy.yml` (static payload), `CLAUDE.md`, Hetzner DNS records
- `biff-repl/`: `src/biff_repl/home.clj` (remove the `/` redirect), `src/biff_repl/ui.clj` (asset paths), `Dockerfile` (OTel Java agent + env)
- `game-of-three/` and `game-of-three-d/`: new `.github/workflows/release.yml`
