package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/fordneild/omega/internal/globals"
	"github.com/fordneild/omega/internal/infra"
	"github.com/fordneild/omega/internal/install"
	"github.com/fordneild/omega/internal/resources"
	"github.com/fordneild/omega/projects"
)

var cli struct {
	resources.ResourcesCLI
	infra.InfraCLI
	install.InstallCommandLine
	Debug bool `short:"d" help:"Enable debug mode."`
}

func main() {
	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&globals.Globals{Debug: cli.Debug, ProjectService: projects.ProjectService})
	logLevel := log.InfoLevel
	if cli.Debug {
		logLevel = log.DebugLevel
	}
	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: false,
	})
	logger.SetLevel(logLevel)
	log.SetDefault(logger)
	ctx.FatalIfErrorf(err)
}
