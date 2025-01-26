package crdk8samazonawscom


// ENIConfigSpec defines the desired state of ENIConfig.
type EniConfigSpec struct {
	Subnet *string `field:"required" json:"subnet" yaml:"subnet"`
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

