package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// LoggingMiddleware logs request details
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Started %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)

		// Log the time taken to process the request
		fmt.Printf("Completed %s in %v\n", r.URL.Path, time.Since(start))
	})
}
