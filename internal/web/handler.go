// Package web holds the site's HTTP handlers.
package web

import (
	"fmt"
	"net/http"
)

// Handler returns the site's routes. Every response is marked with
// X-Served-By so the acceptance suite can tell the site from other upstreams
// behind the reverse proxy.
func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<!doctype html><title>James Kent</title><h1>James Kent</h1>")
	})
	return servedBy("site", mux)
}

func servedBy(name string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Served-By", name)
		next.ServeHTTP(w, r)
	})
}
