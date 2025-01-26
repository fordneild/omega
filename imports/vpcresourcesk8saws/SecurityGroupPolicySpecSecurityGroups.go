package vpcresourcesk8saws


// GroupIds contains the list of security groups that will be applied to the network interface of the pod matching the criteria.
type SecurityGroupPolicySpecSecurityGroups struct {
	// Groups is the list of EC2 Security Groups Ids that need to be applied to the ENI of a Pod.
	GroupIds *[]*string `field:"optional" json:"groupIds" yaml:"groupIds"`
}

