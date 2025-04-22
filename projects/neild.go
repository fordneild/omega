package projects

import (
	"github.com/fordneild/omega/internal/project"
)

const NeildId = "neild"

func neild() project.Project {
	project := project.OmegaProject(NeildId, "neild", nil)
	return project
}

var NeildProject = neild()
