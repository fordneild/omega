package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

// Number of replicas.
type ApplicationOperationSyncSourceKustomizeReplicasCount interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationOperationSyncSourceKustomizeReplicasCount
type jsiiProxy_ApplicationOperationSyncSourceKustomizeReplicasCount struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationOperationSyncSourceKustomizeReplicasCount) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationOperationSyncSourceKustomizeReplicasCount_FromNumber(value *float64) ApplicationOperationSyncSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationOperationSyncSourceKustomizeReplicasCount_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationOperationSyncSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationOperationSyncSourceKustomizeReplicasCount",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationOperationSyncSourceKustomizeReplicasCount_FromString(value *string) ApplicationOperationSyncSourceKustomizeReplicasCount {
	_init_.Initialize()

	if err := validateApplicationOperationSyncSourceKustomizeReplicasCount_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationOperationSyncSourceKustomizeReplicasCount

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationOperationSyncSourceKustomizeReplicasCount",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

