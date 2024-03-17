package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the path hit
		log.Println("Path hit:", r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
