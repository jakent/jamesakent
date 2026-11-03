package drivers

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Web reads projects the way a browser would see them: it fetches the home
// page and parses the rendered HTML. Page requests use plain HTTP.
type Web struct{ *HTTP }

func NewWeb(baseURL string) (*Web, error) {
	h, err := NewHTTP(baseURL)
	if err != nil {
		return nil, err
	}
	return &Web{h}, nil
}

func (d *Web) ListProjects(ctx context.Context, tag string) ([]Project, error) {
	res, err := d.Get(ctx, "", "/"+tagQuery(tag))
	if err != nil {
		return nil, err
	}
	if res.Status != 200 {
		return nil, fmt.Errorf("GET /: status %d", res.Status)
	}
	doc, err := html.Parse(strings.NewReader(res.Body))
	if err != nil {
		return nil, err
	}
	var out []Project
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "article" && hasClass(n, "project") {
			out = append(out, parseProject(n))
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)
	return out, nil
}

func parseProject(article *html.Node) Project {
	p := Project{Tags: strings.Fields(attr(article, "data-tags"))}
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch {
			case n.Data == "h3" && hasClass(n, "title"):
				p.Title = text(n)
			case n.Data == "a" && hasClass(n, "source"):
				p.SourceURL = attr(n, "href")
			case n.Data == "a" && hasClass(n, "demo"):
				p.DemoURL = attr(n, "href")
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(article)
	return p
}

func attr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

func hasClass(n *html.Node, class string) bool {
	for _, c := range strings.Fields(attr(n, "class")) {
		if c == class {
			return true
		}
	}
	return false
}

func text(n *html.Node) string {
	var b strings.Builder
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.TextNode {
			b.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return strings.TrimSpace(b.String())
}

// GetProfile reads the About page. Markup contract the page must follow:
//
//	<h1 class="name">…</h1> <p class="headline">…</p> <p class="now">…</p>
//	<ul class="contact"><li><a href="…">label</a></li>…</ul>
func (d *Web) GetProfile(ctx context.Context) (Profile, error) {
	res, err := d.Get(ctx, "", "/about")
	if err != nil {
		return Profile{}, err
	}
	if res.Status != 200 {
		return Profile{}, fmt.Errorf("GET /about: status %d", res.Status)
	}
	doc, err := html.Parse(strings.NewReader(res.Body))
	if err != nil {
		return Profile{}, err
	}
	var p Profile
	var inContact bool
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch {
			case n.Data == "h1" && hasClass(n, "name"):
				p.Name = text(n)
			case n.Data == "p" && hasClass(n, "headline"):
				p.Headline = text(n)
			case n.Data == "p" && hasClass(n, "now"):
				p.Now = text(n)
			case n.Data == "ul" && hasClass(n, "contact"):
				inContact = true
				defer func() { inContact = false }()
			case n.Data == "a" && inContact:
				p.Contact = append(p.Contact, Contact{Label: text(n), URL: attr(n, "href")})
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)
	return p, nil
}
