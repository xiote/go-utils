package stringutil

import (
	"testing"
)

func TestPadLeft(t *testing.T) {
	cases := []struct {
		in1  string
		in2  string
		in3  int
		want string
	}{
		{"12", "a", 4, "aa12"},
	}
	for _, c := range cases {
		got := PadLeft(c.in1, c.in2, c.in3)
		if got != c.want {
			t.Errorf("PadLeft(%q, %q, %d) == %q, want %q", c.in1, c.in2, c.in3, got, c.want)
		}
	}
}

func TestPadRight(t *testing.T) {
	cases := []struct {
		in1  string
		in2  string
		in3  int
		want string
	}{
		{"12", "k", 4, "12kk"},
	}
	for _, c := range cases {
		got := PadRight(c.in1, c.in2, c.in3)
		if got != c.want {
			t.Errorf("PadRight(%q, %q, %d) == %q, want %q", c.in1, c.in2, c.in3, got, c.want)
		}
	}
}
