package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

type ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount
type jsiiProxy_ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount_FromNumber(value *float64) ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount_FromString(value *string) ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSourceKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

