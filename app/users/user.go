package users

//[Todo] move users to framework because view depends on users
import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
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

// User model
type User struct {
	gorm.Model
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	createdAt time.Time
	updatedAt time.Time
}

// UserRepository type
type UserRepository struct {
	Db database.Database
}

// Save persists to database a User type
func (uR *UserRepository) Save(user *User) {
	uR.Db.ORMConnection.Create(user)
}

// Update persists to database changes for a user model
func (uR *UserRepository) Update(user *User) error {

	err := ValidateUpdate(uR, *user)

	if err != nil {
		return err
	}
	uR.Db.ORMConnection.Save(user)
	return nil
}

// ByID returns a User object that has the provided id
func (uR *UserRepository) ByID(usernameID uint) User {

	user := User{}

	uR.Db.ORMConnection.Where("id = ?", usernameID).First(&user)
	return user
}

//ByEmail returns a User object that has a given email address
func (uR *UserRepository) ByEmail(email string) User {

	user := User{}

	uR.Db.ORMConnection.Where("email = ?", email).First(&user)
	return user
}

// ByUsername returns a User object that has a given username
func (uR *UserRepository) ByUsername(username string) User {

	user := User{}

	uR.Db.ORMConnection.Where("username = ?", username).First(&user)
	return user
}

// NewUser creates validates input and creates a new user instance
func (uR *UserRepository) NewUser(
	username string,
	firstName string,
	lastName string,
	email string,
	password string,
	passwordVerification string) (User, error) {
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

	user.FirstName = firstName
	user.LastName = lastName
	return user, err
}

// LoginUser tries to login a user using given credentials
func LoginUser(sess *sessions.Session, uR *UserRepository, username string, password string) (bool, error) {

	err := ValidateLogin(username, password)

	// Check if we have the needed values for login
	if err != nil {
		return false, err
	}

	// Get the username object that has this username
	user := uR.ByUsername(username)

	// Check if the username exists
	if user.ID == 0 {
		return false, errBadCredentials
	}

	// If we have a username, check if passwords are matching
	passMatch := hash.CompareWithHash([]byte(user.Password), password)

	if passMatch == false {
		return false, errBadCredentials
	}

	// Login successful, clear all session variables and add the user details in session
	// Need to thing more of this if it's really necessary
	session.Empty(sess)

	sess.Values["user_id"] = user.ID
	sess.Values["username"] = user.Username

	return true, nil
}
