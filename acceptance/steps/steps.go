// Package steps binds Gherkin steps to the DSL. Each step is one line.
package steps

import (
	"context"

	"github.com/cucumber/godog"

	"github.com/jakent/jamesakent/acceptance/drivers"
	"github.com/jakent/jamesakent/acceptance/dsl"
	"github.com/jakent/jamesakent/acceptance/support"
)

// Register wires the step definitions to a fresh World per scenario.
func Register(sc *godog.ScenarioContext, d drivers.Driver, domain string) {
	w := &support.World{Driver: d, Domain: domain}

	sc.Step(`^a visitor opens "([^"]*)"$`, func(ctx context.Context, path string) error {
		return dsl.OpenPage(ctx, w, "", path)
	})
	sc.Step(`^a visitor opens "([^"]*)" on the "([^"]*)" host$`, func(ctx context.Context, path, host string) error {
		return dsl.OpenPage(ctx, w, host, path)
	})
	sc.Step(`^the response is successful$`, func() error { return dsl.ResponseIsSuccessful(w) })
	sc.Step(`^the page is served by the site$`, func() error { return dsl.PageIsServedBySite(w) })
	sc.Step(`^the page is served by the REPL$`, func() error { return dsl.PageIsServedByRepl(w) })
	sc.Step(`^a visitor lists the projects$`, func(ctx context.Context) error { return dsl.ListProjects(ctx, w) })
	sc.Step(`^a visitor lists the projects tagged "([^"]*)"$`, func(ctx context.Context, tag string) error {
		return dsl.ListProjectsTagged(ctx, w, tag)
	})
	sc.Step(`^every project links to its source repository$`, func() error { return dsl.EveryProjectLinksToItsSource(w) })
	sc.Step(`^"([^"]*)" has a live demo$`, func(title string) error { return dsl.ProjectHasLiveDemo(w, title) })
	sc.Step(`^"([^"]*)" has no live demo$`, func(title string) error { return dsl.ProjectHasNoLiveDemo(w, title) })
	sc.Step(`^"([^"]*)" has a live demo at "([^"]*)"$`, func(title, path string) error {
		return dsl.ProjectHasLiveDemoAt(w, title, path)
	})
	sc.Step(`^every listed project is tagged "([^"]*)"$`, func(tag string) error { return dsl.EveryListedProjectIsTagged(w, tag) })
	sc.Step(`^"([^"]*)" is not listed$`, func(title string) error { return dsl.ProjectIsNotListed(w, title) })
	sc.Step(`^a visitor reads the About page$`, func(ctx context.Context) error { return dsl.ReadAboutPage(ctx, w) })
	sc.Step(`^the profile has a name and a headline$`, func() error { return dsl.ProfileHasNameAndHeadline(w) })
	sc.Step(`^the profile says what James is working on now$`, func() error { return dsl.ProfileSaysWhatJamesIsWorkingOnNow(w) })
	sc.Step(`^there is at least one contact link$`, func() error { return dsl.AtLeastOneContactLink(w) })
	sc.Step(`^every contact link is a secure web link or an email link$`, func() error {
		return dsl.EveryContactLinkIsSecureOrEmail(w)
	})
	sc.Step(`^the page links to the About page$`, func() error { return dsl.PageLinksToAboutPage(w) })
	sc.Step(`^they are redirected permanently to the bare domain$`, func() error {
		return dsl.RedirectedPermanentlyToBareDomain(w)
	})
}
