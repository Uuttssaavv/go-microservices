package utilities

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

func SignToken(Data map[string]interface{}, expiresAt time.Duration) (string, error) {
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

func VerifyToken(accessToken string) (*jwt.Token, error) {
	jwtSecretKey := GodotEnv("JWT_SECRET")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}
