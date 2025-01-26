package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

type ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount
type jsiiProxy_ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount_FromNumber(value *float64) ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount_FromString(value *string) ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

