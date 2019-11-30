package handlers

import (
	"fmt"
	"io"
	"net/http"
	"nse_scrapper/utils"
)

type nseHandlers struct {
	// service
	// utils
}

// INSEHandlers :
type INSEHandlers interface {
	SearchCompanyName(w http.ResponseWriter, r *http.Request)
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

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
