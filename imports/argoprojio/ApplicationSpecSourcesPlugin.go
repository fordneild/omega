package argoprojio


// Plugin holds config management plugin specific options.
type ApplicationSpecSourcesPlugin struct {
	// Env is a list of environment variable entries.
	Env *[]*ApplicationSpecSourcesPluginEnv `field:"optional" json:"env" yaml:"env"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Parameters *[]*ApplicationSpecSourcesPluginParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

