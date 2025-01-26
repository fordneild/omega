package argoprojio


type ApplicationSpecSourcesKustomizePatches struct {
	Options *map[string]*bool `field:"optional" json:"options" yaml:"options"`
	Patch *string `field:"optional" json:"patch" yaml:"patch"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Target *ApplicationSpecSourcesKustomizePatchesTarget `field:"optional" json:"target" yaml:"target"`
}

