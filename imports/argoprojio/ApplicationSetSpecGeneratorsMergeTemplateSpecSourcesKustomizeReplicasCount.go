package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

type ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount
type jsiiProxy_ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount_FromNumber(value *float64) ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount_FromString(value *string) ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMergeTemplateSpecSourcesKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

