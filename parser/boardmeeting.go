package parser

import (
	"net/http"
	"nse_scrapper/models"
)

// ParseBoardMeeting :
func ParseBoardMeeting(resp *http.Response) (m []models.BoardMeeting) {

	defer resp.Body.Close()

	bm, _ := HTMLBodyToTextList(resp)

	if bm == nil || len(bm) < 3 {
		// Error
		return
	}

	bm = bm[3:]

	for i := 0; i < len(bm); i += 2 {
		m = append(m, models.BoardMeeting{
			Date:    bm[i],
			Purpose: bm[i+1],
		})
	}
	return
}
