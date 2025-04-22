package awsupboundio


// URL lets you configure the endpoint URL to be used in SDK calls.
type ProviderConfigSpecEndpointUrl struct {
	// You can provide a static URL that will be used regardless of the service and region by choosing Static type.
	//
	// Alternatively, you can provide
	// configuration for dynamically resolving the URL with the config you provide
	// once you set the type as Dynamic.
	Type ProviderConfigSpecEndpointUrlType `field:"required" json:"type" yaml:"type"`
	// Dynamic lets you configure the behavior of endpoint URL resolver.
	Dynamic *ProviderConfigSpecEndpointUrlDynamic `field:"optional" json:"dynamic" yaml:"dynamic"`
	// Static is the full URL you'd like the AWS SDK to use.
	//
	// Recommended for using tools like localstack where a single host is exposed
	// for all services and regions.
	Static *string `field:"optional" json:"static" yaml:"static"`
}

