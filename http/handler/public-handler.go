package handler

import "net/http"

// PublicHandler is a handler for public routes.
// This function allows access to anyone without authentication.
func PublicHandler(w http.ResponseWriter, r *http.Request) {
	// Send a plain text response to the client
	w.Write([]byte("Everyone can view this endpoint"))
}
