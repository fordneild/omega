package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsPlugin struct {
	ConfigMapRef *ApplicationSetSpecGeneratorsMatrixGeneratorsPluginConfigMapRef `field:"required" json:"configMapRef" yaml:"configMapRef"`
	Input *ApplicationSetSpecGeneratorsMatrixGeneratorsPluginInput `field:"optional" json:"input" yaml:"input"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplate `field:"optional" json:"template" yaml:"template"`
	Values *map[string]*string `field:"optional" json:"values" yaml:"values"`
}

