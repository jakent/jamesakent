// Package dsl is the driver-agnostic vocabulary scenarios are written in.
// Functions take the World first; step definitions only call these.
package dsl

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jakent/jamesakent/acceptance/support"
)

func OpenPage(ctx context.Context, w *support.World, host, path string) error {
	res, err := w.Driver.Get(ctx, host, path)
	if err != nil {
		return err
	}
	w.Last = res
	return nil
}

func ResponseIsSuccessful(w *support.World) error {
	if s := w.Last.Status; s < 200 || s > 299 {
		return fmt.Errorf("expected a successful response, got status %d", s)
	}
	return nil
}

// PageIsServedBySite checks the marker the site adds to every response.
func PageIsServedBySite(w *support.World) error {
	if got := w.Last.Header.Get("X-Served-By"); got != "site" {
		return fmt.Errorf("expected X-Served-By: site, got %q", got)
	}
	return nil
}

// PageIsServedByRepl is provisional: the REPL does not set a marker, so it is
// identified as "a successful page the site did not serve".
func PageIsServedByRepl(w *support.World) error {
	if got := w.Last.Header.Get("X-Served-By"); got == "site" {
		return fmt.Errorf("expected the REPL, but the site served this page")
	}
	return nil
}

func RedirectedPermanentlyToBareDomain(w *support.World) error {
	if s := w.Last.Status; s != 301 && s != 308 {
		return fmt.Errorf("expected a permanent redirect, got status %d", s)
	}
	loc, err := url.Parse(w.Last.Location)
	if err != nil {
		return err
	}
	if loc.Hostname() != w.Domain {
		return fmt.Errorf("expected redirect to %q, got %q", w.Domain, w.Last.Location)
	}
	return nil
}
