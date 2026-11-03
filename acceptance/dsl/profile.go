package dsl

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/jakent/jamesakent/acceptance/support"
)

// aboutPath is where the home page is expected to link for the About page.
const aboutPath = "/about"

func ReadAboutPage(ctx context.Context, w *support.World) error {
	p, err := w.Driver.GetProfile(ctx)
	if err != nil {
		return err
	}
	w.Profile = p
	return nil
}

func ProfileHasNameAndHeadline(w *support.World) error {
	if strings.TrimSpace(w.Profile.Name) == "" {
		return fmt.Errorf("the profile has no name")
	}
	if strings.TrimSpace(w.Profile.Headline) == "" {
		return fmt.Errorf("the profile has no headline")
	}
	return nil
}

func ProfileSaysWhatJamesIsWorkingOnNow(w *support.World) error {
	if strings.TrimSpace(w.Profile.Now) == "" {
		return fmt.Errorf("the profile does not say what James is working on now")
	}
	return nil
}

func AtLeastOneContactLink(w *support.World) error {
	if len(w.Profile.Contact) == 0 {
		return fmt.Errorf("the profile has no contact links")
	}
	return nil
}

func EveryContactLinkIsSecureOrEmail(w *support.World) error {
	for _, c := range w.Profile.Contact {
		if !strings.HasPrefix(c.URL, "https://") && !strings.HasPrefix(c.URL, "mailto:") {
			return fmt.Errorf("contact link %q (%s) is neither https nor mailto", c.Label, c.URL)
		}
	}
	return nil
}

func PageLinksToAboutPage(w *support.World) error {
	if !slices.Contains(w.Last.Links, aboutPath) {
		return fmt.Errorf("the page has no link to the About page (%s); links: %v", aboutPath, w.Last.Links)
	}
	return nil
}
