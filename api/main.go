package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"toggler.in/api/database"
	"toggler.in/api/router"
)

func main() {
	app := fiber.New()
	database.ConnectDB()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":9090"))
}
