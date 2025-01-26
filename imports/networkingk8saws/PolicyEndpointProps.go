package networkingk8saws

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// PolicyEndpoint is the Schema for the policyendpoints API.
type PolicyEndpointProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// PolicyEndpointSpec defines the desired state of PolicyEndpoint.
	Spec *PolicyEndpointSpec `field:"optional" json:"spec" yaml:"spec"`
}

