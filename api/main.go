package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"toggler.in/api/database"
	"toggler.in/api/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:3000",
    AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
}))
	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":9090"))
}
