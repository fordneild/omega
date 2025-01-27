package project

import (
	"path/filepath"
)

const OmegaProjectBasePath = "rendered"

func OmegaProject(id string, name string) Project {
	project := NewProject(id, name, filepath.Join(OmegaProjectBasePath, name), "https://github.com/fordneild/omega.git")
	return project
}
