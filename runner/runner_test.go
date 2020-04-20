package runner

import (
	"testing"
)

type YourT1 struct{}

func (y YourT1) Concat(input1 string, input2 string) (string, error) {
	return input1 + input2, nil
}

func (y YourT1) Get(input1 string) error {
	return nil
}

func TestCall(t *testing.T) {
	inMethodName := "Concat"
	inString1 := "A"
	inString2 := "B"
	want := "AB"

	r := NewRunner(YourT1{})
	result, _ := r.Call(inMethodName, []string{inString1, inString2})

	if result != want {
		t.Errorf("Call(%q) == %v, want %v", inMethodName, result, want)
	}
}

func TestCall2(t *testing.T) {
	inMethodName := "Get"

	r := NewRunner(YourT1{})
	result, _ := r.Call("Get", []string{"http://play.golang.org/?simple=1"})

	if result != "" {
		t.Errorf("Call(%q) == %v, want %v", inMethodName, result, "")
	}
}
