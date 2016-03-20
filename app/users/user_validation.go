package users

import (
	"errors"
	"github.com/skmetaly/pbblog/framework/validation"
)

var (
	errNoUsername       = validation.ValidationError(errors.New("You must supply a username"))
	errNoEmail          = validation.ValidationError(errors.New("You must supply a email"))
	errNoPassword       = validation.ValidationError(errors.New("You must supply a password"))
	errPasswordTooShort = validation.ValidationError(errors.New("Your password is too short"))
)
