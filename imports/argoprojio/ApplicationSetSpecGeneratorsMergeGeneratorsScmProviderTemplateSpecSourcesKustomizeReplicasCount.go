package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

type ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount
type jsiiProxy_ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount_FromNumber(value *float64) ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount_FromString(value *string) ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourcesKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

