package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

type ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount
type jsiiProxy_ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount_FromNumber(value *float64) ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount_FromString(value *string) ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSourceKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

