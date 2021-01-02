package security

import (
	"net/http"
	"tasquest.com/server/common/errorhandler"
)

var ErrUserNotFound = errorhandler.ApplicationError{
	HTTPCode: http.StatusNotFound,
	Title:    "User not found",
	Message:  "The informed user was not found",
}

var ErrFailedToSaveUser = errorhandler.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Failed to save user",
	Message:  "It was not possible to save the informed user",
}

var ErrUserAlreadyExists = errorhandler.ApplicationError{
	HTTPCode: http.StatusConflict,
	Title:    "User already exists",
	Message:  "There is already an user registered for the informed e-mail",
}

var ErrFailedPasswordGenerate = errorhandler.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Password failure",
	Message:  "The system failed to generate a safe password",
}

var ErrPasswordNotMatch = errorhandler.ApplicationError{
	HTTPCode: http.StatusBadRequest,
	Title:    "Invalid Password",
	Message:  "The Password and the confirmation does not match",
}
