package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	_ "github.com/diogosonsim/crud-go-test/docs"
)

func SwaggerRoute(app *fiber.App) {
	// Routes for GET Swagger-ui:
	app.Get("/swagger/*", swagger.Handler)
}
