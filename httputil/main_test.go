package httputil

import (
	"bytes"
	//"github.com/stretchr/testify/mock"
	"github.com/xiote/go-utils/httputil/mocks"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	mockHTTPClient := &mocks.HTTPClient{}
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"fields":[],"records":[]}`)))
	httpResponse := &http.Response{
		StatusCode: 200,
		Body:       r,
	}
	mockHTTPClient.On("Get", "http://abc").Return(httpResponse, nil).Once()

	cases := []struct {
		in   string
		want string
	}{
		{"http://abc", `{"fields":[],"records":[]}`},
	}

	client := Client{mockHTTPClient}

	for _, c := range cases {

		got := client.Get(c.in)

		if got != c.want {
			t.Errorf("Get(%q) == %q, want %q", c.in, got, c.want)
		}

		mockHTTPClient.AssertExpectations(t)
	}
}
