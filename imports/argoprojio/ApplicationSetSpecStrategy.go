package argoprojio


type ApplicationSetSpecStrategy struct {
	RollingSync *ApplicationSetSpecStrategyRollingSync `field:"optional" json:"rollingSync" yaml:"rollingSync"`
	Type *string `field:"optional" json:"type" yaml:"type"`
}

