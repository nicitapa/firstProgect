package errs

import "errors"

var (
	ErrNotfound                    = errors.New("not found")
	ErrUserNotFound                = errors.New("user not found")
	ErrEmployeesNotfound           = errors.New("employees not found")
	ErrInvalidEmployeesID          = errors.New("invalid employees id")
	ErrInvalidRequestBody          = errors.New("invalid request body")
	ErrInvalidFieldValue           = errors.New("invalid field value")
	ErrInvalidEmployeesName        = errors.New("invalid eployees name, min 4 symbols")
	ErrUsernameAlreadyExists       = errors.New("username already exists")
	ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password")
	ErrInvalidToken                = errors.New("invalid token")
	ErrSomethingWentWrong          = errors.New("something went wrong")
)
