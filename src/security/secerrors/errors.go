package secerrors

import (
	"net/http"
	"tasquest-server/src/common/errorhandler"
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
