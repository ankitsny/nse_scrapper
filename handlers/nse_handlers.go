package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nse_scrapper/parser"
	"nse_scrapper/utils"
)

type nseHandlers struct {
	// service
	// utils
}

// INSEHandlers :
type INSEHandlers interface {
	SearchCompanyName(w http.ResponseWriter, r *http.Request)
	GetCompanyDetails(w http.ResponseWriter, r *http.Request)
}

// NewNSEHandlers : create instance of nse handlers
func NewNSEHandlers() INSEHandlers {
	return &nseHandlers{}
}

func (ns *nseHandlers) SearchCompanyName(w http.ResponseWriter, r *http.Request) {

	req := utils.HTTPRequest{
		URL:   "https://www.nseindia.com/corporates/common/getCompanyListMktTracker.jsp",
		Query: r.URL.Query(),
	}

	resp, err := req.Do(http.MethodGet)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

}

const (
	boardMeeting int = iota
	corpAction
	corpAnnounce
	corpInfo
)

var urls = []string{
	"https://www.nseindia.com/marketinfo/companyTracker/boardMeeting.jsp",
	"https://www.nseindia.com/marketinfo/companyTracker/corpAction.jsp",
	"https://www.nseindia.com/marketinfo/companyTracker/corpAnnounce.jsp",
	"https://www.nseindia.com/marketinfo/companyTracker/compInfo.jsp",
}

func (ns *nseHandlers) GetCompanyDetails(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	query.Add("cName", "cmtracker_nsedef.css")
	query.Add("series", "EQ")

	resp := utils.ParallelGET(urls, query, 4)

	// fmt.Printf("%+v\n", resp)

	cA := func() interface{} {
		return parser.ParseCorpAction(resp[corpAction].Response)
	}
	bM := func() interface{} {
		return parser.ParseBoardMeeting(resp[boardMeeting].Response)
	}
	cAn := func() interface{} {
		return parser.ParseCorpAnnouncement(resp[corpAnnounce].Response)
	}
	cI := func() interface{} {
		return parser.ParseCorpInfo(resp[corpInfo].Response)
	}

	results := utils.ExecParallel(4, cA, bM, cAn, cI)

	cResp := map[string]interface{}{
		"corpAction":       results[0].Data,
		"boardMeeting":     results[1].Data,
		"corpAnnouncement": results[2].Data,
		"corpInfo":         results[3].Data,
	}

	cRespB, _ := json.Marshal(cResp)

	// copyHeader(w.Header(), _resp.Header)
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(_resp.StatusCode)
	// io.Copy(w,)
	w.WriteHeader(200)
	w.Write(cRespB)

}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
