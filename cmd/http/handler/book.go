package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/maaarkin/hero-api-golang/cmd/http/helper"
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
		helper.HandleError(w, err)
	} else {
		if results != nil && len(*results) > 0 {
			helper.JsonResponse(w, results, http.StatusOK)
		} else {
			helper.NoContent(w)
		}
	}
}
