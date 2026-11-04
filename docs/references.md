# References

Where the ideas in this repo come from. Statements marked *(not verified)* were
read in search results or summaries, not in the source itself.

## Executable specifications (the four-layer model)
**Dave Farley, "The Industry Is Catching Up: The Only Way to Generate AI Code
That Works"** (Modern Software Engineering channel).
https://www.youtube.com/watch?v=EkH_uoybj8k&t=527s (the executable-spec
section starts around 8:47)

The model this repo follows: test cases in Given/When/Then, written in the
language of the problem domain -> a DSL -> protocol drivers -> the system under
test. Test cases never mention how the system is built, so the implementation
(or the way in) can change underneath them. In this repo that is
`specs/features/` -> `acceptance/dsl/` -> `acceptance/drivers/` -> the site.
Other Farley ideas used here: functional isolation by aliasing names, time as an
external dependency, and human-owned holdout scenarios the agent cannot edit.

## Farley's example repos (GitHub account `davef77`)
- **Flight-Search-ATDD** (Python): "a demonstration of prompting an AI to write
  code with only acceptance tests". Closest to our factory idea. Contains
  `sample_prompt.txt`, `todo.md`, `flight_search/`, `test/` and test-run logs
  from different stages. Only the top-level listing was read.
  https://github.com/davef77/Flight-Search-ATDD
- **acceptance-testing** (Java): sample code for a DSL-based approach, with a
  Selenium script, so a UI driver exists. A Cucumber example is mentioned in a
  video description *(not verified)*. https://github.com/davef77/acceptance-testing
- **atdd-course-examples** (Java): from his "ATDD & BDD: Stories to Exec Specs"
  course. https://github.com/davef77/atdd-course-examples
- **ATDD-Tools** (Java): helpers for building acceptance-test DSLs.
  https://github.com/davef77/ATDD-Tools

## Community examples
- Learn Go with Tests, "Scaling acceptance tests" (Go; the same layering with
  an HTTP driver): https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/scaling-acceptance-tests
- computerguycj/order-shop (C#/NUnit; its README says it is the author's
  illustration, not Farley's course code): https://github.com/computerguycj/order-shop
- dawid-dahl-umain/augmented-ai-development-demo (TypeScript; DSL plus a
  protocol driver in an AI-assisted workflow): https://github.com/dawid-dahl-umain/augmented-ai-development-demo
- LMAX Simple-DSL (lineage): https://github.com/LMAX-Exchange/Simple-DSL
- Farley's ATDD guide: https://continuous-delivery.co.uk/downloads/ATDD%20Guide%2026-03-21.pdf

## To do
- Read Flight-Search-ATDD's prompt and `test/` layout and compare with the
  `change` -> `write-feature` -> `implement-steps` -> `implement` skills.
- Check whether Farley's Claude Code ATDD skill (mentioned in a third-party
  LinkedIn post) exists and is public *(not verified)*.
