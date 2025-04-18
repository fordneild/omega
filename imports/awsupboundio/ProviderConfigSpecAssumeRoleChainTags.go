package awsupboundio


// Tag is session tag that can be used to assume an IAM Role.
type ProviderConfigSpecAssumeRoleChainTags struct {
	// Name of the tag.
	//
	// Key is a required field.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value of the tag.
	//
	// Value is a required field.
	Value *string `field:"required" json:"value" yaml:"value"`
}

