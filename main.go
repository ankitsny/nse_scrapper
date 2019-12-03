package main

import (
	"fmt"
	"net/http"
	"nse_scrapper/app"
	"nse_scrapper/handlers"
	"nse_scrapper/middleware"
	"nse_scrapper/routes"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	if err := app.LoadConfig("env/dev.yaml"); err != nil {
		panic(err)
	}

	now := time.Now()

	router := mux.NewRouter()

	router.Use(middleware.EnableAccessLog)

	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, fmt.Sprintf(`
		Date: %v
		Version: %v
		`, now, app.Version))
	})

	router.HandleFunc("/nse", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	registerRoutes(router)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", app.Config.ServerPort),
		Handler: router,
	}

	// SPA Handler
	// TODO: Serve Static file using nginx or apache or CDN
	router.PathPrefix("/").Handler(handlers.SpaHandler{
		IndexPath:  "nse_view/build/index.html",
		StaticPath: "nse_view/build",
	})

	fmt.Println("Listening on Port ", app.Config.ServerPort)
	panic(server.ListenAndServe())

}

func registerRoutes(r *mux.Router) {
	nseHandler := handlers.NewNSEHandlers()
	routes.NewNSERoutes(r, nseHandler).RegisterRoutes()
}
