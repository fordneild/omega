package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

type ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount
type jsiiProxy_ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount_FromNumber(value *float64) ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount_FromString(value *string) ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSourcesKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

