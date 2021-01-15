package adventurers

import (
	"net/http"
	"tasquest.com/server"
)

var ErrAdventurerNotFound = server.ApplicationError{
	HTTPCode: http.StatusNotFound,
	Title:    "Adventurer not found",
	Message:  "The informed profile was not found",
}

var ErrFailedToSaveAdventurer = server.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Failed to save the profile",
	Message:  "It was not possible to save the informed profile",
}

var ErrAdventurerAlreadyExists = server.ApplicationError{
	HTTPCode: http.StatusConflict,
	Title:    "Adventurer already exists",
	Message:  "There is already a profile registered for the informed user",
}

var ErrFailedToFetchAdventurer = server.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Failed to Fetch the Adventurer",
	Message:  "An unexpected error occurred",
}
