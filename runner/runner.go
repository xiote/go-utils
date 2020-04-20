package runner

import (
	"reflect"
)

func NewRunner(any interface{}) Runner {
	return Runner{any}
}

type Runner struct {
	Any interface{}
}

func (r *Runner) Call(methodName string, args []string) (error, string) {

	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	result := reflect.ValueOf(r.Any).MethodByName(methodName).Call(inputs)
	err := result[0].Interface()
	if err == nil {
		return nil, result[1].Interface().(string)
	} else {
		return result[0].Interface().(error), result[1].Interface().(string)
	}

}
