package security

import (
	"net/http"
	"tasquest.com/server/commons"
)

var ErrUserNotFound = commons.ApplicationError{
	HTTPCode: http.StatusNotFound,
	Title:    "User not found",
	Message:  "The informed user was not found",
}

var ErrFailedToSaveUser = commons.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Failed to save user",
	Message:  "It was not possible to save the informed user",
}

var ErrUserAlreadyExists = commons.ApplicationError{
	HTTPCode: http.StatusConflict,
	Title:    "User already exists",
	Message:  "There is already an user registered for the informed e-mail",
}

var ErrFailedPasswordGenerate = commons.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Password failure",
	Message:  "The system failed to generate a safe password",
}

var ErrPasswordNotMatch = commons.ApplicationError{
	HTTPCode: http.StatusBadRequest,
	Title:    "Invalid Password",
	Message:  "The Password and the confirmation does not match",
}
