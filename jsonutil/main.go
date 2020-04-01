package jsonutil

import (
	"github.com/PaesslerAG/jsonpath"
)

func Get(jv interface{}, path string) interface{} {
	result, err := jsonpath.Get(path, jv)
	if err != nil {
		panic(err)
	}
	return result
}
