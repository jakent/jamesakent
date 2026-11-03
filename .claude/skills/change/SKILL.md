---
name: change
description: Start a new change in this repo. Scaffolds specs/changes/<name>/ with a proposal (including the required Testability section), opens the linked GitHub issue and the root proposal PR. Use when the owner wants a new feature, fix or content change, before any spec or code is written.
---

# change

Every unit of work starts as a change folder, a GitHub issue and a proposal PR.
Nothing else (features, steps, code) is written until the proposal exists.

A change travels as a stack of three PRs, merged top first:
`implement-<name>` -> `spec-<name>` -> `change-<name>` -> `main`. This skill
creates the root of that stack. Intermediate merges need no review; only the
root PR, which finally holds proposal, spec and implementation together, goes
into `main` (see `implement` for the order).

## Steps
1. Pick a short kebab-case name: `specs/changes/<name>/`. No date: git records when
   a change started. If the name already exists in `specs/changes/` or
   `specs/changes/archive/`, append `-2` (or a clearer word) so names stay unique.
2. Read `docs/roadmap.md`, `specs/model/entities.md` and `specs/model/glossary.md`
   so the proposal uses existing vocabulary.
3. Write `proposal.md` with these sections, no others unless needed:
   - **Why**: the need, in a few sentences.
   - **Scope**: what changes (features, model entries, endpoints, content). Say
     what is out of scope.
   - **Open questions for the owner**: anything you would otherwise guess. Ask
     instead of choosing for them; record answers under **Decisions** afterwards.
   - **Testability** (required): content/config behind an interface with fixtures;
     time and randomness injectable; external services behind a fake; which
     drivers run the scenarios. If you cannot answer, the design is not ready.
   - **Tasks**: ordered, each ending in a checkable result.
   Point at features and model entries; never restate scenarios (behavior lives
   only in `.feature` files, ADR 0001).
4. Add `design.md` only if there is a real design choice. Decisions that outlive
   the change become an ADR in `docs/adr/`.
5. Open the issue. Build the body with `printf`, because `\n` inside a
   double-quoted string is not a newline:
   `gh issue create --title "<title>" --body "$(printf 'Change: specs/changes/<name>/\nCovers: @req to be assigned in write-feature\n\n<one-sentence summary>')"`.
   Make the proposal's first line `Issue: #<number>`.
6. Branch `change-<name>`, commit the proposal, push, and open the root PR against
   `main`, titled "Change: <title> (proposal)". Its body links the issue, says
   "Closes #<number>" so the issue closes when the stack lands, and says not to
   merge until the spec and implementation PRs are merged into it.
7. Do not write the feature or any code. Hand off to `write-feature`.

## Done when
The folder, issue and root PR exist, open questions are listed, and the
Testability section is filled in.
