package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/fordneild/omega/internal/commands"
	"github.com/fordneild/omega/internal/globals"
	"github.com/fordneild/omega/projects"
)

var cli struct {
	commands.AllCommands
	Debug bool `short:"d" help:"Enable debug mode."`
}

func main() {
	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	logLevel := log.InfoLevel
	if cli.Debug {
		logLevel = log.DebugLevel
	}
	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: false,
	})
	logger.SetLevel(logLevel)
	log.SetDefault(logger)
	err := ctx.Run(&globals.Globals{Debug: cli.Debug, ProjectService: projects.ProjectService})
	ctx.FatalIfErrorf(err)
}
