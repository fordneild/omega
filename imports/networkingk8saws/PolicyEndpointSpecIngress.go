package networkingk8saws


// EndpointInfo defines the network endpoint information for the policy ingress/egress.
type PolicyEndpointSpecIngress struct {
	// CIDR is the network address(s) of the endpoint.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Except is the exceptions to the CIDR ranges mentioned above.
	Except *[]*string `field:"optional" json:"except" yaml:"except"`
	// Ports is the list of ports.
	Ports *[]*PolicyEndpointSpecIngressPorts `field:"optional" json:"ports" yaml:"ports"`
}

