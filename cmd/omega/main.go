package main

import (
	"github.com/alecthomas/kong"
	"github.com/fordneild/omega/internal/globals"
	"github.com/fordneild/omega/internal/infra"
	"github.com/fordneild/omega/internal/install"
	"github.com/fordneild/omega/internal/resources"
)

var cli struct {
	resources.ResourcesCLI
	infra.InfraCLI
	install.InstallCommandLine
	Debug bool `help:"Enable debug mode."`
}

func main() {
	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&globals.Globals{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
