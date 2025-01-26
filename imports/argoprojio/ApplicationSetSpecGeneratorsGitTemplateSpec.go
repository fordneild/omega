package argoprojio


type ApplicationSetSpecGeneratorsGitTemplateSpec struct {
	Destination *ApplicationSetSpecGeneratorsGitTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	IgnoreDifferences *[]*ApplicationSetSpecGeneratorsGitTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecGeneratorsGitTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	Source *ApplicationSetSpecGeneratorsGitTemplateSpecSource `field:"optional" json:"source" yaml:"source"`
	Sources *[]*ApplicationSetSpecGeneratorsGitTemplateSpecSources `field:"optional" json:"sources" yaml:"sources"`
	SyncPolicy *ApplicationSetSpecGeneratorsGitTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

