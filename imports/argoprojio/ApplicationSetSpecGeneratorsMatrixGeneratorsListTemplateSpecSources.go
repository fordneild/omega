package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSources struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	Directory *ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSourcesDirectory `field:"optional" json:"directory" yaml:"directory"`
	Helm *ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSourcesHelm `field:"optional" json:"helm" yaml:"helm"`
	Kustomize *ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSourcesKustomize `field:"optional" json:"kustomize" yaml:"kustomize"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Plugin *ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSourcesPlugin `field:"optional" json:"plugin" yaml:"plugin"`
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	TargetRevision *string `field:"optional" json:"targetRevision" yaml:"targetRevision"`
}

