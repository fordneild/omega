package awsupboundio


// Source is the source of the web identity token.
type ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigSource string

const (
	// Secret.
	ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigSource_SECRET ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigSource = "SECRET"
	// Filesystem.
	ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigSource_FILESYSTEM ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigSource = "FILESYSTEM"
)

