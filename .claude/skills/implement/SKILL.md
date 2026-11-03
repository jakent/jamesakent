---
name: implement
description: Implement application code until an approved change's scenarios pass through every driver, then open the PR and archive the change folder after merge. Use after implement-steps. Never edits specs/features/.
---

# implement

Make the red scenarios green with the smallest honest implementation, then ship
it through a PR.

## Rules
- Never edit `specs/features/**`. If a scenario seems wrong, stop and ask the
  owner; spec changes go through `write-feature` in their own PR.
- Content is data behind an interface; time, randomness and external calls are
  injectable (the change's Testability section says how).
- No `Co-Authored-By` or "Generated with" lines in commits or PR bodies.

## Steps
1. Read the change's `proposal.md` (Scope, Decisions, Tasks) and the feature.
2. Work task by task. For domain rules from `specs/model/entities.md`, write a
   unit test first (`internal/.../*_test.go`), including rejection cases, then
   the code. Scenarios prove behavior; unit tests pin invariants.
3. Run `make acceptance` after each task. All scenarios must pass through **every
   driver** (`api`, `web`, ...). A pass through only one driver is not done.
4. Prove the tests can fail: temporarily break the behavior (for example drop a
   field from the page) and confirm a scenario fails with a clear message. **Undo
   the break by editing back, never with `git checkout <file>` on uncommitted
   work**; commit or stash first. (A past `git checkout` here threw away an
   uncommitted new handler.)
5. If the change touches routing or the stack, also run `make up` and
   `make acceptance-stack`, then `make down` if you started it.
6. Before pushing: `gofmt -l .` empty, `go vet ./...`, `go test ./... -count=1`.
7. Commit small, imperative messages. Branch `implement-<name>`. Open a PR whose
   body lists scope, what you checked (and how), and known gaps. Say if no
   `specs/features/` file changed.
8. Branch from the spec branch and open this PR stacked with `--base spec-<name>`.
   The spec PR alone has red CI (undefined steps), so it never merges first.
   CI only runs on PRs to `main`, so run the full suite locally and say so.
9. Do not merge. Review and merge are the owner's (agent PRs need their
   approval; branch protection enforces it). State the merge order in the PR
   body, top of the stack first:
   1. merge this PR into `spec-<name>` (no review needed; its base is not `main`);
   2. wait for CI to rerun on the spec PR, which now contains spec and
      implementation and should be green;
   3. merge the spec PR into `main` (needs the owner's code-owner approval for
      `specs/features/`).
   Specs and their implementation always land together.
10. After both merge: move `specs/changes/<folder>/` to
    `specs/changes/archive/` in a small follow-up PR and close the issue.

## Done when
All scenarios pass on every driver, unit tests and vet are green, and the PR is
open with an honest list of gaps.
