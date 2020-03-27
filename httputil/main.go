package httputil

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

type Client struct {
	HTTPClient HTTPClient
}

func (c *Client) Get(url string) string {
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if status := resp.StatusCode; status != http.StatusOK {
		fmt.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	return string(body)
}
