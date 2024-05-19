package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "JWT_SECRET_KEY"

func CreateToken(now func() time.Time) (string, error) {
	key := []byte(os.Getenv(secretKey))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": now().Add(10 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, withHMAC)
	return err
}

func withHMAC(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	key := []byte(os.Getenv(secretKey))
	return key, nil
}
