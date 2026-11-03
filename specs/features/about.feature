@req:SITE-003
Feature: About
  Visitors can find out who the site is about, what James is working on now,
  and how to get in touch.

  Rule: The About page says who and what

    Scenario: A visitor reads the About page
      When a visitor reads the About page
      Then the profile has a name and a headline
      And the profile says what James is working on now

  Rule: Visitors can reach James through valid links

    Scenario: A visitor finds a way to get in touch
      When a visitor reads the About page
      Then there is at least one contact link
      And every contact link is a secure web link or an email link

  Rule: The About page is reachable from the home page

    Scenario: A visitor navigates from the home page to the About page
      When a visitor opens "/"
      Then the page links to the About page
