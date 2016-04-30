package hash

import (
	"golang.org/x/crypto/bcrypt"
)

//HashBytes hashes a string using bcrypt with default cost
func HashBytes(stringPassword []byte) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(stringPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}

//CompareWithHash compares a hash and a string password
func CompareWithHash(hashPassword []byte, stringPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(stringPassword))

	if err == nil {
		return true
	}

	return false
}

func CreateFromPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword)
}
