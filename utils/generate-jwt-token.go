package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT_SIGNING_KEY is the secret key used to sign the JWT.
const JWT_SIGNING_KEY = "123"

// GenerateJwtTokenWithClaims generates a JWT token with custom claims.
// It returns the token string, the claims, and any error that occurs during the process.
func GenerateJwtTokenWithClaims(userName string) (string, jwt.MapClaims, error) {
	// Current time (issued at)
	now := time.Now()
	// Token expiration time (current time + 15 seconds)
	expires := now.Add(time.Second * 15).Unix()

	// Define the JWT claims
	claims := jwt.MapClaims{
		"sub": userName,   // Subject (user identifier)
		"iat": now.Unix(), // Issued at time (timestamp)
		"exp": expires,    // Expiration time (timestamp)
	}

	// Create a new token using the HMAC SHA256 signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret signing key
	tokenString, err := token.SignedString([]byte(JWT_SIGNING_KEY))
	if err != nil {
		// Return error if signing fails
		return "", nil, err
	}

	// Return the signed token string, the claims, and nil error
	return tokenString, claims, nil
}
