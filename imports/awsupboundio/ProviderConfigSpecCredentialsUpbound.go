package awsupboundio


// Upbound defines the options for authenticating using Upbound as an identity provider.
type ProviderConfigSpecCredentialsUpbound struct {
	// WebIdentity defines the options for assuming an IAM role with a Web Identity.
	WebIdentity *ProviderConfigSpecCredentialsUpboundWebIdentity `field:"optional" json:"webIdentity" yaml:"webIdentity"`
}

