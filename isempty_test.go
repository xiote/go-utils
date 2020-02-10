package strutil

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		inStr       string
		wantIsEmpty bool
	}{
		{"NotEmptyString", false},
		{"", true},
	}
	for _, c := range cases {
		gotIsEmpty := IsEmpty(c.inStr)
		if gotIsEmpty != c.wantIsEmpty {
			t.Errorf("IsEmpty(%q) == %v, want %v", c.inStr, gotIsEmpty, c.wantIsEmpty)
		}
	}
}

func TestIsSpaceOrEmpty(t *testing.T) {
	cases := []struct {
		inStr              string
		wantIsSpaceOrEmpty bool
	}{
		{"NotSpaceOrEmptyString", false},
		{" NotSpaceOrEmptyString ", false},
		{"", true},
		{"  ", true},
	}
	for _, c := range cases {
		gotIsSpaceOrEmpty := IsSpaceOrEmpty(c.inStr)
		if gotIsSpaceOrEmpty != c.wantIsSpaceOrEmpty {
			t.Errorf("IsSpaceOrEmpty(%q) == %v, want %v", c.inStr, gotIsSpaceOrEmpty, c.wantIsSpaceOrEmpty)
		}
	}
}

func TestIsEmptyS(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"NotSpaceOrEmptyString", false},
		{" NotSpaceOrEmptyString ", false},
		{"", true},
		{"  ", true},
	}
	for _, c := range cases {
		got := IsEmptyS(c.in)
		if got != c.want {
			t.Errorf("IsEmptyS(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
