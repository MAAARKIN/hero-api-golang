package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

type handlers struct {
	Book *BookHandler
}

func initHandlers() handlers {
	return handlers{
		Book: NewBookHandler(),
	}
}

func StartServer() {
	log.Println("Initializing the Http Server at localhost:8080")
	handlers := initHandlers()

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
