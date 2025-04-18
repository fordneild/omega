package awsupboundio


// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	Credentials *ProviderConfigSpecCredentials `field:"required" json:"credentials" yaml:"credentials"`
	// AssumeRoleChain defines the options for assuming an IAM role.
	AssumeRoleChain *[]*ProviderConfigSpecAssumeRoleChain `field:"optional" json:"assumeRoleChain" yaml:"assumeRoleChain"`
	// Endpoint is where you can override the default endpoint configuration of AWS calls made by the provider.
	Endpoint *ProviderConfigSpecEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Whether to enable the request to use path-style addressing, i.e., https://s3.amazonaws.com/BUCKET/KEY.
	S3UsePathStyle *bool `field:"optional" json:"s3UsePathStyle" yaml:"s3UsePathStyle"`
	// Whether to skip credentials validation via the STS API.
	//
	// This can be useful for testing and for AWS API implementations that do not have STS available.
	SkipCredentialsValidation *bool `field:"optional" json:"skipCredentialsValidation" yaml:"skipCredentialsValidation"`
	// Whether to skip the AWS Metadata API check Useful for AWS API implementations that do not have a metadata API endpoint.
	SkipMetadataApiCheck *bool `field:"optional" json:"skipMetadataApiCheck" yaml:"skipMetadataApiCheck"`
	// Whether to skip validation of provided region name.
	//
	// Useful for AWS-like implementations that use their own region names or to bypass the validation for
	// regions that aren't publicly available yet.
	SkipRegionValidation *bool `field:"optional" json:"skipRegionValidation" yaml:"skipRegionValidation"`
	// Whether to skip requesting the account ID.
	//
	// Useful for AWS API implementations that do not have the IAM, STS API, or metadata API.
	SkipRequestingAccountId *bool `field:"optional" json:"skipRequestingAccountId" yaml:"skipRequestingAccountId"`
}

