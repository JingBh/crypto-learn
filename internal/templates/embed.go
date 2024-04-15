package templates

import "embed"

//go:embed *.html
var EmbedFS embed.FS
