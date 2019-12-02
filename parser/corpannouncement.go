package parser

import (
	"net/http"
	"nse_scrapper/utils"
	"strings"
)

// ParseCorpAnnouncement :
func ParseCorpAnnouncement(resp *http.Response) (m []string) {
	defer resp.Body.Close()

	m, _ = HTMLBodyToTextList(resp)

	if m == nil || len(m) == 1 {
		return nil
	}

	m = utils.MapStr(m, func(s string) string {
		return strings.TrimSpace(s)
	})

	return
}
