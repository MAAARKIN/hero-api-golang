package repository

import "github.com/maaarkin/hero-api-golang/internal/domain/book"

type BookStore interface {
	GetAll() (*[]book.Book, error)
	Get(id uint64) (*book.Book, error)
	Create(item book.Book) (uint64, error)
	Update(id uint64, item book.Book) (*book.Book, error)
	Delete(id uint64) error
}
