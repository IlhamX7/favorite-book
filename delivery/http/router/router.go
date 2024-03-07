package router

import (
	"favorite-book/delivery/http/controller"
	"favorite-book/domain"
	"favorite-book/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
	bookController := controller.NewBookController(domain)

	app.Get("", bookController.Welcome)
	app.Post("/create", middleware.DeserializeUser, bookController.SaveBook)
	app.Get("/get-all", middleware.DeserializeUser, bookController.GetAllBook)
	app.Get("/get-one/:bookId", middleware.DeserializeUser, bookController.GetOneBook)
	app.Put("/update/:bookId", middleware.DeserializeUser, bookController.UpdateBook)
	app.Delete("/delete/:bookId", middleware.DeserializeUser, bookController.DeleteBook)

	userController := controller.NewUserController(domain)

	app.Post("/sign-up", userController.SignUp)
	app.Post("/login", userController.Login)
	app.Post("/logout", middleware.DeserializeUser, userController.Logout)
}
