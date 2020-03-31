package env

import (
	"os"
	"strconv"
)

func Getenv(name string) string {
        value := os.Getenv(string)
        if value == nil {
                fmt.Errorf("ENV[%s] is nil!", name)                                                                                                       
        }
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
