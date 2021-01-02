package profiles

import (
	"net/http"
	"tasquest.com/server/common/errorhandler"
)

var ErrProfileNotFound = errorhandler.ApplicationError{
	HTTPCode: http.StatusNotFound,
	Title:    "Profile not found",
	Message:  "The informed profile was not found",
}

var ErrFailedToSaveProfile = errorhandler.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Failed to save the profile",
	Message:  "It was not possible to save the informed profile",
}

var ErrProfileAlreadyExists = errorhandler.ApplicationError{
	HTTPCode: http.StatusConflict,
	Title:    "Profile already exists",
	Message:  "There is already a profile registered for the informed user",
}

var ErrFailedToFetchProfile = errorhandler.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Failed to Fetch Profile",
	Message:  "An unexpected error occurred",
}
