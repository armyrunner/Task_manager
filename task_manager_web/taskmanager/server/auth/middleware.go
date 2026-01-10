package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/armyrunner/task_manager/models"
)

// ContextKey type for context values
type ContextKey string

const (
	// UserContextKey is the key for storing user info in context
	UserContextKey ContextKey = "user"
)

// AuthMiddleware validates JWT tokens and adds user info to request context
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SetHeaders(w)

		// Handle OPTIONS preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Check Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format. Use: Bearer <token>", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// Validate token
		claims, err := ValidateAccessToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Add user info to context
		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// GetUserFromContext retrieves the user claims from the request context
func GetUserFromContext(r *http.Request) *models.JWTClaims {
	if claims, ok := r.Context().Value(UserContextKey).(*models.JWTClaims); ok {
		return claims
	}
	return nil
}
