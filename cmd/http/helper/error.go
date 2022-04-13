package helper

import "net/http"

type HttpError struct {
	Status      int      `json:"-"`
	Description string   `json:"description,omitempty"`
	Messages    []string `json:"messages,omitempty"`
}

func DealWith(err error) HttpError {

	return HttpError{
		Status:      http.StatusInternalServerError,
		Description: "Internal error, please report to unico Team",
	}
}
