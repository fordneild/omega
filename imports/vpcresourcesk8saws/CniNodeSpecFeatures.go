package vpcresourcesk8saws


// Feature is a type of feature being supported by VPC resource controller and other AWS Services.
type CniNodeSpecFeatures struct {
	// FeatureName is a type of feature name supported by AWS VPC CNI.
	//
	// It can be Security Group for Pods, custom networking, or others.
	Name *string `field:"optional" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
}

