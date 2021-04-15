package main

import (
	"github.com/diogosonsim/crud-go-test/database"
	"github.com/diogosonsim/crud-go-test/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Routes(app)

	app.Listen(":3000")
}
