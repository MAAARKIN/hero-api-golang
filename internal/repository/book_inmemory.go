package repository

import "github.com/maaarkin/hero-api-golang/internal/domain/book"

type bookStoreInMemory struct {
}

func NewBookStoreInMemory() BookStore {
	return &bookStoreInMemory{}
}

func (store bookStoreInMemory) GetAll() ([]book.Book, error) {

	return nil, nil
}

func (store bookStoreInMemory) Get(id uint64) (*book.Book, error) {
	return nil, nil
}

func (store bookStoreInMemory) Create(item book.Book) (uint64, error) {
	return 0, nil
}

func (store bookStoreInMemory) Update(id uint64, item book.Book) (*book.Book, error) {
	return nil, nil
}

func (store bookStoreInMemory) Delete(id uint64) error {
	return nil
}
