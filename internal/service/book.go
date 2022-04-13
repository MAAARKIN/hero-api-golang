package service

import (
	"github.com/maaarkin/hero-api-golang/internal/domain/book"
	"github.com/maaarkin/hero-api-golang/internal/repository"
)

type BookService interface {
	Save(book book.Book) (*book.Book, error)
	FindAll() (*[]book.Book, error)
	FindById(id uint64) (*book.Book, error)
	Delete(id uint64) error
	Update(book book.Book) error
}

type bookServiceImpl struct {
	bookStore repository.BookStore
}

func NewBookService(bookStore repository.BookStore) BookService {
	return bookServiceImpl{bookStore}
}

func (bs bookServiceImpl) Save(book book.Book) (*book.Book, error) {
	id, err := bs.bookStore.Create(book)

	if err != nil {
		return nil, err
	}

	book.Id = id
	return &book, nil
}

func (bs bookServiceImpl) FindAll() (*[]book.Book, error) {
	return bs.bookStore.GetAll()
}

func (bs bookServiceImpl) FindById(id uint64) (*book.Book, error) {
	return bs.bookStore.Get(id)
}

func (bs bookServiceImpl) Delete(id uint64) error {
	return bs.bookStore.Delete(id)
}

func (bs bookServiceImpl) Update(book book.Book) error {
	_, err := bs.bookStore.Update(book.Id, book)

	if err != nil {
		return err
	}

	return nil
}
