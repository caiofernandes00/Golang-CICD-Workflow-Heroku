package util

import "reflect"

func GrabAllFunctionParamsToValue(params []any) []reflect.Value {
	var values []reflect.Value
	for _, param := range params {
		values = append(values, reflect.ValueOf(param))
	}
	return values
}
