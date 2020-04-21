package runner

import (
	"fmt"
	"reflect"
)

func NewRunner(any interface{}) Runner {
	return Runner{any}
}

type Runner struct {
	Any interface{} // pointer
}

func (r *Runner) Call(methodName string, args []string) (string, error) {

	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	f := reflect.ValueOf(r.Any).MethodByName(methodName)
	fmt.Printf("%v", f)
	result := f.Call(inputs)
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

func (r *Runner) Call2(infos ...string) (string, error) {

	methodName := infos[0]
	infos = append(infos[:0], infos[1:]...) // remove 0 idx, https://stackoverflow.com/questions/25025409/delete-element-in-a-slice
	inputs := make([]reflect.Value, len(infos))
	for i, _ := range infos {
		inputs[i] = reflect.ValueOf(infos[i])
	}
	f := reflect.ValueOf(r.Any).MethodByName(methodName)
	fmt.Printf("%v", f)
	result := f.Call(inputs)
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
