package resources

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/fordneild/omega/internal/project"
)

func NewProjectRoot(project project.Project) cdk8s.App {
	app := cdk8s.NewApp(nil)
	chart := cdk8s.NewChart(app, jsii.String(fmt.Sprintf("%s-project", project.GetId())), nil)
	NewProjectArgocdResources(chart, project.GetId(), ProjectProps{
		ProjectName:    project.GetName(),
		RepoUrl:        project.GetRepo(),
		Path:           project.GetPath(),
		TargetRevision: project.GetTargetRevision(),
	})

	return app
}
