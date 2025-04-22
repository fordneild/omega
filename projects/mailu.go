package projects

import (
	"github.com/fordneild/omega/internal/project"
)

const MailuId = "mailu"

func mailu() project.Project {
	project := project.OmegaProject(MailuId, "mailu", nil)
	return project
}

var MailuProject = mailu()
