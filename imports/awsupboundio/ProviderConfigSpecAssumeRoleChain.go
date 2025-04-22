package awsupboundio


// AssumeRoleOptions define the options for assuming an IAM Role Fields are similar to the STS AssumeRoleOptions in the AWS SDK.
type ProviderConfigSpecAssumeRoleChain struct {
	// ExternalID is the external ID used when assuming role.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// AssumeRoleARN to assume with provider credentials.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags is list of session tags that you want to pass.
	//
	// Each session tag consists of a key
	// name and an associated value. For more information about session tags, see
	// Tagging STS Sessions
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html).
	Tags *[]*ProviderConfigSpecAssumeRoleChainTags `field:"optional" json:"tags" yaml:"tags"`
	// TransitiveTagKeys is a list of keys for session tags that you want to set as transitive.
	//
	// If you set a
	// tag key as transitive, the corresponding key and value passes to subsequent
	// sessions in a role chain. For more information, see Chaining Roles with Session Tags
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html#id_session-tags_role-chaining).
	TransitiveTagKeys *[]*string `field:"optional" json:"transitiveTagKeys" yaml:"transitiveTagKeys"`
}

