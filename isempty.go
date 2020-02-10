package strutil

import (
	"strings"
)

func IsEmpty(s string) bool {
	return (len(s) == 0)
}

func IsSpaceOrEmpty(s string) bool {
	return IsEmptyS(s)
}

func IsEmptyS(s string) bool {
	return (len(strings.TrimSpace(s)) == 0)
}
