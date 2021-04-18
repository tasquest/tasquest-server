package commons

import (
	"emperror.dev/errors"
	"net/http"
)

type ApplicationError struct {
	HTTPCode int    `json:"code"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}

func (ar ApplicationError) Error() string {
	return ar.Message
}

func ParseError(err error) (ApplicationError, error) {
	var appError ApplicationError
	var unwrap = errors.Cause(err)

	if errors.As(unwrap, &appError) {
		return appError, nil
	}

	return ApplicationError{
		HTTPCode: http.StatusInternalServerError,
		Title:    "Unknown Error",
		Message:  "An unexpected error occurred",
	}, err
}
