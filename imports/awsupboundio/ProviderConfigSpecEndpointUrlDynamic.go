package awsupboundio


// Dynamic lets you configure the behavior of endpoint URL resolver.
type ProviderConfigSpecEndpointUrlDynamic struct {
	// Host is the address of the main host that the resolver will use to prepend protocol, service and region configurations.
	//
	// For example, the final URL for EC2 in us-east-1 looks like https://ec2.us-east-1.amazonaws.com
	// You would need to use "amazonaws.com" as Host and "https" as protocol
	// to have the resolver construct it.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Protocol is the HTTP protocol that will be used in the URL.
	//
	// Currently,
	// only http and https are supported.
	Protocol ProviderConfigSpecEndpointUrlDynamicProtocol `field:"required" json:"protocol" yaml:"protocol"`
}

