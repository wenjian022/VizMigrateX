package main

import (
	"embed"

	"VizMigrateX/cmd"
)

var EmbedFs embed.FS

func main() {
	cmd.Execute(EmbedFs)
}
