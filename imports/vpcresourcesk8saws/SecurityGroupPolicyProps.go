package vpcresourcesk8saws

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Custom Resource Definition for applying security groups to pods.
type SecurityGroupPolicyProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// SecurityGroupPolicySpec defines the desired state of SecurityGroupPolicy.
	Spec *SecurityGroupPolicySpec `field:"optional" json:"spec" yaml:"spec"`
}

