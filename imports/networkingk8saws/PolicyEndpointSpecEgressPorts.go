package networkingk8saws


// Port contains information about the transport port/protocol.
type PolicyEndpointSpecEgressPorts struct {
	// Endport specifies the port range port to endPort port must be defined and an integer, endPort > port.
	EndPort *float64 `field:"optional" json:"endPort" yaml:"endPort"`
	// Port specifies the numerical port for the protocol.
	//
	// If empty applies to all ports.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Protocol specifies the transport protocol, default TCP.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

