package httputil

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	// "net/http/httptest"
	"testing"
	"time"
)

func Test_Httputil_Do(t *testing.T) {

	var req *http.Request
	var err error
	url := "https://poticket.interpark.com/Book/null.html"
	// url := "https://google.com"

	// --------------------------------
	// ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, `{"common":{"messageId":null,"message":"success","requestUri":"/v1/goods/20010593/playSeq","gtid":"","timestamp":"20210101165750","internalHttpStatusCode":200},"data":[{"playSeq":"001","playDate":"20210105","playTime":"1930","bookableDate":"202012311600","bookingEndDate":"202101051830","cancelableDate":"202101042359","remainSeat":null,"casting":null,"limitMaxStayDate":null},{"playSeq":"002","playDate":"20210106","playTime":"1930","bookableDate":"202012311600","bookingEndDate":"202101061830","cancelableDate":"202101052359","remainSeat":null,"casting":null,"limitMaxStayDate":null}]}`)
	// }))
	// defer ts.Close()
	// url = ts.URL
	// --------------------------------

	req, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err, "err")

	src, err := Do(&http.Client{}, req, "Google")

	assert.NoError(t, err, "err")
	assert.NotEmpty(t, src, "src")

	// fmt.Println(src)
	<-time.After(time.Second)

}
