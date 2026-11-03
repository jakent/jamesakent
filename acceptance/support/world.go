// Package support holds the per-scenario state shared by DSL and steps.
package support

import "github.com/jakent/jamesakent/acceptance/drivers"

// World is the scenario context passed first to every DSL function.
type World struct {
	Driver drivers.Driver
	// Domain is the bare domain used to check redirects.
	Domain string
	// Last is the most recent response a scenario received.
	Last drivers.Response
}
