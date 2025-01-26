package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

type ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount
type jsiiProxy_ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount_FromNumber(value *float64) ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount_FromString(value *string) ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

