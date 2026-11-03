---
name: change
description: Start a new change in this repo. Scaffolds specs/changes/<date>-<name>/ with a proposal (including the required Testability section) and opens the linked GitHub issue. Use when the owner wants a new feature, fix or content change, before any spec or code is written.
---

# change

Every unit of work starts as a change folder plus a GitHub issue. Nothing else
(features, steps, code) is written until the proposal exists.

## Steps
1. Pick a short kebab-case name and today's date: `specs/changes/YYYY-MM-DD-<name>/`.
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
5. Open the issue: `gh issue create --title "<title>" --body "Change: specs/changes/<folder>/\nCovers: <@req ids, once known>"`.
   Put the issue number in the proposal's first line.
6. Do not write the feature or any code. Hand off to `write-feature`.

## Done when
The folder and issue exist, open questions are listed, and the Testability
section is filled in.
