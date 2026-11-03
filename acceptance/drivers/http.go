package drivers

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// HTTP drives the system with plain HTTP requests against BaseURL. Redirects
// are not followed, so scenarios can assert on them.
type HTTP struct {
	baseURL *url.URL
	client  *http.Client
}

func NewHTTP(baseURL string) (*HTTP, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &HTTP{
		baseURL: u,
		client: &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		}},
	}, nil
}

// Domain is the bare host name that subdomain labels are added to.
func (d *HTTP) Domain() string { return d.baseURL.Hostname() }

func (d *HTTP) Get(ctx context.Context, host, path string) (Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, d.baseURL.String()+path, nil)
	if err != nil {
		return Response{}, err
	}
	if host != "" {
		req.Host = host + "." + d.Domain()
	}
	res, err := d.client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}
	out := Response{
		Status:   res.StatusCode,
		Header:   res.Header,
		Body:     string(body),
		Location: res.Header.Get("Location"),
	}
	if strings.Contains(res.Header.Get("Content-Type"), "html") {
		out.Links = hrefs(out.Body)
	}
	return out, nil
}

// hrefs returns the href of every anchor in an HTML document.
func hrefs(body string) []string {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return nil
	}
	var out []string
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			if h := attr(n, "href"); h != "" {
				out = append(out, h)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)
	return out
}
