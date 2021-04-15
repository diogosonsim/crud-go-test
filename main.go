package main

import (
	"github.com/diogosonsim/crud-go-test/database"
	"github.com/diogosonsim/crud-go-test/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Go test Swagger API
// @version 1.0
// @description Swagger API for Golang Test.
// @termsOfService http://swagger.io/terms/
// @BasePath /api/public
func main() {
	database.Connect()

	// Define a new Fiber app with config.
	app := fiber.New()

	// Allow Cors.
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Routes.
	routes.Routes(app)
	routes.SwaggerRoute(app)

	app.Listen(":3000")
}
