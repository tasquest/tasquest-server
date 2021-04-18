package adventurers

import (
	"net/http"
	"tasquest.com/server/commons"
)

var ErrFailedToSaveAdventurer = commons.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Failed to save the adventurer",
	Message:  "It was not possible to save the informed adventurer",
}

var ErrAdventurerAlreadyExists = commons.ApplicationError{
	HTTPCode: http.StatusConflict,
	Title:    "Adventurer already exists",
	Message:  "There is already an adventurer registered for the informed user",
}

var ErrFailedToFetchAdventurer = commons.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Failed to Fetch the Adventurer",
	Message:  "An unexpected error occurred",
}
