package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(exp time.Time) (string, error) {
	key := getJwtSecret()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": exp.Unix(),
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

func getJwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func withHMAC(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	key := getJwtSecret()
	return key, nil
}
