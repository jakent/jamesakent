// Package drivers are the ways into the system under test. Each driver
// implements Driver; the DSL never knows which one it is using.
package drivers

import (
	"context"
	"net/http"
)

// Response is what any driver reports back for a page request.
type Response struct {
	Status   int
	Header   http.Header
	Body     string
	Location string
}

// Project is a project as a visitor sees it, whichever way they look.
type Project struct {
	Title     string
	SourceURL string
	DemoURL   string
	Tags      []string
}

// Driver opens pages and lists projects. Host is a subdomain label ("" is
// the bare domain, "www" is www.<domain>). Tag "" means no filter.
type Driver interface {
	Get(ctx context.Context, host, path string) (Response, error)
	ListProjects(ctx context.Context, tag string) ([]Project, error)
}
