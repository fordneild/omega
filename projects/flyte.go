package projects

import (
	"github.com/fordneild/omega/internal/project"
	"github.com/fordneild/omega/internal/resources"
)

const FlyteId = "flyte"

func flyte() project.Project {
	project := project.OmegaProject(FlyteId, "flyte")
	project.WithApp(resources.NewFlyteHelmChart(project))
	return project
}

var FlyteProject = flyte()
