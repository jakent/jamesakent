Issue: #6

# About page

## Why
Visitors land on a list of projects with no context about who built them. The
roadmap's "About/Now" idea needs a first, small version: who this is, what James
is working on, and how to get in touch.

## Scope
- A `Profile` entry in `specs/model/entities.md`: name, headline, current focus
  (the "now"), and contact links.
- `content/profile.yaml` (embedded, validated like `projects.yaml`).
- `GET /about` (HTML) and `GET /api/profile` (JSON); a link to `/about` from the home page.
- `specs/features/about.feature` (new; written by `write-feature`, needs owner approval).

Out of scope: a blog/writing section, a resume download, the REPL's `(about-me)`
command reading this data (later driver), analytics.

## Open questions for the owner
Nothing below is decided; I will not guess any of it.
1. **Text:** what is the headline and the "now" paragraph? (I will not invent a
   biography. Give me the words or approve a placeholder to replace later.)
2. **Contact links:** which should appear (GitHub, LinkedIn, email, other)? Is
   showing an email address acceptable, or should it be a link only?
3. **Employer:** should Chariot Solutions be named? Client names (KAOS, Fanatics)
   stay out per the roadmap.
4. **Route:** a separate `/about` page, or a section on the home page?

## Testability
- Profile content is data loaded through the same kind of interface as projects,
  with fixture content in `acceptance/fixtures/profile.yaml` for in-process runs.
- No time, randomness or external services involved.
- Scenarios run through the `api` and `web` drivers on every PR. The `Driver`
  interface gains `GetProfile`, implemented by both.

## Tasks
1. Owner answers the open questions; record them under **Decisions** below.
2. `write-feature`: draft `about.feature`, update the model and glossary.
3. `implement-steps`: steps, DSL, `GetProfile` on both drivers, fixtures.
4. `implement`: content type and validation, `/about`, `/api/profile`, home-page link.
5. Archive this folder after merge.

## Decisions
_None yet._
