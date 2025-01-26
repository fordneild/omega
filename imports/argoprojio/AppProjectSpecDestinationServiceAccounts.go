package argoprojio


// ApplicationDestinationServiceAccount holds information about the service account to be impersonated for the application sync operation.
type AppProjectSpecDestinationServiceAccounts struct {
	// DefaultServiceAccount to be used for impersonation during the sync operation.
	DefaultServiceAccount *string `field:"required" json:"defaultServiceAccount" yaml:"defaultServiceAccount"`
	// Server specifies the URL of the target cluster's Kubernetes control plane API.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Namespace specifies the target namespace for the application's resources.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

