package euckr

import (
	"bytes"
	"golang.org/x/text/encoding/korean"
	tf "golang.org/x/text/transform"
)

func Euckr(source string) string {
	var bufs bytes.Buffer
	tf.NewWriter(&bufs, korean.EUCKR.NewEncoder()).Write([]byte(source))
	return bufs.String()
}
