package vpcresourcesk8saws

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type CniNodeProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Important: Run "make" to regenerate code after modifying this file CNINodeSpec defines the desired state of CNINode.
	Spec *CniNodeSpec `field:"optional" json:"spec" yaml:"spec"`
}

