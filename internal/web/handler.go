// Package web holds the site's HTTP handlers.
package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jakent/jamesakent/internal/content"
)

var page = template.Must(template.New("home").Parse(`<!doctype html>
<meta charset="utf-8">
<title>James Kent</title>
<h1>James Kent</h1>
<nav><a href="/about">About</a></nav>
<h2>Projects</h2>
{{range .}}<article class="project" data-id="{{.ID}}" data-tags="{{range $i, $t := .Tags}}{{if $i}} {{end}}{{$t}}{{end}}">
  <h3 class="title">{{.Title}}</h3>
  <p class="summary">{{.Summary}}</p>
  <a class="source" href="{{.SourceURL}}">Source</a>
  {{if .DemoURL}}<a class="demo" href="{{.DemoURL}}">Live demo</a>{{else}}<span class="no-demo">No hosted demo</span>{{end}}
</article>
{{end}}`))

var aboutPage = template.Must(template.New("about").Parse(`<!doctype html>
<meta charset="utf-8">
<title>About · {{.Name}}</title>
<nav><a href="/">Home</a></nav>
<h1 class="name">{{.Name}}</h1>
<p class="headline">{{.Headline}}</p>
<h2>Now</h2>
<p class="now">{{.Now}}</p>
<h2>Contact</h2>
<ul class="contact">
{{range .Contact}}  <li><a href="{{.URL}}">{{.Label}}</a></li>
{{end}}</ul>
`))

// Handler returns the site's routes. Every response is marked with
// X-Served-By so the acceptance suite can tell the site from other upstreams
// behind the reverse proxy.
func Handler(projects []content.Project, profile content.Profile) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		page.Execute(w, content.Filter(projects, r.URL.Query().Get("tag")))
	})
	mux.HandleFunc("GET /about", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		aboutPage.Execute(w, profile)
	})
	mux.HandleFunc("GET /api/projects", func(w http.ResponseWriter, r *http.Request) {
		list := content.Filter(projects, r.URL.Query().Get("tag"))
		if list == nil {
			list = []content.Project{}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"projects": list})
	})
	mux.HandleFunc("GET /api/profile", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(profile)
	})
	return servedBy("site", mux)
}

func servedBy(name string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Served-By", name)
		next.ServeHTTP(w, r)
	})
}
