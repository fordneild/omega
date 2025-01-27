package project

import "fmt"

func NewProjectService(projects []Project, rootProject Project) *ProjectService {
	return &ProjectService{RootProject: rootProject, Projects: projects}
}

type ProjectService struct {
	RootProject Project
	Projects    []Project
}

func (s *ProjectService) GetProjectById(id string) (*Project, error) {
	for _, project := range s.GetAllProjects() {
		if project.GetId() == id {
			return &project, nil
		}
	}
	return nil, fmt.Errorf("project not found")
}

func (s *ProjectService) GetAllProjects() []Project {
	return append([]Project{s.RootProject}, s.Projects...)
}
