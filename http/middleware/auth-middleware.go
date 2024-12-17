package middleware

import (
	"fmt"
	"net/http"

	"github.com/eduardonakaidev/go-jwt-auth-example/utils"
)

// AuthMiddleware is a middleware function for JWT token validation.
// It intercepts HTTP requests, checks for a valid token in the request header,
// and allows the request to proceed only if the token is valid.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the token from the request header
		token := r.Header.Get("X-Api-Token")
		
		// If the token is not provided, return an unauthorized error
		if len(token) == 0 {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		// Validate the token using the ValidateToken function
		claims, err := utils.ValidateToken(token)
		if err != nil {
			// If the token is invalid, return an unauthorized error
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		// Print the claims for debugging purposes
		fmt.Println(claims)

		// Call the next handler if the token is valid
		next.ServeHTTP(w, r)
	}
}
