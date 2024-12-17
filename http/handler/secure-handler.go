package handler

import "net/http"

// SecureHandler is a handler for authenticated routes.
// This function is called only after passing through the AuthMiddleware,
// which validates the JWT token.
func SecureHandler(w http.ResponseWriter, r *http.Request) {
	// Respond to the client indicating successful authentication
	w.Write([]byte("You are authenticated"))
}
