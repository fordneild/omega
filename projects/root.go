package projects

import (
	"github.com/fordneild/omega/internal/project"
	"github.com/fordneild/omega/internal/resources"
)

func root() project.Project {
	rootProj := project.OmegaProject("root", "root", nil)
	for _, project := range AllProjects {
		app := resources.NewProjectRoot(project)
		rootProj.WithApp(app)
	}
	return rootProj
}

var RootProject = root()
