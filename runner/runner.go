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

func (r *Runner) Call(methodName string, args []string) []reflect.Value {

	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	result := reflect.ValueOf(r.Any).MethodByName(methodName).Call(inputs)
	return result

	//err := result[0].Interface()
	//if err == nil {
	//    fmt.Println("No error returned by", m)
	//} else {
	//    fmt.Printf("Error calling %s: %v", m, err)
	//}
}
