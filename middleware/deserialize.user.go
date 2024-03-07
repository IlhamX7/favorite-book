package middleware

import (
	"favorite-book/domain/entity"
	"favorite-book/infrastructure"
	"favorite-book/initializer"
	"favorite-book/shared/util"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DeserializeUser(ctx *fiber.Ctx) error {
	var tokenString string
	authorization := ctx.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if ctx.Cookies("token") != "" {
		tokenString = ctx.Cookies("token")
	}

	if tokenString == "" {
		resp, statusCode := util.ConstructResponseError(fiber.StatusUnauthorized, "You are not logged in")
		return ctx.Status(statusCode).JSON(resp)
	}

	config, _ := initializer.LoadConfig(".")

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(config.JwtSecret), nil
	})

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusUnauthorized, fmt.Sprintf("invalidate token: %v", err))
		return ctx.Status(statusCode).JSON(resp)
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		resp, statusCode := util.ConstructResponseError(fiber.StatusUnauthorized, "invalid token claim")
		return ctx.Status(statusCode).JSON(resp)
	}

	var user entity.User
	infrastructure.NewDatabaseConnection().First(&user, "username = ?", fmt.Sprint(claims["sub"]))

	if user.Username != claims["sub"] {
		resp, statusCode := util.ConstructResponseError(fiber.StatusForbidden, "the user belonging to this token no logger exists")
		return ctx.Status(statusCode).JSON(resp)
	}

	ctx.Locals("user", entity.FilterUserRecord(&user))

	return ctx.Next()
}
