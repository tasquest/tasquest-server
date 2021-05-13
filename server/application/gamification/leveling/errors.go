package leveling

import (
	"net/http"

	"tasquest.com/server/commons"
)

var ErrNewLevelCantBeLowerThanLargestExistingLevel = commons.ApplicationError{
	HTTPCode: http.StatusBadRequest,
	Title:    "Failed to save level information",
	Message:  "The new experience level can't be lower than the latest existing experience level.",
}
