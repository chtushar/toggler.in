package main

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"toggler.in/api/database"
	"toggler.in/api/router"
)

func main() {

	// Create waitgroup
	wg := new(sync.WaitGroup)
	wg.Add(2)

	database.ConnectDB()

	// Client server
	go func ()  {
		app := fiber.New()
		app.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost:3000",
			AllowHeaders:  "Origin, Content-Type, Accept",
			AllowCredentials: true,
		}))
		router.SetupRoutes(app)
		log.Fatal(app.Listen(":9090"))
		wg.Done()
	} ()

	// // Application server
	go func () {
		app2 := fiber.New()
		app2.Get("/", func (c *fiber.Ctx) error {
			return c.JSON(fiber.Map{ "message": "ok" })
		})
		log.Fatal(app2.Listen(":9091"))
		wg.Done()
	} ()

	wg.Wait()
}
