package persistence

import "embed"

//go:embed "migrations"
var EmbeddedFiles embed.FS
