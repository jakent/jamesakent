// Package content defines the site's content types and enforces the
// invariants from specs/model/entities.md.
package content

import (
	"fmt"
	"net/url"
	"regexp"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

// Project is something that is showcased on the site.
type Project struct {
	ID        string   `yaml:"id" json:"id"`
	Title     string   `yaml:"title" json:"title"`
	Summary   string   `yaml:"summary" json:"summary"`
	Tags      []string `yaml:"tags" json:"tags"`
	SourceURL string   `yaml:"source_url" json:"source_url"`
	DemoURL   string   `yaml:"demo_url,omitempty" json:"demo_url,omitempty"`
}

// HasTag reports whether the project carries the tag.
func (p Project) HasTag(tag string) bool { return slices.Contains(p.Tags, tag) }

const siteDomain = "jamesakent.com"

var slug = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// ParseProjects decodes projects.yaml and validates every invariant.
func ParseProjects(data []byte) ([]Project, error) {
	var ps []Project
	if err := yaml.Unmarshal(data, &ps); err != nil {
		return nil, fmt.Errorf("parse projects: %w", err)
	}
	if err := Validate(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Validate checks the Project invariants: unique URL-safe ids, a title, a
// source repository, and a demo (when present) on the site's own domain.
func Validate(ps []Project) error {
	seen := map[string]bool{}
	for i, p := range ps {
		where := fmt.Sprintf("project %d (%q)", i+1, p.ID)
		switch {
		case !slug.MatchString(p.ID):
			return fmt.Errorf("%s: id must be a lowercase URL-safe slug", where)
		case seen[p.ID]:
			return fmt.Errorf("%s: duplicate id", where)
		case strings.TrimSpace(p.Title) == "":
			return fmt.Errorf("%s: title is required", where)
		case p.SourceURL == "":
			return fmt.Errorf("%s: source_url is required", where)
		}
		seen[p.ID] = true
		if err := checkURL(p.SourceURL, false); err != nil {
			return fmt.Errorf("%s: source_url: %w", where, err)
		}
		if p.DemoURL != "" {
			if err := checkURL(p.DemoURL, true); err != nil {
				return fmt.Errorf("%s: demo_url: %w", where, err)
			}
		}
	}
	return nil
}

func checkURL(raw string, onSiteDomain bool) error {
	u, err := url.Parse(raw)
	if err != nil || u.Scheme != "https" || u.Host == "" {
		return fmt.Errorf("%q must be an absolute https URL", raw)
	}
	if onSiteDomain {
		h := u.Hostname()
		if h != siteDomain && !strings.HasSuffix(h, "."+siteDomain) {
			return fmt.Errorf("%q must be on %s or a subdomain", raw, siteDomain)
		}
	}
	return nil
}

// Filter returns the projects carrying tag; an empty tag returns all of them.
func Filter(ps []Project, tag string) []Project {
	if tag == "" {
		return ps
	}
	var out []Project
	for _, p := range ps {
		if p.HasTag(tag) {
			out = append(out, p)
		}
	}
	return out
}
