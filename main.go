package main

import (
	"github.com/eduardonakaidev/go-jwt-auth-example/http/handler"
	"github.com/eduardonakaidev/go-jwt-auth-example/http/middleware"
	"net/http"
)

func main() {
	// Route for user authentication (login)
	// Example: POST http://localhost:3000/api/auth
	http.HandleFunc("POST /api/auth", handler.LoginHandler)

	// Public route that does not require authentication
	// Example: GET http://localhost:3000/api/public
	http.HandleFunc("GET /api/public", handler.PublicHandler)

	// Secure route that requires authentication
	// The AuthMiddleware validates the JWT token before accessing SecureHandler
	// Example: GET http://localhost:3000/api/secure
	http.HandleFunc("GET /api/secure", middleware.AuthMiddleware(handler.SecureHandler))

	// Start the HTTP server and listen on port 3000
	// Accessible at http://localhost:3000
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic("Server failed to start: " + err.Error())
	}
}
