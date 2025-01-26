package crdk8samazonawscom

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ENIConfig is the Schema for the eniconfigs API.
type EniConfigProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ENIConfigSpec defines the desired state of ENIConfig.
	Spec *EniConfigSpec `field:"optional" json:"spec" yaml:"spec"`
}

