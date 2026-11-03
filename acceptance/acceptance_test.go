package acceptance

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cucumber/godog"

	"github.com/jakent/jamesakent/acceptance/drivers"
	"github.com/jakent/jamesakent/acceptance/steps"
	"github.com/jakent/jamesakent/internal/content"
	"github.com/jakent/jamesakent/internal/web"
)

// TestFeatures runs specs/features once per driver. With BASE_URL set it
// drives that running system (the full stack, including @stack scenarios).
// Without it, it starts the site in-process on fixture content and skips
// scenarios that need other services.
func TestFeatures(t *testing.T) {
	baseURL := os.Getenv("BASE_URL")
	tags := os.Getenv("GODOG_TAGS")
	if baseURL == "" {
		data, err := os.ReadFile("fixtures/projects.yaml")
		if err != nil {
			t.Fatal(err)
		}
		projects, err := content.ParseProjects(data)
		if err != nil {
			t.Fatal(err)
		}
		pdata, err := os.ReadFile("fixtures/profile.yaml")
		if err != nil {
			t.Fatal(err)
		}
		profile, err := content.ParseProfile(pdata)
		if err != nil {
			t.Fatal(err)
		}
		srv := httptest.NewServer(web.Handler(projects, profile))
		defer srv.Close()
		baseURL = srv.URL
		if tags == "" {
			tags = "~@stack"
		}
	}

	api, err := drivers.NewAPI(baseURL)
	if err != nil {
		t.Fatal(err)
	}
	web, err := drivers.NewWeb(baseURL)
	if err != nil {
		t.Fatal(err)
	}

	for name, driver := range map[string]interface {
		drivers.Driver
		Domain() string
	}{"api": api, "web": web} {
		t.Run(name, func(t *testing.T) {
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
		})
	}
}
