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
		t.Setenv("JWT_SECRET_KEY", "test-key")
		_, got := CreateToken(current)
		assert.ErrorIs(t, got, test.want)
	}
}

func TestVerifyToken(t *testing.T) {
	type exp struct {
		with time.Time
		want error
	}

	type sig struct {
		with string
		want error
	}

	exps := []exp{
		{with: current, want: nil},
		{with: expired, want: jwt.ErrTokenExpired},
	}

	sigs := []sig{
		{with: "test-key", want: nil},
		{with: "fake-key", want: jwt.ErrSignatureInvalid},
	}

	for _, test := range exps {
		t.Setenv("JWT_SECRET_KEY", "test-key")
		jwt, _ := CreateToken(test.with)
		got := VerifyToken(jwt)
		assert.ErrorIs(t, got, test.want)
	}

	for _, test := range sigs {
		t.Setenv("JWT_SECRET_KEY", test.with)
		jwt, _ := CreateToken(current)
		t.Setenv("JWT_SECRET_KEY", "test-key")
		got := VerifyToken(jwt)
		assert.ErrorIs(t, got, test.want)
	}
}

var current time.Time = time.Date(2025, time.January,
	00, 00, 00, 00, 0, time.UTC)

var expired time.Time = time.Now().Add(-1 * time.Hour)
