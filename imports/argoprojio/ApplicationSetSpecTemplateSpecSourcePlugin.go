package argoprojio


type ApplicationSetSpecTemplateSpecSourcePlugin struct {
	Env *[]*ApplicationSetSpecTemplateSpecSourcePluginEnv `field:"optional" json:"env" yaml:"env"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Parameters *[]*ApplicationSetSpecTemplateSpecSourcePluginParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

