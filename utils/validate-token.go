package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}

		return []byte(JWT_SIGNING_KEY), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}
	// check token valdity
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	// get claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	// check expire time
	expValue, ok := claims["exp"]
	if !ok {
		return nil, fmt.Errorf("exp field not found in token")
	}

	expires, ok := expValue.(float64)
	if !ok {
		return nil, fmt.Errorf("invalid type for exp field")
	}

	if time.Now().Unix() > int64(expires) {
		return nil, fmt.Errorf("token expired")
	}
	return claims, nil
}
