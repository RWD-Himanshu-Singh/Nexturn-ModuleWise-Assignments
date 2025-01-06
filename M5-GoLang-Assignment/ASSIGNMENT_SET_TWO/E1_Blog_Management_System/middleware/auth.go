package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// BasicAuthMiddleware is the middleware function for basic authentication
func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Decode the Base64-encoded credentials
		credentials := strings.TrimPrefix(authHeader, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(credentials)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Split username and password
		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		username, password := parts[0], parts[1]

		// Example: Validate username and password (use secure validation in a real-world app)
		if username != "admin" || password != "password" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authentication passes, call the next handler
		next.ServeHTTP(w, r)
	})
}
