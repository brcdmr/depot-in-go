package api

import (
	"log"
	"net/http"
	"time"
)

// Logging HTTP requests
func LoggingMiddleware(target http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTimer := time.Now()
		target.ServeHTTP(w, r)
		log.Printf("url: %s - method: %s\t remoteIP: %s\t perf: %v", r.RequestURI, r.Method, r.RemoteAddr, time.Since(startTimer))

	})

}
