package acceptance

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"net/http/httptest"

	"github.com/jakent/jamesakent/acceptance/drivers"
	"github.com/jakent/jamesakent/acceptance/steps"
	"github.com/jakent/jamesakent/internal/web"
)

// TestFeatures runs specs/features. With BASE_URL set it drives that running
// system (the full stack, including @stack scenarios). Without it, it starts
// the site in-process and skips scenarios that need other services.
func TestFeatures(t *testing.T) {
	baseURL := os.Getenv("BASE_URL")
	tags := os.Getenv("GODOG_TAGS")
	if baseURL == "" {
		srv := httptest.NewServer(web.Handler())
		defer srv.Close()
		baseURL = srv.URL
		if tags == "" {
			tags = "~@stack"
		}
	}

	driver, err := drivers.NewHTTP(baseURL)
	if err != nil {
		t.Fatal(err)
	}

	suite := godog.TestSuite{
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			steps.Register(sc, driver, driver.Domain())
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../specs/features"},
			Tags:     tags,
			Strict:   true,
			TestingT: t,
		},
	}
	if suite.Run() != 0 {
		t.Fatal("acceptance scenarios failed")
	}
}
