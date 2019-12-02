package parser

import (
	"net/http"
	"strings"
)

// ParseCorpInfo :
func ParseCorpInfo(resp *http.Response) (m map[string]string) {

	defer resp.Body.Close()

	m = make(map[string]string)

	cI, _ := HTMLBodyToTextList(resp)

	if cI == nil || len(cI) <= 1 {
		return nil
	}

	m["Company Name"] = cI[0]
	cI = cI[1 : len(cI)-1]

	for i := 0; i < len(cI); i++ {
		if strings.Contains(cI[i], ":") {
			t := strings.Split(cI[i], ":")
			m[strings.TrimSpace(t[0])] = strings.TrimSpace(t[1])
		}
	}
	return
}
