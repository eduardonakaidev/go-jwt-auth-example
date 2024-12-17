package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestValidateToken_ValidToken(t *testing.T) {
	username := "testuser"

	// Generate a valid token
	token, _, err := GenerateJwtTokenWithClaims(username)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Validate the token
	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check claims
	if claims["sub"] != username {
		t.Errorf("Expected sub claim to be %s, got %s", username, claims["sub"])
	}
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	// Manually create an expired token
	claims := jwt.MapClaims{
		"sub": "testuser",
		"exp": time.Now().Add(-time.Second * 10).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(JWT_SIGNING_KEY))

	// Validate the token
	_, err := ValidateToken(tokenString)
	if err == nil {
		t.Error("Expected error for expired token, got none")
	}
}

func TestValidateToken_InvalidToken(t *testing.T) {
	// Use an invalid token string
	invalidToken := "invalid.token.string"

	_, err := ValidateToken(invalidToken)
	if err == nil {
		t.Error("Expected error for invalid token, got none")
	}
}