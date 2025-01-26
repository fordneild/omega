package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

type ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount
type jsiiProxy_ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount_FromNumber(value *float64) ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount_FromString(value *string) ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

