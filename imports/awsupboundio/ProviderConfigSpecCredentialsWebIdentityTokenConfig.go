package awsupboundio


// TokenConfig is the Web Identity Token config to assume the role.
type ProviderConfigSpecCredentialsWebIdentityTokenConfig struct {
	// Source is the source of the web identity token.
	Source ProviderConfigSpecCredentialsWebIdentityTokenConfigSource `field:"required" json:"source" yaml:"source"`
	// Fs is a reference to a filesystem location that contains credentials that must be used to obtain the web identity token.
	Fs *ProviderConfigSpecCredentialsWebIdentityTokenConfigFs `field:"optional" json:"fs" yaml:"fs"`
	// A SecretRef is a reference to a secret key that contains the credentials that must be used to obtain the web identity token.
	SecretRef *ProviderConfigSpecCredentialsWebIdentityTokenConfigSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

