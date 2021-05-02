package leveling

import (
	"net/http"

	"tasquest.com/server/commons"
)

var FromExpCannotBeHigherThanToExpError = commons.ApplicationError{
	HTTPCode: http.StatusBadRequest,
	Title:    "Failed to save the adventurer",
	Message:  "It was not possible to save the informed adventurer",
}

var ToExpCannotBeLowerThanFromExpError = commons.ApplicationError{
	HTTPCode: http.StatusBadRequest,
	Title:    "Failed to save the adventurer",
	Message:  "It was not possible to save the informed adventurer",
}

var LevelAlreadyExistsError = commons.ApplicationError{
	HTTPCode: http.StatusConflict,
	Title:    "a",
	Message:  "a",
}

var ErrExperienceOverlaps = commons.ApplicationError{
	HTTPCode: http.StatusConflict,
	Title:    "a",
	Message:  "a",
}
