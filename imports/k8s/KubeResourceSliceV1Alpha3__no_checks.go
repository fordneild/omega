//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeResourceSliceV1Alpha3_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeResourceSliceV1Alpha3_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeResourceSliceV1Alpha3_ManifestParameters(props *KubeResourceSliceV1Alpha3Props) error {
	return nil
}

func validateKubeResourceSliceV1Alpha3_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeResourceSliceV1Alpha3Parameters(scope constructs.Construct, id *string, props *KubeResourceSliceV1Alpha3Props) error {
	return nil
}

