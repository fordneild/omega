package vpcresourcesk8saws


// Important: Run "make" to regenerate code after modifying this file CNINodeSpec defines the desired state of CNINode.
type CniNodeSpec struct {
	Features *[]*CniNodeSpecFeatures `field:"optional" json:"features" yaml:"features"`
}

