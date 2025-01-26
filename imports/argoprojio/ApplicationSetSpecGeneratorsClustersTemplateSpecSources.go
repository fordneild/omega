package argoprojio


type ApplicationSetSpecGeneratorsClustersTemplateSpecSources struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	Directory *ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesDirectory `field:"optional" json:"directory" yaml:"directory"`
	Helm *ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesHelm `field:"optional" json:"helm" yaml:"helm"`
	Kustomize *ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomize `field:"optional" json:"kustomize" yaml:"kustomize"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Plugin *ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesPlugin `field:"optional" json:"plugin" yaml:"plugin"`
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	TargetRevision *string `field:"optional" json:"targetRevision" yaml:"targetRevision"`
}

