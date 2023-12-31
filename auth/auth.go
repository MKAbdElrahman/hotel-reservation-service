package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const expirationPeriod = time.Hour * 24

func GenerateAuthToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"id":      userID,
		"expires": time.Now().Add(expirationPeriod).Unix(),
		"issued":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

func IsTokenNotExpired(token *jwt.Token) bool {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expirationTime := int64(claims["expires"].(float64))
		return time.Now().Unix() < expirationTime
	}
	return false
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
