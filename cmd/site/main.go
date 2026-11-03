package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

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

	addr := ":" + envOr("PORT", "8080")
	log.Printf("site listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, otelhttp.NewHandler(web.Handler(), "site")))
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
