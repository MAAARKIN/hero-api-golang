package helper

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var (
	errorHandlerMap = make(map[error]int)

	ErrDecodeJson     = errors.New("error decoding json")
	ErrJsonValidation = errors.New("error validating json")
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
	if ok, httpError := isValidationError(err); ok {
		return *httpError
	} else if errCode, ok := errorHandlerMap[err]; ok {
		return HttpError{Status: errCode, Description: err.Error()}
	} else {
		return HttpError{
			Status:      http.StatusInternalServerError,
			Description: "Internal error, please report to admin",
		}
	}
}

func isValidationError(err error) (bool, *HttpError) {
	v := &validator.ValidationErrors{}
	if errors.As(err, v) {
		validationErrors := HttpError{Status: http.StatusBadRequest, Description: ErrJsonValidation.Error()}
		for _, err := range err.(validator.ValidationErrors) {
			message := generateErrorMessage(err)
			validationErrors.Messages = append(validationErrors.Messages, message.Error())
		}
		return true, &validationErrors
	}
	return false, nil
}

func generateErrorMessage(err validator.FieldError) error {
	return fmt.Errorf("error: field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag())
}
