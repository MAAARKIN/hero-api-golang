package handler

import (
	"github.com/go-chi/chi"
	"net/http"
)

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (b *BookHandler) Route(r chi.Router) {
	r.Get("/hello", b.hello)
}

func (*BookHandler) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
