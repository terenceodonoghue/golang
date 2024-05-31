package assert

import (
	"errors"
	"testing"
)

func ErrorIs(t *testing.T, got error, want error) {
	if !errors.Is(got, want) {
		t.Errorf(`got %v, wanted %v`, got, want)
	}
}
