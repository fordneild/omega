package argoprojio


type ApplicationSetSpecGeneratorsPluginTemplate struct {
	Metadata *ApplicationSetSpecGeneratorsPluginTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	Spec *ApplicationSetSpecGeneratorsPluginTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

