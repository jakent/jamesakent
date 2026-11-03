---
name: implement-steps
description: Turn an approved feature file into runnable acceptance tests: step definitions, DSL functions and driver methods, with no application code. Use after the owner has approved the feature. Leaves the new scenarios failing for the right reason.
---

# implement-steps

Build the test layers for an approved feature so the scenarios run and fail
because the product does not do the thing yet. Do not touch app code and do not
edit `specs/features/`.

## Layers (from AGENTS.md)
`.feature` -> steps (`acceptance/steps/`, one line each) -> DSL
(`acceptance/dsl/`, plain functions, World first) -> drivers
(`acceptance/drivers/`, one `Driver` interface, implemented per way in) -> system.

## Steps
1. Run `make acceptance`; the undefined steps are your to-do list. Copy the
   suggested regexes exactly.
2. For each step: add a one-line `sc.Step(...)` in `steps.go` that only calls a
   DSL function. Business logic in a step is a bug.
3. Add DSL functions in `acceptance/dsl/`. They talk only to `w.Driver` and
   `w.<state>`; they never know which driver is in use, never hardcode URLs.
   Error messages must name the project/entity and say what was expected.
4. If the DSL needs something new from the system, extend the `Driver`
   interface and implement it in **every** driver (`api`, `web`, ...). Keep the
   shared DTO in `drivers.go` in the vocabulary of the model, not of HTTP.
5. If new content or data is needed, add it to `acceptance/fixtures/` (the
   in-process run uses fixtures, not real content).
6. Run `make acceptance`. The new scenarios must now be **red because the
   behavior is missing** (404, wrong data, absent field), not because of a
   missing step, panic, or a typo in a regex. If red for the wrong reason, fix
   the test layer.
7. `gofmt -l .` and `go vet ./...` must be clean. Commit on the change branch.

## Done when
Every step is defined, scenarios fail with a clear product-level message, and
no file outside `acceptance/` changed (except fixtures).
