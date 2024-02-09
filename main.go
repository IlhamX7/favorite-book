package main

import (
	"favorite-book/delivery/http"
	"fmt"

	"os"

	"favorite-book/domain"

	"github.com/gofiber/template/html/v2"
)

func main() {
	domain := domain.ConstructDomain()
	engine := html.New("./views", ".html")
	app := http.NewHttpDelivery(domain, engine)
	port := os.Getenv("PORT_APP")
	if port == "" {
		port = "8080"
	}
	app.Listen(fmt.Sprintf(":%s", port))
}
