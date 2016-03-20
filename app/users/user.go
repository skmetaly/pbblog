package users

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uuid.UUID
	Username       string
	Email          string
	Password       string
	HashedPassword string
}

//  [todo] Move this constants to config
const (
	passwordLength = 8
	hashCost       = 10
)

func NewUser(username string, email string, password string) (User, error) {
	user := User{
		Username: username,
		Email:    email,
		Password: password,
	}

	if username == "" {
		return user, errNoUsername
	}

	if password == "" {
		return user, errNoPassword
	}

	if email == "" {
		return user, errNoEmail
	}

	if len(password) < passwordLength {
		return user, errPasswordTooShort
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.HashedPassword = string(hashedPassword)

	user.ID = uuid.NewV4()

	return user, err
}
