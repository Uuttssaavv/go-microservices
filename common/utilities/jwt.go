package utilities

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Sign(Data map[string]interface{}, expiresAt time.Duration) (string, error) {
	expiringTime := time.Now().Add(expiresAt)
	secretKey := GodotEnv("JWT_SECRET")

	claims := jwt.MapClaims{}
	claims["expiresAt"] = expiringTime
	claims["authorization"] = true

	for key, value := range Data {
		claims[key] = value
	}

	secret := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := secret.SignedString([]byte(secretKey))

	if err == nil {
		return accessToken, nil
	}
	return accessToken, err

}
