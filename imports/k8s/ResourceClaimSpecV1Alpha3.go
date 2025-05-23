package k8s


// ResourceClaimSpec defines what is being requested in a ResourceClaim and how to configure it.
type ResourceClaimSpecV1Alpha3 struct {
	// Controller is the name of the DRA driver that is meant to handle allocation of this claim.
	//
	// If empty, allocation is handled by the scheduler while scheduling a pod.
	//
	// Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver.
	//
	// This is an alpha field and requires enabling the DRAControlPlaneController feature gate.
	Controller *string `field:"optional" json:"controller" yaml:"controller"`
	// Devices defines how to request devices.
	Devices *DeviceClaimV1Alpha3 `field:"optional" json:"devices" yaml:"devices"`
}

