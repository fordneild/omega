//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeStorageVersionMigrationV1Alpha1_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeStorageVersionMigrationV1Alpha1_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeStorageVersionMigrationV1Alpha1_ManifestParameters(props *KubeStorageVersionMigrationV1Alpha1Props) error {
	return nil
}

func validateKubeStorageVersionMigrationV1Alpha1_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeStorageVersionMigrationV1Alpha1Parameters(scope constructs.Construct, id *string, props *KubeStorageVersionMigrationV1Alpha1Props) error {
	return nil
}

