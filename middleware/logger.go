package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received for: %s %s",r.Method,r.URL.Path);
		next.ServeHTTP(w,r)
		log.Printf("Request processed for: %s %s",r.Method,r.URL.Path)
	})
}