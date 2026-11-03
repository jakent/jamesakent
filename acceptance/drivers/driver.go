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

// Driver opens pages. Host is a subdomain label ("" is the bare domain,
// "www" is www.<domain>).
type Driver interface {
	Get(ctx context.Context, host, path string) (Response, error)
}
