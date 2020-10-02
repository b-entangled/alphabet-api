package alphabet

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// NewAlphabetRouter - a separate router for alphabet service
func NewAlphabetRouter(mux *mux.Router, base string) *mux.Router {
	alphabetRouter := mux.PathPrefix(base).Subrouter()
	alphabetRouter.HandleFunc("/validate", Alphabet).Methods("GET")
	// Use Log middleware to log duration of each request
	// Can add Authentication and Authorization
	alphabetRouter.Use(LogDurationMiddleWare)
	return alphabetRouter
}

// LogDurationMiddleWare - a log middleware that prints time taken to process a single request
func LogDurationMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(res, req)
		end := time.Since(start)
		log.Println("Total Time to process Request : URL - ", req.URL.String(), "Time - ", end)
	})
}
