package argoprojio


type ApplicationSetSpecTemplateSpecSources struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	Directory *ApplicationSetSpecTemplateSpecSourcesDirectory `field:"optional" json:"directory" yaml:"directory"`
	Helm *ApplicationSetSpecTemplateSpecSourcesHelm `field:"optional" json:"helm" yaml:"helm"`
	Kustomize *ApplicationSetSpecTemplateSpecSourcesKustomize `field:"optional" json:"kustomize" yaml:"kustomize"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Plugin *ApplicationSetSpecTemplateSpecSourcesPlugin `field:"optional" json:"plugin" yaml:"plugin"`
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	TargetRevision *string `field:"optional" json:"targetRevision" yaml:"targetRevision"`
}

