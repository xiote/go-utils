package env

import (
	"os"
	"strconv"
	"testing"
)

func TestS(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"URL", "https://google.com"},
	}
	for _, c := range cases {
		os.Setenv(c.in, c.want)
		got := S(c.in)
		if got != c.want {
			t.Errorf("S(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestN(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"COUNT", 5},
	}
	for _, c := range cases {
		os.Setenv(c.in, strconv.Itoa(c.want))
		got := N(c.in)
		if got != c.want {
			t.Errorf("N(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
