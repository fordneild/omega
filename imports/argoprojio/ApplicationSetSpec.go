package argoprojio


type ApplicationSetSpec struct {
	Generators *[]*ApplicationSetSpecGenerators `field:"required" json:"generators" yaml:"generators"`
	Template *ApplicationSetSpecTemplate `field:"required" json:"template" yaml:"template"`
	ApplyNestedSelectors *bool `field:"optional" json:"applyNestedSelectors" yaml:"applyNestedSelectors"`
	GoTemplate *bool `field:"optional" json:"goTemplate" yaml:"goTemplate"`
	GoTemplateOptions *[]*string `field:"optional" json:"goTemplateOptions" yaml:"goTemplateOptions"`
	IgnoreApplicationDifferences *[]*ApplicationSetSpecIgnoreApplicationDifferences `field:"optional" json:"ignoreApplicationDifferences" yaml:"ignoreApplicationDifferences"`
	PreservedFields *ApplicationSetSpecPreservedFields `field:"optional" json:"preservedFields" yaml:"preservedFields"`
	Strategy *ApplicationSetSpecStrategy `field:"optional" json:"strategy" yaml:"strategy"`
	SyncPolicy *ApplicationSetSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
	TemplatePatch *string `field:"optional" json:"templatePatch" yaml:"templatePatch"`
}

