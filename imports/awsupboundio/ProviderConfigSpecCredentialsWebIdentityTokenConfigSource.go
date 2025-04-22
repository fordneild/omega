package awsupboundio


// Source is the source of the web identity token.
type ProviderConfigSpecCredentialsWebIdentityTokenConfigSource string

const (
	// Secret.
	ProviderConfigSpecCredentialsWebIdentityTokenConfigSource_SECRET ProviderConfigSpecCredentialsWebIdentityTokenConfigSource = "SECRET"
	// Filesystem.
	ProviderConfigSpecCredentialsWebIdentityTokenConfigSource_FILESYSTEM ProviderConfigSpecCredentialsWebIdentityTokenConfigSource = "FILESYSTEM"
)

