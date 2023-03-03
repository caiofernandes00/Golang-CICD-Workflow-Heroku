package util

import (
	"fmt"
	"reflect"
	"strings"
)

func Memoize(fn any) func(params ...any) any {
	cache := make(map[string]any)
	return func(params ...any) any {
		sb := strings.Builder{}
		for _, param := range params {
			sb.Write([]byte(fmt.Sprintf("%v", param)))
		}
		key := sb.String()
		if v, ok := cache[key]; ok {
			return v
		}
		result := reflect.ValueOf(fn).Call(GrabAllFunctionParamsToValue(params))
		cache[key] = result[0].Interface()
		return cache[key]
	}
}
