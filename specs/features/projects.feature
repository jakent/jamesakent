@req:SITE-002
Feature: Projects
  Visitors can see what James has built, and reach each project's source and,
  when there is one, its live demo.

  Rule: Every project links to its source repository

    Scenario: A visitor lists the projects
      When a visitor lists the projects
      Then every project links to its source repository

  Rule: The Game of Life demos are findable and live

    Scenario: A visitor finds the Game of Life demos
      When a visitor lists the projects
      Then "Game of Three-D" has a live demo
      And "Game of Three" has a live demo

  Rule: Projects without a hosted demo say so instead of showing a broken link

    Scenario: A desktop-only project has no demo link
      When a visitor lists the projects
      Then "Game of Clojure" has no live demo

  Rule: The REPL is listed as a project that opens at /repl

    Scenario: A visitor finds the REPL
      When a visitor lists the projects
      Then "Clojure REPL" has a live demo at "/repl"

  Rule: Projects can be filtered by tag

    Scenario: A visitor filters by tag
      When a visitor lists the projects tagged "game-of-life"
      Then every listed project is tagged "game-of-life"
      And "Clojure REPL" is not listed
