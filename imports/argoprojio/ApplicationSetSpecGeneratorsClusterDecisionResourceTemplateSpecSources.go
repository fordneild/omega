package argoprojio


type ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSources struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	Directory *ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourcesDirectory `field:"optional" json:"directory" yaml:"directory"`
	Helm *ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourcesHelm `field:"optional" json:"helm" yaml:"helm"`
	Kustomize *ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomize `field:"optional" json:"kustomize" yaml:"kustomize"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Plugin *ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourcesPlugin `field:"optional" json:"plugin" yaml:"plugin"`
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	TargetRevision *string `field:"optional" json:"targetRevision" yaml:"targetRevision"`
}

