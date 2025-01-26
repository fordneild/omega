package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

type ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount
type jsiiProxy_ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount_FromNumber(value *float64) ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount_FromString(value *string) ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSourceKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

