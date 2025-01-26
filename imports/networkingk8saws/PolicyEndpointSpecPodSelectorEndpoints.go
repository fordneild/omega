package networkingk8saws


// PodEndpoint defines the summary information for the pods.
type PolicyEndpointSpecPodSelectorEndpoints struct {
	// HostIP is the IP address of the host the pod is currently running on.
	HostIp *string `field:"required" json:"hostIp" yaml:"hostIp"`
	// Name is the pod name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace is the pod namespace.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// PodIP is the IP address of the pod.
	PodIp *string `field:"required" json:"podIp" yaml:"podIp"`
}

