package presenter

import "github.com/maaarkin/hero-api-golang/internal/domain/book"

type BookPersist struct {
	Title       string `json:"title" validate:"required"`
	Author      string `json:"author" validate:"required"`
	NumberPages int    `json:"numberPages" validate:"required"`
}

func (b *BookPersist) ToDomain() book.Book {
	return book.Book{
		Title:       b.Title,
		Author:      b.Author,
		NumberPages: b.NumberPages,
	}
}
