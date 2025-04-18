package commands

import (
	"github.com/charmbracelet/log"
	"github.com/fordneild/omega/internal/globals"
)

type SynthCmd struct {
	ProjectId string `arg:""`
}

func (cmd *SynthCmd) Run(globals *globals.Globals) error {
	project, err := globals.ProjectService.GetProjectById(cmd.ProjectId)
	if err != nil {
		return err
	}
	proj := *project
	log.Infof("Synthing project %s", proj.GetId())
	log.Infof("Project %s has %d apps", proj.GetId(), len(proj.GetApps()))
	proj.Synth()
	return nil
}
