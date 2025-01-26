package argoprojio


type ApplicationSetSpecStrategyRollingSyncSteps struct {
	MatchExpressions *[]*ApplicationSetSpecStrategyRollingSyncStepsMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MaxUpdate ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate `field:"optional" json:"maxUpdate" yaml:"maxUpdate"`
}

