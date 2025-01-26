package argoprojio


type ApplicationSetSpecIgnoreApplicationDifferences struct {
	JqPathExpressions *[]*string `field:"optional" json:"jqPathExpressions" yaml:"jqPathExpressions"`
	JsonPointers *[]*string `field:"optional" json:"jsonPointers" yaml:"jsonPointers"`
	Name *string `field:"optional" json:"name" yaml:"name"`
}

