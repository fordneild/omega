package argoprojio


type ApplicationSetSpecTemplateSpec struct {
	Destination *ApplicationSetSpecTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	IgnoreDifferences *[]*ApplicationSetSpecTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	Source *ApplicationSetSpecTemplateSpecSource `field:"optional" json:"source" yaml:"source"`
	Sources *[]*ApplicationSetSpecTemplateSpecSources `field:"optional" json:"sources" yaml:"sources"`
	SyncPolicy *ApplicationSetSpecTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

