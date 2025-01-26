package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

type ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount
type jsiiProxy_ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount_FromNumber(value *float64) ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount_FromString(value *string) ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSourceKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

