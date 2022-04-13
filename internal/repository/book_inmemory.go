package repository

import (
	"errors"
	"sync"

	"github.com/maaarkin/hero-api-golang/internal/domain/book"
)

var (
	keyInstance = uint64(3)

	books = map[uint64]book.Book{
		1: {Id: 1, Title: "Title 1", Author: "MarkMark", NumberPages: 101},
		2: {Id: 2, Title: "Title 2", Author: "MarkMark 2", NumberPages: 203},
	}
)

type bookStoreInMemory struct {
	mu sync.Mutex
}

func NewBookStoreInMemory() BookStore {
	return &bookStoreInMemory{}
}

func (store bookStoreInMemory) GetAll() (*[]book.Book, error) {
	m := make([]book.Book, 0)

	for _, value := range books {
		m = append(m, value)
	}

	return &m, nil
}

func (store bookStoreInMemory) Get(id uint64) (*book.Book, error) {
	book, has := books[id]
	if !has {
		return nil, errors.New("No book in database")
	}
	return &book, nil
}

func (store bookStoreInMemory) Create(item book.Book) (uint64, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	keyInstance = keyInstance + 1
	books[keyInstance] = item

	return keyInstance, nil
}

func (store bookStoreInMemory) Update(id uint64, item book.Book) (*book.Book, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	books[keyInstance] = item

	return &item, nil
}

func (store bookStoreInMemory) Delete(id uint64) error {
	delete(books, id)
	return nil
}
