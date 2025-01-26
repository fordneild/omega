package networkingk8saws


// PolicyRef is a reference to the Kubernetes NetworkPolicy resource.
type PolicyEndpointSpecPolicyRef struct {
	// Name is the name of the Policy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace is the namespace of the Policy.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

