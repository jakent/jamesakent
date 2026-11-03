package content_test

import (
	"strings"
	"testing"

	sitecontent "github.com/jakent/jamesakent/content"
	"github.com/jakent/jamesakent/internal/content"
)

func TestEmbeddedProjectsAreValid(t *testing.T) {
	ps, err := content.ParseProjects(sitecontent.Projects)
	if err != nil {
		t.Fatal(err)
	}
	if len(ps) == 0 {
		t.Fatal("no projects")
	}
}

func TestValidateRejects(t *testing.T) {
	ok := content.Project{ID: "a", Title: "A", SourceURL: "https://github.com/jakent/a"}
	with := func(f func(*content.Project)) []content.Project { p := ok; f(&p); return []content.Project{p} }

	cases := map[string]struct {
		in   []content.Project
		want string
	}{
		"bad slug":       {with(func(p *content.Project) { p.ID = "Not A Slug" }), "slug"},
		"duplicate id":   {[]content.Project{ok, ok}, "duplicate"},
		"no title":       {with(func(p *content.Project) { p.Title = " " }), "title"},
		"no source":      {with(func(p *content.Project) { p.SourceURL = "" }), "source_url is required"},
		"http source":    {with(func(p *content.Project) { p.SourceURL = "http://x.io/a" }), "https"},
		"foreign demo":   {with(func(p *content.Project) { p.DemoURL = "https://example.com/demo" }), "jamesakent.com"},
		"lookalike demo": {with(func(p *content.Project) { p.DemoURL = "https://eviljamesakent.com" }), "jamesakent.com"},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := content.Validate(c.in)
			if err == nil || !strings.Contains(err.Error(), c.want) {
				t.Fatalf("want error containing %q, got %v", c.want, err)
			}
		})
	}
}

func TestValidateAcceptsSubdomainAndNoDemo(t *testing.T) {
	ps := []content.Project{
		{ID: "a", Title: "A", SourceURL: "https://github.com/jakent/a", DemoURL: "https://life.jamesakent.com"},
		{ID: "b", Title: "B", SourceURL: "https://github.com/jakent/b"},
	}
	if err := content.Validate(ps); err != nil {
		t.Fatal(err)
	}
}

func TestFilter(t *testing.T) {
	ps := []content.Project{{ID: "a", Tags: []string{"x"}}, {ID: "b", Tags: []string{"y"}}}
	if got := content.Filter(ps, "x"); len(got) != 1 || got[0].ID != "a" {
		t.Fatalf("got %v", got)
	}
	if got := content.Filter(ps, ""); len(got) != 2 {
		t.Fatalf("empty tag should return all, got %v", got)
	}
}
