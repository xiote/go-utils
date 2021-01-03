package waiter

import (
	"github.com/stretchr/testify/assert"
	// "github.com/xiote/ticketing/ticketinginfo"
	// "encoding/json"
	// "fmt"
	"testing"
	// . "github.com/xiote/go-utils/keyvalueutil"
	// . "github.com/xiote/go-utils/euckr"
	// "golang.org/x/net/publicsuffix"
	// "io/ioutil"
	// "log"
	// "math/rand"
	// "net/http"
	// "net/http/cookiejar"
	// "net/http/httptest"
	// "net/url"
	// "os"
	// "strconv"
	"context"
	"time"
)

func Test_Waiter_WaitTill2(t *testing.T) {
	now := time.Now()
	then := now.Add(5 * time.Second)

	WaitTill2(now, then)
}

func Test_Waiter_WaitUntilOk(t *testing.T) {
	chkfn := func(timeoutctx context.Context) (isok bool, err error) {
		isok = true
		err = nil
		return
	}

	err := WaitUntilOk(chkfn, 1*time.Second)
	assert.NoError(t, err, "err")
}

func Test_Waiter_WaitUntilOkTimeout(t *testing.T) {
	chkfn := func(timeoutctx context.Context) (isok bool, err error) {
		isok = false
		err = nil
		return
	}

	err := WaitUntilOk(chkfn, 1*time.Second)
	assert.Error(t, err, "err")
}
