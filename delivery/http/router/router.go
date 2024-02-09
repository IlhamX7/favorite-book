package router

import (
	"favorite-book/delivery/http/controller"
	"favorite-book/domain"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
	bookController := controller.NewBookController(domain)

	app.Post("/", bookController.SaveBook)
	app.Get("/", bookController.GetAllBook)
}
