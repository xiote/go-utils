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

func (r *Runner) Call(methodName string, args []string) (string, error) {

	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	result := reflect.ValueOf(r.Any).MethodByName(methodName).Call(inputs)
	if len(result) == 1 {
		err := result[0].Interface()
		if err == nil {
			return "", nil
		} else {
			return "", result[0].Interface().(error)
		}
	} else if len(result) >= 2 {
		err := result[1].Interface()
		if err == nil {
			return result[0].Interface().(string), nil
		} else {
			return result[0].Interface().(string), result[1].Interface().(error)
		}
	}

	return "", nil
}
