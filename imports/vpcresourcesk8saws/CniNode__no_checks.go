//go:build no_runtime_type_checking

package vpcresourcesk8saws

// Building without runtime type checking enabled, so all the below just return nil

func validateCniNode_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateCniNode_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCniNode_ManifestParameters(props *CniNodeProps) error {
	return nil
}

func validateCniNode_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewCniNodeParameters(scope constructs.Construct, id *string, props *CniNodeProps) error {
	return nil
}

