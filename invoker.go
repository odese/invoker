package invoker

import (
	"reflect"
)

// Invoke ...
func Invoke(function interface{}, inputs []interface{}) (outputTypes []reflect.Value, outputValues []interface{}, err error) {
	f := reflect.ValueOf(function)
	inputValues := make([]reflect.Value, len(inputs))
	for i, input := range inputs {
		inputValues[i] = reflect.ValueOf(input)
	}
	outputTypes = f.Call(inputValues)
	for o := range outputTypes {
		outputValues = append(outputValues, outputTypes[o].Interface())
	}
	return
}

// FuncArguments ...
type FuncArguments struct {
	Function interface{}
	Inputs   []interface{}
}
