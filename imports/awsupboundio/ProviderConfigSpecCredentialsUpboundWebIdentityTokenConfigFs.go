package awsupboundio


// Fs is a reference to a filesystem location that contains credentials that must be used to obtain the web identity token.
type ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigFs struct {
	// Path is a filesystem path.
	Path *string `field:"required" json:"path" yaml:"path"`
}

