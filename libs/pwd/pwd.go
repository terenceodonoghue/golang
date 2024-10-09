package pwd

import (
	"golang.org/x/crypto/bcrypt"
)

func Compare(hash []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

func HashAndSalt(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	return hash, err
}
