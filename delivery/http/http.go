package http

import (
	"favorite-book/domain"
	"os"

	"favorite-book/delivery/http/router"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func NewHttpDelivery(domain domain.Domain, engine *html.Engine) *fiber.App {
	config := fiber.Config{
		AppName:           os.Getenv("APP_NAME"),
		EnablePrintRoutes: true,
		JSONEncoder:       sonic.Marshal,
		JSONDecoder:       sonic.Unmarshal,
		Views:             engine,
	}

	if os.Getenv("GO_ENV") == "production" {
		config.Prefork = true
	}

	app := fiber.New(config)
	app.Use(logger.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	router.NewRouter(app, domain)

	return app
}
