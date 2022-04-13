package helper

import (
	"encoding/json"
	"net/http"
)

func addDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}

func writeData(w http.ResponseWriter, data interface{}) {
	if bytes, e := json.Marshal(data); e != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		if _, e := w.Write(bytes); e != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func NoContent(w http.ResponseWriter) {
	addDefaultHeaders(w)
	w.WriteHeader(http.StatusNoContent)
}

func JsonResponse(w http.ResponseWriter, data interface{}, httpStatus int) {
	addDefaultHeaders(w)
	w.WriteHeader(httpStatus)
	writeData(w, data)
}

func HandleError(w http.ResponseWriter, err error) {
	addDefaultHeaders(w)
	httpError := DealWith(err)
	w.WriteHeader(httpError.Status)

	if httpError.Status != http.StatusNoContent {
		writeData(w, httpError)
	}
}
