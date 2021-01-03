package jsonutil

import (
	"encoding/json"
	"testing"
)

func TestGet(t *testing.T) {
	cases := []struct {
		in1  string
		in2  string
		want string
	}{
		{`
        {"@type":"response","@service":"vcenter","@version":"1.0.0","result":{"aptNo":7743,"aptType":"A01","aptName":"현대","aptTypeName":"아파트","cortarNo":"4812312600","addr":"경상남도 창원시 성산구 반림동 3-1","city":"경상남도","dvsn":"창원시 성산구","sec":"반림동","dtlAddr":"3-1","totHsehCnt":1200,"dongCount":13,"planTypeCount":7}}
        `,
			"$.result.cortarNo", "4812312600"},
	}

	for _, c := range cases {
		jv := interface{}(nil)
		json.Unmarshal([]byte(c.in1), &jv)
		got := Get(jv, c.in2)
		if got != c.want {
			t.Errorf("Get(%q, %q) == %q, not %q", c.in1, c.in2, got, c.want)
		}
	}
}
