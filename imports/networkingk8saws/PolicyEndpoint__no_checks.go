//go:build no_runtime_type_checking

package networkingk8saws

// Building without runtime type checking enabled, so all the below just return nil

func validatePolicyEndpoint_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validatePolicyEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePolicyEndpoint_ManifestParameters(props *PolicyEndpointProps) error {
	return nil
}

func validatePolicyEndpoint_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewPolicyEndpointParameters(scope constructs.Construct, id *string, props *PolicyEndpointProps) error {
	return nil
}

