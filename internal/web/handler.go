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
<h2>Projects</h2>
{{range .}}<article class="project" data-id="{{.ID}}" data-tags="{{range $i, $t := .Tags}}{{if $i}} {{end}}{{$t}}{{end}}">
  <h3 class="title">{{.Title}}</h3>
  <p class="summary">{{.Summary}}</p>
  <a class="source" href="{{.SourceURL}}">Source</a>
  {{if .DemoURL}}<a class="demo" href="{{.DemoURL}}">Live demo</a>{{else}}<span class="no-demo">No hosted demo</span>{{end}}
</article>
{{end}}`))

// Handler returns the site's routes. Every response is marked with
// X-Served-By so the acceptance suite can tell the site from other upstreams
// behind the reverse proxy.
func Handler(projects []content.Project) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		page.Execute(w, content.Filter(projects, r.URL.Query().Get("tag")))
	})
	mux.HandleFunc("GET /api/projects", func(w http.ResponseWriter, r *http.Request) {
		list := content.Filter(projects, r.URL.Query().Get("tag"))
		if list == nil {
			list = []content.Project{}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"projects": list})
	})
	return servedBy("site", mux)
}

func servedBy(name string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Served-By", name)
		next.ServeHTTP(w, r)
	})
}
