package euckr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Euckr_Euckr(t *testing.T) {
	euckrstr := Euckr(`abc`)
	assert.NotEmpty(t, euckrstr, "euckrstr")
}
