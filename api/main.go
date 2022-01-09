package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"toggler.in/api/database"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	app.Get("/", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	// signup

	// login
	// logout

	log.Fatal(app.Listen(":9090"))
}
