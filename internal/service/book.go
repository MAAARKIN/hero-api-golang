package service

import (
	"github.com/maaarkin/hero-api-golang/internal/domain/book"
)

type BookService interface {
	Save(book book.Book) (*book.Book, error)
	FindAll() error
	FindById(id uint64) (*book.Book, error)
	Delete(id uint64) error
	Update(book book.Book) error
}

type bookServiceImpl struct {
}

func NewBookService() BookService {
	return bookServiceImpl{}
}

func (bookServiceImpl) Save(book book.Book) (*book.Book, error) {
	return nil, nil
}

func (bookServiceImpl) FindAll() error {
	return nil
}

func (bookServiceImpl) FindById(id uint64) (*book.Book, error) {
	return nil, nil
}

func (bookServiceImpl) Delete(id uint64) error {
	return nil
}

func (bookServiceImpl) Update(book book.Book) error {
	return nil
}
