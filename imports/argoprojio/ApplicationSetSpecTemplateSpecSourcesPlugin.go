package argoprojio


type ApplicationSetSpecTemplateSpecSourcesPlugin struct {
	Env *[]*ApplicationSetSpecTemplateSpecSourcesPluginEnv `field:"optional" json:"env" yaml:"env"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Parameters *[]*ApplicationSetSpecTemplateSpecSourcesPluginParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

