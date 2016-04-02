package users

import (
	"errors"
	"github.com/skmetaly/pbblog/framework/validation"
)

var (
	errDuplicateUser       = validation.ValidationError(errors.New("This username already exists"))
	errNoUsername          = validation.ValidationError(errors.New("You must supply a username"))
	errNoEmail             = validation.ValidationError(errors.New("You must supply a email"))
	errNoPassword          = validation.ValidationError(errors.New("You must supply a password"))
	errPasswordDoesntMatch = validation.ValidationError(errors.New("Passwords doesn't match"))
	errPasswordTooShort    = validation.ValidationError(errors.New("Your password is too short"))
	errDuplicateEmail      = validation.ValidationError(errors.New("This email already exists"))
)
