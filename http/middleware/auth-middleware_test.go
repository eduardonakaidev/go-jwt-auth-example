package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardonakaidev/go-jwt-auth-example/utils"
)

func TestSecureHandler_WithValidToken(t *testing.T) {
	// Generate a valid token
	token, _, _ := utils.GenerateJwtTokenWithClaims("eduardo")

	req, err := http.NewRequest("GET", "/api/secure", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("X-Api-Token", token)

	rr := httptest.NewRecorder()
	handler := AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You are authenticated"))
	}))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
	}

	if rr.Body.String() != "You are authenticated" {
		t.Errorf("Unexpected response body: %s", rr.Body.String())
	}
}

func TestSecureHandler_WithInvalidToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/secure", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("X-Api-Token", "invalid.token.string")

	rr := httptest.NewRecorder()
	handler := AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You are authenticated"))
	}))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, status)
	}
}
