package tasks

import (
	"net/http"
	"tasquest.com/server/commons"
)

var ErrFailedToFetchTask = commons.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Failed to Fetch the Task",
	Message:  "An unexpected error occurred",
}
