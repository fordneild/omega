package awsupboundio


// Credentials required to authenticate to this provider.
type ProviderConfigSpecCredentials struct {
	// Source of the provider credentials.
	Source ProviderConfigSpecCredentialsSource `field:"required" json:"source" yaml:"source"`
	// Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.
	Env *ProviderConfigSpecCredentialsEnv `field:"optional" json:"env" yaml:"env"`
	// Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.
	Fs *ProviderConfigSpecCredentialsFs `field:"optional" json:"fs" yaml:"fs"`
	// A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.
	SecretRef *ProviderConfigSpecCredentialsSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	// Upbound defines the options for authenticating using Upbound as an identity provider.
	Upbound *ProviderConfigSpecCredentialsUpbound `field:"optional" json:"upbound" yaml:"upbound"`
	// WebIdentity defines the options for assuming an IAM role with a Web Identity.
	WebIdentity *ProviderConfigSpecCredentialsWebIdentity `field:"optional" json:"webIdentity" yaml:"webIdentity"`
}

