package handler

import (
	"github.com/go-chi/chi"
	"github.com/maaarkin/hero-api-golang/internal/service"
	"net/http"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bs service.BookService) *BookHandler {
	return &BookHandler{bs}
}

func (b *BookHandler) Route(r chi.Router) {
	r.Get("/hello", b.hello)
}

func (*BookHandler) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
