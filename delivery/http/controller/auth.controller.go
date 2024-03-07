package controller

import (
	"favorite-book/delivery/http/dto/request"
	"favorite-book/domain"
	"favorite-book/domain/entity"
	"favorite-book/initializer"
	"favorite-book/shared/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	domain domain.Domain
}

func NewUserController(domain domain.Domain) UserController {
	return UserController{
		domain: domain,
	}
}

func (u *UserController) SignUp(ctx *fiber.Ctx) error {
	var user *request.RequestUserDTO

	if err := ctx.BodyParser(&user); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request password")
		return ctx.Status(statusCode).JSON(resp)
	}

	newUser := entity.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
		Role:     user.Role,
	}

	if err := u.domain.UserUsecase.SaveUser(newUser); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to sign up")
		return ctx.Status(statusCode).JSON(resp)
	}

	resp, statusCode := util.ConstructResponseSuccess(fiber.StatusCreated, &user)
	return ctx.Status(statusCode).JSON(resp)
}

func (u *UserController) Login(ctx *fiber.Ctx) error {
	var signIn *request.RequestLoginDTO

	if err := ctx.BodyParser(&signIn); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	user, err := u.domain.UserUsecase.Login(signIn.Username)
	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch user")
		return ctx.Status(statusCode).JSON(resp)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signIn.Password))
	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid username or Password")
		return ctx.Status(statusCode).JSON(resp)
	}

	config, _ := initializer.LoadConfig(".")

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.Username
	claims["exp"] = now.Add(config.JwtExpiresIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(config.JwtSecret))

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "generating JWT Token failed")
		return ctx.Status(statusCode).JSON(resp)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   config.JwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	resp, statusCode := util.ConstructResponseSuccess(fiber.StatusOK, tokenString)
	return ctx.Status(statusCode).JSON(resp)
}

func (u *UserController) Logout(ctx *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	resp, statusCode := util.ConstructResponseSuccess(fiber.StatusOK, "Success logout")
	return ctx.Status(statusCode).JSON(resp)
}
