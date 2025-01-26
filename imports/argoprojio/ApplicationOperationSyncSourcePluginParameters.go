package argoprojio


type ApplicationOperationSyncSourcePluginParameters struct {
	// Array is the value of an array type parameter.
	Array *[]*string `field:"optional" json:"array" yaml:"array"`
	// Map is the value of a map type parameter.
	Map *map[string]*string `field:"optional" json:"map" yaml:"map"`
	// Name is the name identifying a parameter.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// String_ is the value of a string type parameter.
	String *string `field:"optional" json:"string" yaml:"string"`
}

