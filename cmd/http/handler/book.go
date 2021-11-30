package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/maaarkin/hero-api-golang/internal/service"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bs service.BookService) *BookHandler {
	return &BookHandler{bs}
}

func (b *BookHandler) Route(r chi.Router) {
	r.Get("/hello", b.hello)
	r.Get("/", b.getAll)
}

func (*BookHandler) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

func (h *BookHandler) getAll(w http.ResponseWriter, r *http.Request) {
	if results, err := h.bookService.FindAll(); err != nil {
		addDefaultHeaders(w)
		writeData(
			w,
			map[string]interface{}{
				"Status":      http.StatusInternalServerError,
				"Description": "Internal error, please report to admin",
			},
		)
	} else {
		if results != nil && len(*results) > 0 {
			addDefaultHeaders(w)
			w.WriteHeader(http.StatusOK)
			writeData(w, results)
		} else {
			addDefaultHeaders(w)
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

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
