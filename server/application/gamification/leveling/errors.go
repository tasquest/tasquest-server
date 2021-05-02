package leveling

import (
	"net/http"

	"tasquest.com/server/commons"
)

var ErrFromExpCannotBeHigherThanToExp = commons.ApplicationError{
	HTTPCode: http.StatusBadRequest,
	Title:    "Failed to save level information",
	Message:  "The 'From' experience can't be higher than the 'To' experience.",
}

var ErrLevelAlreadyExists = commons.ApplicationError{
	HTTPCode: http.StatusConflict,
	Title:    "Failed to save level information",
	Message:  "The informed level already exists",
}

var ErrExperienceOverlaps = commons.ApplicationError{
	HTTPCode: http.StatusConflict,
	Title:    "Failed to save level information",
	Message:  "There is already a level registered between the informed experience range",
}
