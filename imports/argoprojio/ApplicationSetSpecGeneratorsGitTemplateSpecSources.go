package argoprojio


type ApplicationSetSpecGeneratorsGitTemplateSpecSources struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	Directory *ApplicationSetSpecGeneratorsGitTemplateSpecSourcesDirectory `field:"optional" json:"directory" yaml:"directory"`
	Helm *ApplicationSetSpecGeneratorsGitTemplateSpecSourcesHelm `field:"optional" json:"helm" yaml:"helm"`
	Kustomize *ApplicationSetSpecGeneratorsGitTemplateSpecSourcesKustomize `field:"optional" json:"kustomize" yaml:"kustomize"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Plugin *ApplicationSetSpecGeneratorsGitTemplateSpecSourcesPlugin `field:"optional" json:"plugin" yaml:"plugin"`
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	TargetRevision *string `field:"optional" json:"targetRevision" yaml:"targetRevision"`
}

