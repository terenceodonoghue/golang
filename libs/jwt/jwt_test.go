package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/terenceodonoghue/golang/libs/assert"
)

func TestCreateToken(t *testing.T) {
	type err struct {
		want error
	}

	errs := []err{
		{want: nil},
	}

	for _, test := range errs {
		t.Setenv("JWT_SECRET", "test-key")
		_, got := CreateToken(current)
		assert.ErrorIs(t, got, test.want)
	}
}

func TestVerifyToken(t *testing.T) {
	type exp struct {
		with time.Time
		want error
	}

	type key struct {
		with string
		want error
	}

	exps := []exp{
		{with: current, want: nil},
		{with: expired, want: jwt.ErrTokenExpired},
	}

	keys := []key{
		{with: "test-key", want: nil},
		{with: "fake-key", want: jwt.ErrSignatureInvalid},
	}

	for _, test := range exps {
		jwt := createToken(t, test.with)
		got := VerifyToken(jwt)
		assert.ErrorIs(t, got, test.want)
	}

	for _, test := range keys {
		jwt := createToken(t, current)
		t.Setenv("JWT_SECRET", test.with)
		got := VerifyToken(jwt)
		assert.ErrorIs(t, got, test.want)
	}
}

func createToken(t *testing.T, exp time.Time) string {
	t.Setenv("JWT_SECRET", "test-key")
	jwt, _ := CreateToken(exp)
	return jwt
}

var current time.Time = time.Date(2025, time.January,
	00, 00, 00, 00, 0, time.UTC)

var expired time.Time = time.Now().Add(-1 * time.Hour)
