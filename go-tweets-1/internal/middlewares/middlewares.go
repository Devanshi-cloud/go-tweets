package middlewares

import (
	"net/http"
	"strings"
)

// AuthMiddleware checks for the presence of a valid JWT token in the request headers.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Here you would typically validate the token and extract claims.
		// For now, we will just call the next handler.
		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware logs the details of each incoming request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request details (method, URL, etc.)
		// This is a placeholder for actual logging logic.
		next.ServeHTTP(w, r)
	})
}