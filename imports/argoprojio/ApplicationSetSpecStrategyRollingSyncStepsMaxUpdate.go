package argoprojio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/fordneild/omega/imports/argoprojio/jsii"
)

type ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate interface {
	Value() interface{}
}

// The jsii proxy struct for ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate
type jsiiProxy_ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate struct {
	_ byte // padding
}

func (j *jsiiProxy_ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate_FromNumber(value *float64) ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate {
	_init_.Initialize()

	if err := validateApplicationSetSpecStrategyRollingSyncStepsMaxUpdate_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate_FromString(value *string) ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate {
	_init_.Initialize()

	if err := validateApplicationSetSpecStrategyRollingSyncStepsMaxUpdate_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate

	_jsii_.StaticInvoke(
		"argoprojio.ApplicationSetSpecStrategyRollingSyncStepsMaxUpdate",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

