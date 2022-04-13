package container

import (
	"github.com/maaarkin/hero-api-golang/internal/repository"
	"github.com/maaarkin/hero-api-golang/internal/service"
)

// TODO declare services/repositories/components, here
type Container struct {
	Services Services
}

type Services struct {
	BookService service.BookService
}

// Inject represent the starter of our IoC Container, here we will inject
// the necessary structs/functions that we need to build our project.
func Inject() Container {

	//stores
	bookStore := repository.NewBookStoreInMemory()

	//init services
	bs := service.NewBookService(bookStore)

	services := Services{
		BookService: bs,
	}

	return Container{
		Services: services,
	}
}
