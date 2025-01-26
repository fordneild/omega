package infra

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/pulumi/pulumi/pkg/v3/cmd/pulumi/operations"
	"github.com/pulumi/pulumi/sdk/v3/go/common/version"

	"github.com/fordneild/omega/internal/globals"
)

type UpCmd struct {
}

func panicHandler(finished *bool) {
	if panicPayload := recover(); !*finished {
		stack := string(debug.Stack())
		fmt.Fprintln(os.Stderr, "================================================================================")
		fmt.Fprintln(os.Stderr, "The Pulumi CLI encountered a fatal error. This is a bug!")
		fmt.Fprintln(os.Stderr, "We would appreciate a report: https://github.com/pulumi/pulumi/issues/")
		fmt.Fprintln(os.Stderr, "Please provide all of the text below in your report.")
		fmt.Fprintln(os.Stderr, "================================================================================")
		fmt.Fprintf(os.Stderr, "Pulumi Version:   %s\n", version.Version)
		fmt.Fprintf(os.Stderr, "Go Version:       %s\n", runtime.Version())
		fmt.Fprintf(os.Stderr, "Go Compiler:      %s\n", runtime.Compiler)
		fmt.Fprintf(os.Stderr, "Architecture:     %s\n", runtime.GOARCH)
		fmt.Fprintf(os.Stderr, "Operating System: %s\n", runtime.GOOS)
		fmt.Fprintf(os.Stderr, "Panic:            %s\n\n", panicPayload)
		fmt.Fprintln(os.Stderr, stack)
		os.Exit(1)
	}
}

func (cmd *UpCmd) Run(globals *globals.Globals) error {
	finished := new(bool)
	defer panicHandler(finished)
	// up works magically
	// Create a new command
	pulumiUpCmd := operations.NewUpCmd()
	args := []string{"--config-file"}
	pulumiUpCmd.SetArgs(args)
	err := pulumiUpCmd.Execute()
	*finished = true
	return err
}
