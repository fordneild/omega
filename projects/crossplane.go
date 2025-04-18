package projects

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/fordneild/omega/internal/project"
	"github.com/fordneild/omega/internal/resources"
)

const CrossplaneId = "crossplane"

func crossplane() project.Project {
	project := project.OmegaProject(CrossplaneId, "crossplane", jsii.String("fn-turks-vacation-dev"))
	project.WithApp(resources.NewAWSCrossplane(project))
	return project
}

var CrossplaneProject = crossplane()
