package presenter

import "github.com/maaarkin/hero-api-golang/internal/domain/book"

type BookPersist struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	NumberPages int    `json:"numberPages"`
}

func (b *BookPersist) ToDomain() book.Book {
	return book.Book{
		Title:       b.Title,
		Author:      b.Author,
		NumberPages: b.NumberPages,
	}
}
