package jwt

import (
	"fmt"
	"time"

	"github.com/ahmadalaik/be-invitation/config"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(config.LoadConfig().JWTSecret)

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expirationTime,
	})

	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateToken(tokenStr string) (string, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			// Accept only HMAC method (HS256, HS384, HS512)
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	return claims["username"].(string), nil
}
