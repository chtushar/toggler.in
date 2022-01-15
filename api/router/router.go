package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"toggler.in/api/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Flag
	flag := api.Group("/flag")
	flag.Post("/", handler.GetUserFlags)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Post("/", handler.CreateUser)
}