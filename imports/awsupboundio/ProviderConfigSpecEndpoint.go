package awsupboundio


// Endpoint is where you can override the default endpoint configuration of AWS calls made by the provider.
type ProviderConfigSpecEndpoint struct {
	// URL lets you configure the endpoint URL to be used in SDK calls.
	Url *ProviderConfigSpecEndpointUrl `field:"required" json:"url" yaml:"url"`
	// Specifies if the endpoint's hostname can be modified by the SDK's API client.
	//
	// If the hostname is mutable the SDK API clients may modify any part of
	// the hostname based on the requirements of the API, (e.g. adding, or
	// removing content in the hostname). Such as, Amazon S3 API client
	// prefixing "bucketname" to the hostname, or changing the
	// hostname service name component from "s3." to "s3-accesspoint.dualstack."
	// for the dualstack endpoint of an S3 Accesspoint resource.
	//
	//
	// Care should be taken when providing a custom endpoint for an API. If the
	// endpoint hostname is mutable, and the client cannot modify the endpoint
	// correctly, the operation call will most likely fail, or have undefined
	// behavior.
	//
	//
	// If hostname is immutable, the SDK API clients will not modify the
	// hostname of the URL. This may cause the API client not to function
	// correctly if the API requires the operation specific hostname values
	// to be used by the client.
	//
	//
	// This flag does not modify the API client's behavior if this endpoint
	// will be used instead of Endpoint Discovery, or if the endpoint will be
	// used to perform Endpoint Discovery. That behavior is configured via the
	// API Client's Options.
	// Note that this is effective only for resources that use AWS SDK v2.
	HostnameImmutable *bool `field:"optional" json:"hostnameImmutable" yaml:"hostnameImmutable"`
	// The AWS partition the endpoint belongs to.
	PartitionId *string `field:"optional" json:"partitionId" yaml:"partitionId"`
	// Specifies the list of services you want endpoint to be used for.
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
	// The signing method that should be used for signing the requests to the endpoint.
	SigningMethod *string `field:"optional" json:"signingMethod" yaml:"signingMethod"`
	// The service name that should be used for signing the requests to the endpoint.
	SigningName *string `field:"optional" json:"signingName" yaml:"signingName"`
	// The region that should be used for signing the request to the endpoint.
	//
	// For IAM, which doesn't have any region, us-east-1 is used to sign the
	// requests, which is the only signing region of IAM.
	SigningRegion *string `field:"optional" json:"signingRegion" yaml:"signingRegion"`
	// The source of the Endpoint.
	//
	// By default, this will be ServiceMetadata.
	// When providing a custom endpoint, you should set the source as Custom.
	// If source is not provided when providing a custom endpoint, the SDK may not
	// perform required host mutations correctly. Source should be used along with
	// HostnameImmutable property as per the usage requirement.
	// Note that this is effective only for resources that use AWS SDK v2.
	Source ProviderConfigSpecEndpointSource `field:"optional" json:"source" yaml:"source"`
}

