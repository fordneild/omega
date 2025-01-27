package resources

import (
	"fmt"
	"slices"

	"github.com/charmbracelet/log"
	"github.com/fordneild/omega/internal/globals"
	"github.com/fordneild/omega/internal/project"
	"github.com/fordneild/omega/internal/shell"
)

type SyncCmd struct {
	ProjectsIds *[]string `name:"project-ids" help:"Projects to sync"`
}

func SyncProject(project project.Project) error {
	omegaSyncCommand := fmt.Sprintf("omega synth %s", project.GetId())
	log.Infof("Syncing project %s", project.GetId())
	cmd := fmt.Sprintf("cdk8s synth --output=\"%s\" --app=\"%s\"", project.GetPath(), omegaSyncCommand)
	log.Infof("Running command: %s", cmd)
	stdout, stderr, err := shell.Exec(cmd)
	if stdout != "" {
		log.Print(stdout)
	}
	if err != nil {
		if stderr != "" {
			log.Errorf("\n%s", stderr)
		}
		return err
	}
	return err
}

func (cmd *SyncCmd) Run(globals *globals.Globals) error {
	projectIds := cmd.ProjectsIds
	projects := globals.ProjectService.GetAllProjects()
	for _, project := range projects {
		if projectIds != nil && !slices.Contains(*projectIds, project.GetId()) {
			continue
		}
		err := SyncProject(project)
		if err != nil {
			return err
		}
	}
	return nil
}
