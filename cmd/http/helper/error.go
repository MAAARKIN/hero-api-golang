package helper

import (
	"errors"
	"net/http"
)

var (
	errorHandlerMap = make(map[error]int)

	ErrDecodeJson = errors.New("error decoding json")
)

// register how to deal with errors here to HTTP layer
func init() {
	errorHandlerMap[ErrDecodeJson] = http.StatusBadRequest

}

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
