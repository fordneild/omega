package globals

import "github.com/fordneild/omega/internal/project"

type Globals struct {
	Debug          bool
	ProjectService *project.ProjectService
}
