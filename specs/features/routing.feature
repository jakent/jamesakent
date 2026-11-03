@req:SITE-001 @smoke
Feature: Routing
  jamesakent.com serves the personal site at the root, keeps the Clojure REPL
  at /repl, and sends www visitors to the bare domain.

  Rule: The root is the personal site, not the REPL

    Scenario: A visitor opens the home page
      When a visitor opens "/"
      Then the response is successful
      And the page is served by the site

  Rule: The REPL stays available at /repl

    @stack
    Scenario: A visitor opens the REPL
      When a visitor opens "/repl"
      Then the response is successful
      And the page is served by the REPL

  Rule: www redirects to the bare domain

    @stack
    Scenario: A visitor opens the www host
      When a visitor opens "/" on the "www" host
      Then they are redirected permanently to the bare domain
