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
| `demo_url` | Live demo. Optional. A project has a live demo exactly when this is set |

Invariants
- `id` is unique and URL-safe.
- Every project has a `source_url`.
- `demo_url` is on the jamesakent.com domain or one of its subdomains.

Source of truth: `content/projects.yaml`.
