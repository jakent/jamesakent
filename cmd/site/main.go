package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	sitecontent "github.com/jakent/jamesakent/content"
	"github.com/jakent/jamesakent/internal/content"
	"github.com/jakent/jamesakent/internal/telemetry"
	"github.com/jakent/jamesakent/internal/web"
)

func main() {
	ctx := context.Background()
	shutdown, err := telemetry.Setup(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer shutdown(ctx)

	projects, err := content.ParseProjects(sitecontent.Projects)
	if err != nil {
		log.Fatal(err)
	}

	profile, err := content.ParseProfile(sitecontent.Profile)
	if err != nil {
		log.Fatal(err)
	}

	addr := ":" + envOr("PORT", "8080")
	log.Printf("site listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, otelhttp.NewHandler(web.Handler(projects, profile), "site")))
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
