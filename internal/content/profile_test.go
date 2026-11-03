package content_test

import (
	"strings"
	"testing"

	sitecontent "github.com/jakent/jamesakent/content"
	"github.com/jakent/jamesakent/internal/content"
)

func TestEmbeddedProfileIsValid(t *testing.T) {
	if _, err := content.ParseProfile(sitecontent.Profile); err != nil {
		t.Fatal(err)
	}
}

func TestValidateProfileRejects(t *testing.T) {
	ok := content.Profile{
		Name: "N", Headline: "H", Now: "W",
		Contact: []content.Contact{{Label: "GitHub", URL: "https://github.com/jakent"}},
	}
	with := func(f func(*content.Profile)) content.Profile { p := ok; f(&p); return p }

	cases := map[string]struct {
		in   content.Profile
		want string
	}{
		"no name":       {with(func(p *content.Profile) { p.Name = " " }), "name"},
		"no headline":   {with(func(p *content.Profile) { p.Headline = "" }), "headline"},
		"no now":        {with(func(p *content.Profile) { p.Now = "" }), "now"},
		"no contact":    {with(func(p *content.Profile) { p.Contact = nil }), "contact"},
		"http contact":  {with(func(p *content.Profile) { p.Contact[0].URL = "http://x.io" }), "https or mailto"},
		"script scheme": {with(func(p *content.Profile) { p.Contact = []content.Contact{{"x", "javascript:alert(1)"}} }), "https or mailto"},
		"no label":      {with(func(p *content.Profile) { p.Contact = []content.Contact{{"", "https://x.io"}} }), "label"},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := content.ValidateProfile(c.in)
			if err == nil || !strings.Contains(err.Error(), c.want) {
				t.Fatalf("want error containing %q, got %v", c.want, err)
			}
		})
	}
}

func TestValidateProfileAcceptsMailto(t *testing.T) {
	p := content.Profile{
		Name: "N", Headline: "H", Now: "W",
		Contact: []content.Contact{{Label: "Email", URL: "mailto:a@b.io"}},
	}
	if err := content.ValidateProfile(p); err != nil {
		t.Fatal(err)
	}
}
