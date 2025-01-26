//go:build no_runtime_type_checking

package vpcresourcesk8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateSecurityGroupPolicy_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateSecurityGroupPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecurityGroupPolicy_ManifestParameters(props *SecurityGroupPolicyProps) error {
	return nil
}

func validateSecurityGroupPolicy_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewSecurityGroupPolicyParameters(scope constructs.Construct, id *string, props *SecurityGroupPolicyProps) error {
	return nil
}

