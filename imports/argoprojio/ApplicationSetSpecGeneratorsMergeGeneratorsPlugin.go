package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsPlugin struct {
	ConfigMapRef *ApplicationSetSpecGeneratorsMergeGeneratorsPluginConfigMapRef `field:"required" json:"configMapRef" yaml:"configMapRef"`
	Input *ApplicationSetSpecGeneratorsMergeGeneratorsPluginInput `field:"optional" json:"input" yaml:"input"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsMergeGeneratorsPluginTemplate `field:"optional" json:"template" yaml:"template"`
	Values *map[string]*string `field:"optional" json:"values" yaml:"values"`
}

