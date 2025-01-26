package argoprojio


// Plugin holds config management plugin specific options.
type ApplicationOperationSyncSourcesPlugin struct {
	// Env is a list of environment variable entries.
	Env *[]*ApplicationOperationSyncSourcesPluginEnv `field:"optional" json:"env" yaml:"env"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Parameters *[]*ApplicationOperationSyncSourcesPluginParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

