---
name: write-feature
description: Draft or change a Gherkin feature file for an existing change proposal, for the owner to approve. Use after the change skill, before any steps or code. Spec-only PR; never merged without the owner's approval.
---

# write-feature

`specs/features/` is human-owned. You may draft or edit features, but only in a
spec-only PR that CODEOWNERS routes to the owner. Never mix spec and
implementation changes in one PR.

## Steps
1. Read the change's `proposal.md` and its recorded Decisions. If open
   questions remain unanswered, stop and ask the owner.
2. Read existing `specs/features/*.feature`, `acceptance/dsl/`, and
   `acceptance/steps/` so new scenarios reuse existing step wording and DSL
   verbs. Add new wording only when needed.
3. Write the feature:
   - Tag the Feature `@req:SITE-NNN` (next unused number). Add `@smoke` only for
     scenarios that should also run against production; add `@stack` for
     scenarios needing Caddy/biff-repl (they are skipped in-process).
   - Group with `Rule:` (an invariant from `specs/model/entities.md`), then
     `Scenario:` examples, including at least one refusal/negative case per rule.
   - Describe what a visitor or operator sees and does, in glossary terms. No
     URLs, CSS, JSON fields or HTTP verbs in steps: those belong to drivers.
4. Update `specs/model/entities.md` and `glossary.md` if the change adds nouns
   or invariants.
5. Check it: `make acceptance` must run and report the new steps as undefined
   (that is the expected red). It must not report parse errors.
6. Branch `spec-<name>`, commit, push, open a PR titled "Spec: <name> (for owner
   approval)". The body lists the rules added and any wording you were unsure about.
7. Stop. Do not start `implement-steps` until the owner approves the feature.
   If they have approved but not merged, implementation may proceed on a stacked
   branch (see `implement`).

## Done when
The feature parses, its steps are undefined (not broken), and the spec PR is open.
