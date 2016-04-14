package validation

type ValidationError error

func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)

	return ok
}
