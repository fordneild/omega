package awsupboundio


// Source of the provider credentials.
type ProviderConfigSpecCredentialsSource string

const (
	// None.
	ProviderConfigSpecCredentialsSource_NONE ProviderConfigSpecCredentialsSource = "NONE"
	// Secret.
	ProviderConfigSpecCredentialsSource_SECRET ProviderConfigSpecCredentialsSource = "SECRET"
	// IRSA.
	ProviderConfigSpecCredentialsSource_IRSA ProviderConfigSpecCredentialsSource = "IRSA"
	// WebIdentity.
	ProviderConfigSpecCredentialsSource_WEB_IDENTITY ProviderConfigSpecCredentialsSource = "WEB_IDENTITY"
	// PodIdentity.
	ProviderConfigSpecCredentialsSource_POD_IDENTITY ProviderConfigSpecCredentialsSource = "POD_IDENTITY"
	// Upbound.
	ProviderConfigSpecCredentialsSource_UPBOUND ProviderConfigSpecCredentialsSource = "UPBOUND"
)

