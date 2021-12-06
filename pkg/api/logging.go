package api

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(target http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTimer := time.Now()
		target.ServeHTTP(w, r)
		log.Printf("url: %s - method: %s\t remoteIP: %s\t perf: %v", r.RequestURI, r.Method, r.RemoteAddr, time.Since(startTimer))

	})

}

// func LoggingMiddlewareWtRecover(logger log.Logger) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		fn := func(w http.ResponseWriter, r *http.Request) {
// 			defer func() {
// 				if err := recover(); err != nil {
// 					w.WriteHeader(http.StatusInternalServerError)

// 					log.Printf("Recoverd Error: %s ", err)
// 				}
// 			}()

// 			startTimer := time.Now()
// 			next.ServeHTTP(w, r)
// 			log.Printf("url: %s - method: %s\t remoteIP: %s\t perf: %v", r.RequestURI, r.Method, r.RemoteAddr, time.Since(startTimer))

// 		}

// 		return http.HandlerFunc(fn)
// 	}
// }
