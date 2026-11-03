// Package content embeds the site's content files into the binary.
package content

import _ "embed"

// Projects is the raw content/projects.yaml.
//
//go:embed projects.yaml
var Projects []byte

// Profile is the raw content/profile.yaml.
//
//go:embed profile.yaml
var Profile []byte
