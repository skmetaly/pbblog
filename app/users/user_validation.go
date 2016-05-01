package users

import (
	"errors"
	"github.com/skmetaly/pbblog/framework/hash"
	"github.com/skmetaly/pbblog/framework/validation"
	"strings"
)

var (
	errDuplicateUser           = validation.ValidationError(errors.New("This username already exists"))
	errNoUsername              = validation.ValidationError(errors.New("You must supply a username"))
	errNoEmail                 = validation.ValidationError(errors.New("You must supply a email"))
	errNoPassword              = validation.ValidationError(errors.New("You must supply a password"))
	errPasswordDoesntMatch     = validation.ValidationError(errors.New("Passwords doesn't match"))
	errCurrPasswordDoesntMatch = validation.ValidationError(errors.New("Current passwords doesn't match"))
	errPasswordTooShort        = validation.ValidationError(errors.New("Your password is too short"))
	errDuplicateEmail          = validation.ValidationError(errors.New("This email already exists"))
	errBadCredentials          = validation.ValidationError(errors.New("Username and/or password are wrong"))
)

// ValidateUpdate validates the new values for a user
func ValidateUpdate(uR *UserRepository, user User) error {
	if strings.Compare(user.Email, "") == 0 {
		return errNoEmail
	}

	duplicateUserByEmail := uR.ByEmail(user.Email)

	if strings.Compare(duplicateUserByEmail.Email, user.Email) == 0 && duplicateUserByEmail.ID != user.ID {
		return errDuplicateEmail
	}

	if strings.Compare(user.Password, "") == 0 {
		return errNoPassword
	}

	return nil
}

// ValidateLogin validates a user using given credentials
func ValidateLogin(username string, password string) error {

	if username == "" {
		return errNoUsername
	}

	if password == "" {
		return errNoPassword
	}

	return nil
}

// ValidatePasswordChange validates a password change given current password, new password and confirmed password
func ValidatePasswordChange(
	uR *UserRepository,
	user User,
	currentPassword string,
	newPassword string,
	passwordConfirmation string) error {

	// Check if passwords are matching
	passMatch := hash.CompareWithHash([]byte(user.Password), currentPassword)

	if passMatch == false {
		return errCurrPasswordDoesntMatch
	}

	//Current password is true, check if new passwords match"
	if strings.Compare(newPassword, passwordConfirmation) != 0 {
		return errPasswordDoesntMatch
	}

	return nil
}
