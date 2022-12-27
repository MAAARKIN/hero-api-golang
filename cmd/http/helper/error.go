package helper

import (
	"net/http"
)

var (
	errorHandlerMap = make(map[error]int)
)

type HttpError struct {
	Status      int      `json:"-"`
	Description string   `json:"description,omitempty"`
	Messages    []string `json:"messages,omitempty"`
}

func DealWith(err error) HttpError {
	if errCode, ok := errorHandlerMap[err]; ok {
		return HttpError{Status: errCode, Description: err.Error()}
	} else {
		return HttpError{
			Status:      http.StatusInternalServerError,
			Description: "Internal error, please report to admin",
		}
	}
}
