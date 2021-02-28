package httputil

import (
	"bytes"
	"compress/gzip"
	log "github.com/xiote/go-utils/chanlog"
	"golang.org/x/text/encoding/korean"
	tf "golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func DoOnEuckr(client *http.Client, req *http.Request) (src string, err error) {
	var body []byte
	if body, err = do(client, req); err != nil {
		return
	}
	var bufs bytes.Buffer
	wr := tf.NewWriter(&bufs, korean.EUCKR.NewDecoder())
	defer wr.Close()
	wr.Write(body)

	src = bufs.String()
	return
}

func Do(client *http.Client, req *http.Request, reqName string) (src string, err error) {

	var starttime time.Time
	go func() {
		starttime = time.Now()
		go log.Printf("[%s] [START]\n", reqName)
	}()

	var body []byte
	if body, err = do(client, req); err != nil {
		return
	}

	go func() {
		elasped := time.Since(starttime)
		l := int64(len(body) * 8)
		log.Printf("[%s] [END] [%s] [%d] [ %d Mbps ]\n", reqName, elasped, l, l*1000/elasped.Nanoseconds())
	}()

	src = string(body)
	return
}

func do(client *http.Client, req *http.Request) (body []byte, err error) {

	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	if reader, err = ContentDecodingReader(resp.Header.Get("Content-Encoding"), resp.Body); err != nil {
		return
	}

	if body, err = ioutil.ReadAll(reader); err != nil {
		return
	}

	return
}

func ContentDecodingReader(contentEncoding string, body io.ReadCloser) (reader io.ReadCloser, err error) {

	switch contentEncoding {
	case "gzip":
		reader, err = gzip.NewReader(body)
		if err != nil {
			return
		}
		defer reader.Close()
	default:
		reader = body
	}
	return
}
