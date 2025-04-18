package awsupboundio


// Protocol is the HTTP protocol that will be used in the URL.
//
// Currently,
// only http and https are supported.
type ProviderConfigSpecEndpointUrlDynamicProtocol string

const (
	// http.
	ProviderConfigSpecEndpointUrlDynamicProtocol_HTTP ProviderConfigSpecEndpointUrlDynamicProtocol = "HTTP"
	// https.
	ProviderConfigSpecEndpointUrlDynamicProtocol_HTTPS ProviderConfigSpecEndpointUrlDynamicProtocol = "HTTPS"
)

