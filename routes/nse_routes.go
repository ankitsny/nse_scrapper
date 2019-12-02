package routes

import (
	"net/http"
	"nse_scrapper/handlers"

	"github.com/gorilla/mux"
)

type nseRoutes struct {
	router  *mux.Router
	handler handlers.INSEHandlers
}

// INSERoutes :
type INSERoutes interface {
	RegisterRoutes()
}

// NewNSERoutes :
func NewNSERoutes(r *mux.Router, h handlers.INSEHandlers) INSERoutes {
	return &nseRoutes{
		handler: h,
		router:  r,
	}
}

func (nr *nseRoutes) RegisterRoutes() {
	nr.router.HandleFunc("/company", nr.handler.SearchCompanyName).Methods(http.MethodGet)
	nr.router.HandleFunc("/company/details", nr.handler.GetCompanyDetails).Methods(http.MethodGet)
}
