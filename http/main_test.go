package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

var GetMock func(url string) (resp *http.Response, err error)

type MockHttpClient struct{}

func (m MockHttpClient) Get(url string) (resp *http.Response, err error) {
	return GetMock(url)
}

func TestGet(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"http://anywhere", `{"fields":[],"records":[]}`},
	}

	for _, c := range cases {
		GetMock = func(url string) (resp *http.Response, err error) {
			json := `{"fields":[],"records":[]}`
			r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		}
		got := Get(c.in, MockHttpClient{})
		if got != c.want {
			t.Errorf("Get(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
