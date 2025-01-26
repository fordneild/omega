package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsClustersTemplateSpec struct {
	Destination *ApplicationSetSpecGeneratorsMergeGeneratorsClustersTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	IgnoreDifferences *[]*ApplicationSetSpecGeneratorsMergeGeneratorsClustersTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecGeneratorsMergeGeneratorsClustersTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	Source *ApplicationSetSpecGeneratorsMergeGeneratorsClustersTemplateSpecSource `field:"optional" json:"source" yaml:"source"`
	Sources *[]*ApplicationSetSpecGeneratorsMergeGeneratorsClustersTemplateSpecSources `field:"optional" json:"sources" yaml:"sources"`
	SyncPolicy *ApplicationSetSpecGeneratorsMergeGeneratorsClustersTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

