package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	mux.HandleFunc("/nse", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	server := http.Server{
		Addr:    ":4000",
		Handler: mux,
	}
	fmt.Println("Yeah")
	panic(server.ListenAndServe())

}
