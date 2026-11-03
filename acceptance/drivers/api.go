package drivers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// API reads projects from the JSON API. Page requests use plain HTTP.
type API struct{ *HTTP }

func NewAPI(baseURL string) (*API, error) {
	h, err := NewHTTP(baseURL)
	if err != nil {
		return nil, err
	}
	return &API{h}, nil
}

func (d *API) ListProjects(ctx context.Context, tag string) ([]Project, error) {
	res, err := d.Get(ctx, "", "/api/projects"+tagQuery(tag))
	if err != nil {
		return nil, err
	}
	if res.Status != 200 {
		return nil, fmt.Errorf("GET /api/projects: status %d", res.Status)
	}
	var body struct {
		Projects []struct {
			Title     string   `json:"title"`
			SourceURL string   `json:"source_url"`
			DemoURL   string   `json:"demo_url"`
			Tags      []string `json:"tags"`
		} `json:"projects"`
	}
	if err := json.Unmarshal([]byte(res.Body), &body); err != nil {
		return nil, err
	}
	out := make([]Project, 0, len(body.Projects))
	for _, p := range body.Projects {
		out = append(out, Project{p.Title, p.SourceURL, p.DemoURL, p.Tags})
	}
	return out, nil
}

// GetProfile reads GET /api/profile:
// {"name","headline","now","contact":[{"label","url"}]}.
func (d *API) GetProfile(ctx context.Context) (Profile, error) {
	res, err := d.Get(ctx, "", "/api/profile")
	if err != nil {
		return Profile{}, err
	}
	if res.Status != 200 {
		return Profile{}, fmt.Errorf("GET /api/profile: status %d", res.Status)
	}
	var body struct {
		Name     string `json:"name"`
		Headline string `json:"headline"`
		Now      string `json:"now"`
		Contact  []struct {
			Label string `json:"label"`
			URL   string `json:"url"`
		} `json:"contact"`
	}
	if err := json.Unmarshal([]byte(res.Body), &body); err != nil {
		return Profile{}, err
	}
	p := Profile{Name: body.Name, Headline: body.Headline, Now: body.Now}
	for _, c := range body.Contact {
		p.Contact = append(p.Contact, Contact{c.Label, c.URL})
	}
	return p, nil
}

func tagQuery(tag string) string {
	if tag == "" {
		return ""
	}
	return "?tag=" + url.QueryEscape(tag)
}
