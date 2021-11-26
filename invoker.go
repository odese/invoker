package invoker

import (
	"reflect"
)

// FuncArguments ...
type FuncArguments struct {
	Function   interface{}
	Parameters []interface{}
}

// Invoker ...
func Invoker(m map[string]FuncArguments, name string, params []interface{}) (output []reflect.Value, err error) {
	f := reflect.ValueOf(m[name].Function)
	input := make([]reflect.Value, len(params))
	for p, param := range params {
		input[p] = reflect.ValueOf(param)
	}
	output = f.Call(input)
	return
}

// Invoke ...
func Invoke(function interface{}, params []interface{}) (outputs []reflect.Value, err error) {
	f := reflect.ValueOf(function)
	inputs := make([]reflect.Value, len(params))
	for p, param := range params {
		inputs[p] = reflect.ValueOf(param)
	}
	outputs = f.Call(inputs)
	return
}