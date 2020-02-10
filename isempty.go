package strutil

import (
	"strings"
	"unicode"
)

func IsEmpty(s string) bool {
	if s == "" {
		return true
	} else {
		return false
	}
}

func IsSpaceOrEmpty(s string) bool {
	return IsEmptyS(s)
}

// check that string is empty or space.
func IsEmptyS(s string) bool {
	return (len(strings.TrimSpace(s)) == 0)
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func IsSpace(s string) bool {
	for _, char := range s {
		if !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

func IsNotSpace(s string) bool {
	return !IsSpace(s)
}
