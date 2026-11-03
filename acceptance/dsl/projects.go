package dsl

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jakent/jamesakent/acceptance/drivers"
	"github.com/jakent/jamesakent/acceptance/support"
)

func ListProjects(ctx context.Context, w *support.World) error {
	return ListProjectsTagged(ctx, w, "")
}

func ListProjectsTagged(ctx context.Context, w *support.World, tag string) error {
	ps, err := w.Driver.ListProjects(ctx, tag)
	if err != nil {
		return err
	}
	w.Projects = ps
	return nil
}

func EveryProjectLinksToItsSource(w *support.World) error {
	if len(w.Projects) == 0 {
		return fmt.Errorf("no projects were listed")
	}
	for _, p := range w.Projects {
		if p.SourceURL == "" {
			return fmt.Errorf("project %q has no source link", p.Title)
		}
	}
	return nil
}

func ProjectHasLiveDemo(w *support.World, title string) error {
	p, err := find(w, title)
	if err != nil {
		return err
	}
	if p.DemoURL == "" {
		return fmt.Errorf("project %q has no live demo", title)
	}
	return nil
}

func ProjectHasNoLiveDemo(w *support.World, title string) error {
	p, err := find(w, title)
	if err != nil {
		return err
	}
	if p.DemoURL != "" {
		return fmt.Errorf("project %q should have no live demo but links to %q", title, p.DemoURL)
	}
	return nil
}

func ProjectHasLiveDemoAt(w *support.World, title, path string) error {
	if err := ProjectHasLiveDemo(w, title); err != nil {
		return err
	}
	p, _ := find(w, title)
	u, err := url.Parse(p.DemoURL)
	if err != nil {
		return err
	}
	if u.Path != path {
		return fmt.Errorf("project %q demo is at %q, want path %q", title, p.DemoURL, path)
	}
	return nil
}

func EveryListedProjectIsTagged(w *support.World, tag string) error {
	if len(w.Projects) == 0 {
		return fmt.Errorf("no projects were listed")
	}
	for _, p := range w.Projects {
		if !hasTag(p, tag) {
			return fmt.Errorf("project %q is not tagged %q (tags: %v)", p.Title, tag, p.Tags)
		}
	}
	return nil
}

func ProjectIsNotListed(w *support.World, title string) error {
	if _, err := find(w, title); err == nil {
		return fmt.Errorf("project %q should not be listed", title)
	}
	return nil
}

func find(w *support.World, title string) (drivers.Project, error) {
	for _, p := range w.Projects {
		if p.Title == title {
			return p, nil
		}
	}
	return drivers.Project{}, fmt.Errorf("project %q is not listed", title)
}

func hasTag(p drivers.Project, tag string) bool {
	for _, t := range p.Tags {
		if t == tag {
			return true
		}
	}
	return false
}
