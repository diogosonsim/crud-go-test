package routes

import (
	"github.com/diogosonsim/crud-go-test/controllers"
	"github.com/diogosonsim/crud-go-test/repository"
	"github.com/diogosonsim/crud-go-test/services"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe public routes.
func Routes(app *fiber.App) {
	userController := controllers.NewUserController(
		services.NewUserService(
			repository.NewUserRepository(),
		),
	)

	app.Get("/api/public/users", userController.GetUsers)

	app.Get("/api/public/users/:id", userController.GetUser)

	app.Post("/api/public/users", userController.CreateUser)

	app.Put("/api/public/users", userController.UpdateUser)

	app.Delete("/api/public/users/:id", userController.DeleteUser)
}
