package argoprojio


type ApplicationSetSpecGeneratorsScmProviderTemplateSpec struct {
	Destination *ApplicationSetSpecGeneratorsScmProviderTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	IgnoreDifferences *[]*ApplicationSetSpecGeneratorsScmProviderTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecGeneratorsScmProviderTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	Source *ApplicationSetSpecGeneratorsScmProviderTemplateSpecSource `field:"optional" json:"source" yaml:"source"`
	Sources *[]*ApplicationSetSpecGeneratorsScmProviderTemplateSpecSources `field:"optional" json:"sources" yaml:"sources"`
	SyncPolicy *ApplicationSetSpecGeneratorsScmProviderTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

