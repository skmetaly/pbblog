package users

import (
	"github.com/jinzhu/gorm"
	//"github.com/satori/go.uuid"
	"github.com/davecgh/go-spew/spew"
	"github.com/skmetaly/pbblog/framework/database"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

//  [todo] Move this constants to config
const (
	passwordLength = 8
	hashCost       = 10
)

type User struct {
	gorm.Model
	Username  string
	Email     string
	Password  string
	createdAt time.Time
	updatedAt time.Time
}

type UserRepository struct {
	Db database.Database
}

func (uR *UserRepository) ByUsername(username string) User {

	user := User{}

	uR.Db.ORMConnection.Where("username = ?", username).First(&user)
	return user
}

//NewUser creates validates input and creates a new user instance
func (uR *UserRepository) NewUser(username string, email string, password string) (User, error) {
	user := User{
		Username: username,
		Email:    email,
	}

	if username == "" {
		return user, errNoUsername
	}

	duplicateUser := uR.ByUsername(username)

	if strings.Compare(duplicateUser.Username, username) == 0 {
		return user, errDuplicateUser
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
	user.Password = string(hashedPassword)

	return user, err
}
