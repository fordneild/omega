package awsupboundio


// WebIdentity defines the options for assuming an IAM role with a Web Identity.
type ProviderConfigSpecCredentialsUpboundWebIdentity struct {
	// AssumeRoleARN to assume with provider credentials.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// RoleSessionName is the session name, if you wish to uniquely identify this session.
	RoleSessionName *string `field:"optional" json:"roleSessionName" yaml:"roleSessionName"`
	// TokenConfig is the Web Identity Token config to assume the role.
	TokenConfig *ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfig `field:"optional" json:"tokenConfig" yaml:"tokenConfig"`
}

