package jwt

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestCreateToken(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "test-key")

	want := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzU2MDM4MDB9.cltD4WgT33Cc2divxm9yCCupThVn7aYCpIer_GtBPvU"
	if got, err := CreateToken(current); got != want || err != nil {
		t.Fatalf(`CreateToken() got %q, %v, wanted %q, <nil>`, got, err, want)
	}
}

func TestVerifyToken(t *testing.T) {
	t.Setenv("JWT_SECRET_KEY", "test-key")

	type test struct {
		with func() time.Time
		want error
	}

	tests := []test{
		{with: current, want: nil},
		{with: expired, want: jwt.ErrTokenExpired},
	}

	for _, test := range tests {
		token, _ := CreateToken(test.with)
		if got := VerifyToken(token); !errors.Is(got, test.want) {
			t.Fatalf(`VerifyToken() got %v, wanted %v`, got, test.want)
		}
	}
}

func current() time.Time {
	return time.Date(2025, time.January,
		00, 00, 00, 00, 0, time.UTC)
}

func expired() time.Time {
	return time.Now().Add(-1 * time.Hour)
}
