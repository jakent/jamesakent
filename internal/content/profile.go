package content

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// Contact is one way to reach the site's owner.
type Contact struct {
	Label string `yaml:"label" json:"label"`
	URL   string `yaml:"url" json:"url"`
}

// Profile is who the site is about. There is exactly one.
type Profile struct {
	Name     string    `yaml:"name" json:"name"`
	Headline string    `yaml:"headline" json:"headline"`
	Now      string    `yaml:"now" json:"now"`
	Contact  []Contact `yaml:"contact" json:"contact"`
}

// ParseProfile decodes profile.yaml and validates every invariant.
func ParseProfile(data []byte) (Profile, error) {
	var p Profile
	if err := yaml.Unmarshal(data, &p); err != nil {
		return Profile{}, fmt.Errorf("parse profile: %w", err)
	}
	if err := ValidateProfile(p); err != nil {
		return Profile{}, err
	}
	return p, nil
}

// ValidateProfile checks the Profile invariants: name, headline and now are
// set, there is at least one contact link, and every contact has a label and
// an https or mailto URL.
func ValidateProfile(p Profile) error {
	for field, v := range map[string]string{"name": p.Name, "headline": p.Headline, "now": p.Now} {
		if strings.TrimSpace(v) == "" {
			return fmt.Errorf("profile: %s is required", field)
		}
	}
	if len(p.Contact) == 0 {
		return fmt.Errorf("profile: at least one contact link is required")
	}
	for i, c := range p.Contact {
		if strings.TrimSpace(c.Label) == "" {
			return fmt.Errorf("profile: contact %d: label is required", i+1)
		}
		if !strings.HasPrefix(c.URL, "https://") && !strings.HasPrefix(c.URL, "mailto:") {
			return fmt.Errorf("profile: contact %q: url must be https or mailto", c.Label)
		}
	}
	return nil
}
