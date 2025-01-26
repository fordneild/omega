package argoprojio


type ApplicationSetSpecGeneratorsList struct {
	Elements *[]interface{} `field:"optional" json:"elements" yaml:"elements"`
	ElementsYaml *string `field:"optional" json:"elementsYaml" yaml:"elementsYaml"`
	Template *ApplicationSetSpecGeneratorsListTemplate `field:"optional" json:"template" yaml:"template"`
}

