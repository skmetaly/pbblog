package users

import (
	"github.com/jinzhu/gorm"
	//"github.com/satori/go.uuid"
	//"errors"
	//"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/sessions"
	"github.com/skmetaly/pbblog/framework/database"
	"github.com/skmetaly/pbblog/framework/hash"
	"github.com/skmetaly/pbblog/framework/session"
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

//ByID returns a User object that has the provided id
func (uR *UserRepository) ByID(usernameId int) User {

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

//ValidateLogin validates a user using given credentials
func ValidateLogin(username string, password string) error {

	if username == "" {
		return errNoUsername
	}

	if password == "" {
		return errNoPassword
	}

	return nil
}

//LoginUser tries to login a user using given credentials
func LoginUser(sess *sessions.Session, uR *UserRepository, username string, password string) (bool, error) {

	err := ValidateLogin(username, password)

	//Check if we have the needed values for login
	if err != nil {
		return false, err
	}

	//Get the username object that has this username
	user := uR.ByUsername(username)

	//Check if the username exists
	if user.ID == 0 {
		return false, errBadCredentials
	}

	//If we have a username, check if passwords are matching
	passMatch := hash.CompareWithHash([]byte(user.Password), password)

	if passMatch == false {
		return false, errBadCredentials
	}

	//Login successful, clear all session variables and add the user details in session
	//Need to thing more of this if it's really necessary
	session.Empty(sess)

	sess.Values["user_id"] = user.ID
	sess.Values["username"] = user.Username

	return true, nil
}
