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
	sc.Step(`^they are redirected permanently to the bare domain$`, func() error {
		return dsl.RedirectedPermanentlyToBareDomain(w)
	})
}
