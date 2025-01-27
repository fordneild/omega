package projects

import "github.com/fordneild/omega/internal/project"

var AllProjects = []project.Project{
	CrossplaneProject,
}

// singleton injected as global into CLI
var ProjectService = project.NewProjectService(AllProjects, RootProject)
