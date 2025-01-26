//go:build !no_runtime_type_checking

package argoprojio

import (
	"fmt"
)

func validateApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourcesKustomizeReplicasCount_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourcesKustomizeReplicasCount_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

