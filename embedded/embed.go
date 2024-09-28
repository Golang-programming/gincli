// ./embedded/embed.go
package embedded

import (
	"embed"
)

//go:embed templates/**/*
var TemplatesFS embed.FS
