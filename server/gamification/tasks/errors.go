package tasks

import (
	"net/http"
	"tasquest.com/server"
)

var ErrFailedToFetchTask = server.ApplicationError{
	HTTPCode: http.StatusInternalServerError,
	Title:    "Failed to Fetch the Task",
	Message:  "An unexpected error occurred",
}
