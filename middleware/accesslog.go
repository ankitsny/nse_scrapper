package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// construct a custom http.ResponseWriter
// to fetch the status code

// CustomResponseWriter :
type CustomResponseWriter struct {
	w      http.ResponseWriter
	Status int
}

// WriteHeader :
func (cW *CustomResponseWriter) WriteHeader(status int) {
	cW.w.WriteHeader(status)
	cW.Status = status
}

// Header :
func (cW *CustomResponseWriter) Header() http.Header {
	return cW.w.Header()
}

// Write :
func (cW *CustomResponseWriter) Write(b []byte) (int, error) {
	return cW.w.Write(b)
}

// EnableAccessLog :
func EnableAccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cW := &CustomResponseWriter{
			w: w,
		}

		now := time.Now()
		next.ServeHTTP(cW, r)

		// TODO: use logger

		fmt.Printf("%v - %v - %v - %v - %v - %v\n", now.In(time.UTC), r.Method, r.RemoteAddr, r.URL.Path, time.Now().Sub(now), cW.Status)
	})
}
