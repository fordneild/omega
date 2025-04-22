package awsupboundio


// You can provide a static URL that will be used regardless of the service and region by choosing Static type.
//
// Alternatively, you can provide
// configuration for dynamically resolving the URL with the config you provide
// once you set the type as Dynamic.
type ProviderConfigSpecEndpointUrlType string

const (
	// Static.
	ProviderConfigSpecEndpointUrlType_STATIC ProviderConfigSpecEndpointUrlType = "STATIC"
	// Dynamic.
	ProviderConfigSpecEndpointUrlType_DYNAMIC ProviderConfigSpecEndpointUrlType = "DYNAMIC"
	// Auto.
	ProviderConfigSpecEndpointUrlType_AUTO ProviderConfigSpecEndpointUrlType = "AUTO"
)

