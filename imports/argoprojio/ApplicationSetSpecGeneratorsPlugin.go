package argoprojio


type ApplicationSetSpecGeneratorsPlugin struct {
	ConfigMapRef *ApplicationSetSpecGeneratorsPluginConfigMapRef `field:"required" json:"configMapRef" yaml:"configMapRef"`
	Input *ApplicationSetSpecGeneratorsPluginInput `field:"optional" json:"input" yaml:"input"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsPluginTemplate `field:"optional" json:"template" yaml:"template"`
	Values *map[string]*string `field:"optional" json:"values" yaml:"values"`
}

