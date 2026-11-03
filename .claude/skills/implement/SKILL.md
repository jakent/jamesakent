---
name: implement
description: Implement application code until an approved change's scenarios pass through every driver, archive the change folder in the same PR, and open the top PR of the stack. Use after implement-steps. Never edits specs/features/.
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
   field from the page) and confirm a scenario fails with a clear message. **Commit
   first, then break, then undo with `git checkout <file>`; never run `git
   checkout <file>` on uncommitted work.** (A past `git checkout` here threw away
   an uncommitted new handler.)
5. If the change touches routing or the stack, also run `make up` and
   `make acceptance-stack`, then `make down` if you started it.
6. Before pushing: `gofmt -l .` empty, `go vet ./...`, `go test ./... -count=1`.
7. Last task: move `specs/changes/<name>/` to `specs/changes/archive/<name>/` in
   this PR, so the archive lands with the change and needs no follow-up PR.
8. Commit small, imperative messages on `implement-<name>` (created from
   `spec-<name>`). Open the PR with `--base spec-<name>`; its body lists scope,
   what you checked (and how), what you did not check, and known gaps. Say if no
   `specs/features/` file changed. CI only runs on PRs to `main`, so run the
   full suite locally and say so.
9. Do not merge. Review and merge are the owner's (agent PRs need their
   approval; branch protection enforces it). State the merge order in the PR
   body, top of the stack first:
   1. this PR into `spec-<name>` (no review needed; its base is not `main`);
   2. the spec PR into `change-<name>` (no review needed);
   3. wait for CI on the root PR (`change-<name>` -> `main`), which now holds
      proposal, spec and implementation and should be green, then merge it
      into `main`. It needs the owner's code-owner approval for
      `specs/features/`; an owner who authored it uses the admin bypass.
   The three PRs always land together; the root PR's "Closes #<issue>" closes
   the issue.

## Done when
All scenarios pass on every driver, unit tests and vet are green, and the PR is
open with an honest list of gaps.
