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
	r.Get("/", b.findAll)
}

func (b *BookHandler) findAll(w http.ResponseWriter, r *http.Request) {
	//TODO find all books
}

func (b *BookHandler) hello(w http.ResponseWriter, r *http.Request) {
	//TODO hello world
}
