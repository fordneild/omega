package networkingk8saws


// PolicyEndpointSpec defines the desired state of PolicyEndpoint.
type PolicyEndpointSpec struct {
	// PolicyRef is a reference to the Kubernetes NetworkPolicy resource.
	PolicyRef *PolicyEndpointSpecPolicyRef `field:"required" json:"policyRef" yaml:"policyRef"`
	// Egress is the list of egress rules containing resolved network addresses.
	Egress *[]*PolicyEndpointSpecEgress `field:"optional" json:"egress" yaml:"egress"`
	// Ingress is the list of ingress rules containing resolved network addresses.
	Ingress *[]*PolicyEndpointSpecIngress `field:"optional" json:"ingress" yaml:"ingress"`
	// PodIsolation specifies whether the pod needs to be isolated for a particular traffic direction Ingress or Egress, or both.
	//
	// If default isolation is not specified, and there are no ingress/egress rules, then the pod is not isolated from the point of view of this policy. This follows the NetworkPolicy spec.PolicyTypes.
	PodIsolation *[]*string `field:"optional" json:"podIsolation" yaml:"podIsolation"`
	// PodSelector is the podSelector from the policy resource.
	PodSelector *PolicyEndpointSpecPodSelector `field:"optional" json:"podSelector" yaml:"podSelector"`
	// PodSelectorEndpoints contains information about the pods matching the podSelector.
	PodSelectorEndpoints *[]*PolicyEndpointSpecPodSelectorEndpoints `field:"optional" json:"podSelectorEndpoints" yaml:"podSelectorEndpoints"`
}

