# Entity model

Nouns, invariants and states. Behavior (what visitors can do) lives in
`specs/features/`; this file only says what things are.

## Project
Something James built that the site showcases.

| Field | Meaning |
|---|---|
| `id` | Stable slug, unique across projects (`game-of-three-d`) |
| `title` | Display name |
| `summary` | One or two sentences |
| `tags` | Free-form labels (`clojure`, `clojurescript`, `game-of-life`, `htmx`) |
| `source_url` | Source repository. **Required.** |
| `demo_url` | Live demo. Optional |
| `kind` | `demo` (runs in the browser), `app` (a hosted server app), `library`, `desktop` (no hosted demo) |

Invariants
- `id` is unique and URL-safe.
- Every project has a `source_url`.
- A project of kind `demo` or `app` has a `demo_url`; `library` and `desktop` projects do not.
- `demo_url` is on the jamesakent.com domain or one of its subdomains.

Source of truth: `content/projects.yaml`.
