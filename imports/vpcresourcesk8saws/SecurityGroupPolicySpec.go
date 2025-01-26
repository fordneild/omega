package vpcresourcesk8saws


// SecurityGroupPolicySpec defines the desired state of SecurityGroupPolicy.
type SecurityGroupPolicySpec struct {
	// A label selector is a label query over a set of resources.
	//
	// The result of matchLabels and
	// matchExpressions are ANDed. An empty label selector matches all objects. A null
	// label selector matches no objects.
	PodSelector *SecurityGroupPolicySpecPodSelector `field:"optional" json:"podSelector" yaml:"podSelector"`
	// GroupIds contains the list of security groups that will be applied to the network interface of the pod matching the criteria.
	SecurityGroups *SecurityGroupPolicySpecSecurityGroups `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A label selector is a label query over a set of resources.
	//
	// The result of matchLabels and
	// matchExpressions are ANDed. An empty label selector matches all objects. A null
	// label selector matches no objects.
	ServiceAccountSelector *SecurityGroupPolicySpecServiceAccountSelector `field:"optional" json:"serviceAccountSelector" yaml:"serviceAccountSelector"`
}

