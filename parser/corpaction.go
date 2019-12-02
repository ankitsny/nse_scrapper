package parser

import (
	"net/http"
	"nse_scrapper/models"
)

// ParseCorpAction :
func ParseCorpAction(resp *http.Response) (m []models.CorpAction) {
	defer resp.Body.Close()

	cA, _ := HTMLBodyToTextList(resp)

	if cA == nil || len(cA) < 2 {
		// Error
		return
	}

	cA = cA[2:]

	for i := 0; i < len(cA); i += 2 {
		m = append(m, models.CorpAction{
			Date:    cA[i],
			Purpose: cA[i+1],
		})
	}
	return
}
