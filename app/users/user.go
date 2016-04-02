package users

import (
	"github.com/jinzhu/gorm"
	//"github.com/satori/go.uuid"
	//"github.com/davecgh/go-spew/spew"
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

//Save persists to database a User type
func (uR *UserRepository) Save(user User) {
	uR.Db.ORMConnection.Create(&user)
}

//ById returns a User object that has the provided id
func (uR *UserRepository) ById(usernameId int) User {

	user := User{}

	uR.Db.ORMConnection.Where("id = ?", usernameId).First(&user)
	return user
}

//ByEmail returns a User object that has a given email address
func (uR *UserRepository) ByEmail(email string) User {

	user := User{}

	uR.Db.ORMConnection.Where("email = ?", email).First(&user)
	return user
}

//ByUsername returns a User object that has a given username
func (uR *UserRepository) ByUsername(username string) User {

	user := User{}

	uR.Db.ORMConnection.Where("username = ?", username).First(&user)
	return user
}

//NewUser creates validates input and creates a new user instance
func (uR *UserRepository) NewUser(username string, email string, password string, passwordVerification string) (User, error) {
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

	if strings.Compare(password, passwordVerification) != 0 {
		return user, errPasswordDoesntMatch
	}

	if email == "" {
		return user, errNoEmail
	}

	duplicateUserByEmail := uR.ByEmail(email)

	if strings.Compare(duplicateUserByEmail.Email, email) == 0 {
		return user, errDuplicateEmail
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	return user, err
}
