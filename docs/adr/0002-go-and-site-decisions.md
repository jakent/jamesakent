# 0002. Go, public tracing, and GitHub-backed work tracking

Status: accepted (2026-11-03)

## Decisions
- **Go** for the site, the `jak` CLI, the optional MCP server, OTel
  instrumentation and the godog acceptance suite: one module, small static
  binaries, mature OTel and MCP SDKs. KAOS used TypeScript; this project is a
  chance to learn something new.
- **This repo is public and owns `/`.** The existing biff-repl app stays at
  `/repl` behind Caddy.
- **Demos on subdomains** (`life.`, `life3d.`), served as static files by Caddy.
- **Jaeger v2 is public** at `traces.jamesakent.com` (query UI only). OTLP
  ingest ports stay internal. Data is redacted (client IPs, cookies, auth
  headers, truncated eval bodies) and expires after about 72 hours.
- **Tickets are GitHub Issues + Projects.** The site renders a `/work` page
  from the GitHub API, linking each ticket to its change folder, scenarios, PR
  and CI status.
- **`jak` CLI is the factory's tool; MCP is an optional public showcase.** Both
  wrap the same `internal/` packages.

See `docs/roadmap.md` for phases and rationale.
