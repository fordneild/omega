package resources

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/fordneild/omega/internal/globals"
	"github.com/fordneild/omega/internal/project"
	"github.com/fordneild/omega/internal/shell"
)

type SyncCmd struct {
	ProjectsIds *[]string `name:"project-ids" help:"Projects to sync"`
}

func SyncProject(globals *globals.Globals, project project.Project) error {
	flags := []string{}
	if globals.Debug {
		flags = append(flags, "-d")
	}
	omegaSyncCommand := fmt.Sprintf("omega synth %s %s", project.GetId(), strings.Join(flags, " "))
	log.Infof("Syncing project %s", project.GetId())
	cmd := fmt.Sprintf("cdk8s synth --output=\"%s\" --app=\"%s\"", project.GetPath(), omegaSyncCommand)
	log.Debugf("Running command: %s", cmd)
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
		err := SyncProject(globals, project)
		if err != nil {
			return err
		}
	}
	return nil
}
