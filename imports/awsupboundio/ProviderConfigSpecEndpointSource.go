package awsupboundio


// The source of the Endpoint.
//
// By default, this will be ServiceMetadata.
// When providing a custom endpoint, you should set the source as Custom.
// If source is not provided when providing a custom endpoint, the SDK may not
// perform required host mutations correctly. Source should be used along with
// HostnameImmutable property as per the usage requirement.
// Note that this is effective only for resources that use AWS SDK v2.
type ProviderConfigSpecEndpointSource string

const (
	// ServiceMetadata.
	ProviderConfigSpecEndpointSource_SERVICE_METADATA ProviderConfigSpecEndpointSource = "SERVICE_METADATA"
	// Custom.
	ProviderConfigSpecEndpointSource_CUSTOM ProviderConfigSpecEndpointSource = "CUSTOM"
)

