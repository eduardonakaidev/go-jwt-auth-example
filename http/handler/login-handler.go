package handler

import (
	"encoding/json"
	"net/http"

	"github.com/eduardonakaidev/utils"
)

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token  string      `json:"token"`
	Claims interface{} `json:"claims"` // Include full claims
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginParams LoginRequest

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&loginParams)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusBadRequest)
		return
	}

	// Validate username and password (simulated)
	if loginParams.UserName == "eduardo" && loginParams.Password == "123456" {
		// Generate the token and claims
		tokenString, claims, err := utils.GenerateJwtTokenWithClaims(loginParams.UserName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create the response with token and claims
		res := LoginResponse{
			Token:  tokenString,
			Claims: claims,
		}

		// Return the JSON response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	// Invalid credentials
	http.Error(w, "invalid credentials", http.StatusBadRequest)
}
