package main

import (
	"github.com/maaarkin/hero-api-golang/cmd/http/handler"
	"github.com/maaarkin/hero-api-golang/internal/container"
)

func main() {

	di := container.Inject()
	//delegate to start our Http Server
	handler.StartServer(di)
}
