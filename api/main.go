package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
    return c.SendString(msg) // => ✋ register
	})

	log.Fatal(app.Listen(":9090"))
}