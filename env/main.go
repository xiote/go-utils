package env

import (
	"fmt"
	"os"
	"strconv"
)

func Getenv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Errorf("ENV[%s] is nil!", key))
	}
	return val
}

func S(name string) string {
	return os.Getenv(name)
}

func N(name string) int {
	i, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		panic(err)
	}
	return i
}
