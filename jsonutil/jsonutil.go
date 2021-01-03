package jsonutil

import (
	"encoding/json"
	"github.com/PaesslerAG/jsonpath"
	"io/ioutil"
)

func Get(jv interface{}, path string) interface{} {
	result, err := jsonpath.Get(path, jv)
	if err != nil {
		panic(err)
	}
	return result
}

func Unmarshal(filename string, v interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), v)
	if err != nil {
		return err
	}

	return nil
}

func MustUnmarshal(filename string, v interface{}) {
	if err := Unmarshal(filename, v); err != nil {
		panic(err)
	}
}
