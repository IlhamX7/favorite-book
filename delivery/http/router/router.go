package router

import (
	"favorite-book/delivery/http/controller"
	"favorite-book/domain"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
	bookController := controller.NewBookController(domain)

	app.Post("/create", bookController.SaveBook)
	app.Get("/get-all", bookController.GetAllBook)
	app.Get("/get-one/:bookId", bookController.GetOneBook)
	app.Put("/update/:bookId", bookController.UpdateBook)
	app.Delete("/delete/:bookId", bookController.DeleteBook)
}
