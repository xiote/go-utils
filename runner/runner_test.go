package runner

import (
	"testing"
)

type YourT1 struct{}

func (y YourT1) Concat(input1 string, input2 string) (error, string) {
	return nil, input1 + input2
}

func TestCall(t *testing.T) {
	inMethodName := "Concat"
	inString1 := "A"
	inString2 := "B"
	want := "AB"

	r := NewRunner(YourT1{})
	_, result := r.Call(inMethodName, []string{inString1, inString2})

	if result != want {
		t.Errorf("Call(%q) == %v, want %v", inMethodName, result[1], want)
	}
}
