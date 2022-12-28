package helper

import (
	"encoding/json"
	"net/http"
)

func BindJson(r *http.Request, destination interface{}) error {
	e := json.NewDecoder(r.Body).Decode(destination)
	if e != nil {
		return ErrDecodeJson
	}
	return nil
}
