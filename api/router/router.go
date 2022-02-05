package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"toggler.in/api/handler"
	"toggler.in/api/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Flag
	flag := api.Group("/flag")
	flag.Post("/", middleware.Protected(), handler.GetUserFlags)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.LogIn)
	auth.Get("/logout", handler.LogOut)

	// User
	user := api.Group("/user")
	user.Post("/", handler.CreateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)
	user.Get("/status", middleware.Protected(), handler.GetUserStatus)

	// Team
	team := api.Group("/team")
	team.Post("/", middleware.Protected(), handler.CreateTeam)
}