package utils

import (
	"testing"
)

func TestGenerateJwtTokenWithClaims(t *testing.T) {
	username := "testuser"

	token, claims, err := GenerateJwtTokenWithClaims(username)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check if claims are as expected
	if claims["sub"] != username {
		t.Errorf("Expected sub claim to be %s, got %s", username, claims["sub"])
	}

	if claims["exp"] == nil {
		t.Errorf("Expected exp claim to be set, got nil")
	}

	if token == "" {
		t.Errorf("Token should not be empty")
	}
}