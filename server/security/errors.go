package security

import (
	"net/http"
	"tasquest.com/server"
)

var ErrUserNotFound = server.ApplicationError{
	HTTPCode: http.StatusNotFound,
	Title:    "User not found",
	Message:  "The informed user was not found",
}

var ErrFailedToSaveUser = server.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Failed to save user",
	Message:  "It was not possible to save the informed user",
}

var ErrUserAlreadyExists = server.ApplicationError{
	HTTPCode: http.StatusConflict,
	Title:    "User already exists",
	Message:  "There is already an user registered for the informed e-mail",
}

var ErrFailedPasswordGenerate = server.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Password failure",
	Message:  "The system failed to generate a safe password",
}

var ErrPasswordNotMatch = server.ApplicationError{
	HTTPCode: http.StatusBadRequest,
	Title:    "Invalid Password",
	Message:  "The Password and the confirmation does not match",
}
