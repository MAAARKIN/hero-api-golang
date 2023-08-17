package helper

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func BindJson(r *http.Request, destination interface{}) error {
	e := json.NewDecoder(r.Body).Decode(destination)
	if e != nil {
		return ErrDecodeJson
	}
	return validateJsonFields(destination)
}

func validateJsonFields(input interface{}) error {
	validator := validator.New()
	if err := validator.Struct(input); err != nil {
		return err
	}
	return nil
}
