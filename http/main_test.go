package http

import (
	"testing"
)

func TestGet(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"http://ipcheck.re4u.co.kr", "112.223.47.109"},
		{"https://www.data.go.kr/pubr/use/pra/IrosPblonsipSvcReqst/downloadUrlXmlToJson.do?fileName=%EA%B1%B4%EC%B6%95%EB%AC%BC%EC%86%8C%EC%9C%A0%20%EC%A0%95%EB%B3%B4%EC%A1%B0%ED%9A%8C&oagUseAt=Y&url=http://apis.data.go.kr/1611000/OwnerInfoService/getArchitecturePossessionInfo?ServiceKey=dmaUZsmx2hjfF5JA%2FmhuUQjeLdmgRNLmCV8ix1P4S4vNQI1frJzRhzvMDJUN3iKcpkHm3JzckCT8l421i3sK%2FA%3D%3D&sigungu_cd=41410&bjdong_cd=10400&bun=1156&ji=0001&numOfRows=1000&pageNo=3", `{"fields":[],"records":[]}`},
	}

	for _, c := range cases {
		got := Get(c.in)
		if got != c.want {
			t.Errorf("Get(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
