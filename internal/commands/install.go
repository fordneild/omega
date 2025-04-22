package commands

import (
	"github.com/fordneild/omega/internal/globals"
	"github.com/fordneild/omega/internal/shell"
)

type InstallCmd struct {
}

func (cmd *InstallCmd) Run(globals *globals.Globals) error {
	_, _, err := shell.ExecPrint("go install ~/sources/omega/cmd/omega")
	return err
}
