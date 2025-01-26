package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

// Number of replicas.
type ApplicationOperationSyncSourcesKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationOperationSyncSourcesKustomizeReplicasCount
type jsiiProxy_ApplicationOperationSyncSourcesKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationOperationSyncSourcesKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationOperationSyncSourcesKustomizeReplicasCount_FromNumber(value *float64) ApplicationOperationSyncSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationOperationSyncSourcesKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationOperationSyncSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationOperationSyncSourcesKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationOperationSyncSourcesKustomizeReplicasCount_FromString(value *string) ApplicationOperationSyncSourcesKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationOperationSyncSourcesKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationOperationSyncSourcesKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationOperationSyncSourcesKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

