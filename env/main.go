package env

import (
	"os"
	"strconv"
)

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
