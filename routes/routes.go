package routes

import (
	"github.com/diogosonsim/crud-go-test/controllers"
	"github.com/diogosonsim/crud-go-test/repository"
	"github.com/diogosonsim/crud-go-test/services"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	userController := controllers.NewUserController(
		services.NewUserService(
			repository.NewUserRepository(),
		),
	)

	app.Get("/api/users", userController.GetUsers)

	app.Get("/api/users/:id", userController.GetUser)

	app.Post("/api/users", userController.CreateUser)

	app.Put("/api/users", userController.UpdateUser)

	app.Delete("/api/users/:id", userController.DeleteUser)
}
