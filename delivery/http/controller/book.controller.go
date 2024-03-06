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

	// return ctx.Redirect("/")
	resp, statusCode := util.ConstructResponseSuccess(fiber.StatusCreated, &book)
	return ctx.Status(statusCode).JSON(resp)
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
			ID:       book.ID,
			Judul:    book.Judul,
			Penerbit: book.Penerbit,
			Rating:   book.Rating,
		}
	}

	// return ctx.Render("index", fiber.Map{
	// 	"Books": responseBooks,
	// })

	resp, statusCode := util.ConstructResponseSuccess(fiber.StatusOK, responseBooks)
	return ctx.Status(statusCode).JSON(resp)
}

func (b *BookController) GetOneBook(ctx *fiber.Ctx) error {
	bookId := ctx.Params("bookId")
	book, err := b.domain.BookUsecase.GetOneBook(bookId)

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch book")
		return ctx.Status(statusCode).JSON(resp)
	}

	resp, statusCode := util.ConstructResponseSuccess(fiber.StatusOK, book)
	return ctx.Status(statusCode).JSON(resp)
}

func (b *BookController) UpdateBook(ctx *fiber.Ctx) error {
	bookId := ctx.Params("bookId")

	var payload *request.RequestBookDTO

	if err := ctx.BodyParser(&payload); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, err.Error())
		return ctx.Status(statusCode).JSON(resp)
	}

	updates := make(map[string]interface{})
	if payload.Judul != "" {
		updates["judul"] = payload.Judul
	}
	if payload.Penerbit != "" {
		updates["penerbit"] = payload.Penerbit
	}
	if payload.Rating != 0 {
		updates["rating"] = payload.Rating
	}

	book, err := b.domain.BookUsecase.UpdateBook(bookId, updates)

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch book")
		return ctx.Status(statusCode).JSON(resp)
	}

	resp, statusCode := util.ConstructResponseSuccess(fiber.StatusOK, book)
	return ctx.Status(statusCode).JSON(resp)
}

func (b *BookController) DeleteBook(ctx *fiber.Ctx) error {
	bookId := ctx.Params("bookId")
	_, err := b.domain.BookUsecase.DeleteBook(bookId)

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch book")
		return ctx.Status(statusCode).JSON(resp)
	}

	resp, statusCode := util.ConstructResponseSuccess(fiber.StatusOK, "Success delete book")
	return ctx.Status(statusCode).JSON(resp)
}
