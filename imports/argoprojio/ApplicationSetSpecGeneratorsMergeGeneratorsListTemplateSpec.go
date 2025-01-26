package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpec struct {
	Destination *ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	IgnoreDifferences *[]*ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	Source *ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSource `field:"optional" json:"source" yaml:"source"`
	Sources *[]*ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSources `field:"optional" json:"sources" yaml:"sources"`
	SyncPolicy *ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

