package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/maaarkin/hero-api-golang/internal/container"
	"log"
	"net/http"
)

type handlers struct {
	Book *BookHandler
}

func initHandlers(di container.Container) handlers {
	return handlers{
		Book: NewBookHandler(di.Services.BookService),
	}
}

func StartServer(di container.Container) {
	log.Println("Initializing the Http Server at localhost:8080")
	handlers := initHandlers(di)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/books", handlers.Book.Route)
	})
	http.ListenAndServe(":8080", r)
}
