package argoprojio


type ApplicationSetSpecPreservedFields struct {
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

