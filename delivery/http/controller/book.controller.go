package controller

import (
	"favorite-book/delivery/http/dto/request"
	"favorite-book/delivery/http/dto/response"
	"favorite-book/domain"
	"favorite-book/shared/util"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	domain domain.Domain
}

func NewBookController(domain domain.Domain) BookController {
	return BookController{
		domain: domain,
	}
}

func (b *BookController) SaveBook(ctx *fiber.Ctx) error {
	var book request.RequestBookDTO

	if err := ctx.BodyParser(&book); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := b.domain.BookUsecase.SaveBook(book.ToBookEntity()); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to save book")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/")
}

func (b *BookController) GetAllBook(ctx *fiber.Ctx) error {
	book, err := b.domain.BookUsecase.GetAllBook()

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch book")
		return ctx.Status(statusCode).JSON(resp)
	}

	responseBooks := make([]response.ResponseBookDTO, len(*book))

	for i, book := range *book {
		responseBooks[i] = response.ResponseBookDTO{
			ID:       i + 1,
			Judul:    book.Judul,
			Penerbit: book.Penerbit,
			Rating:   book.Rating,
		}
	}

	return ctx.Render("index", fiber.Map{
		"Books": responseBooks,
	})
}
