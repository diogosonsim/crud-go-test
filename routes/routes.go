package routes

import (
	"github.com/diogosonsim/crud-go-test/controllers"
	"github.com/diogosonsim/crud-go-test/middlewares"
	"github.com/diogosonsim/crud-go-test/repository"
	"github.com/diogosonsim/crud-go-test/services"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe public routes.
func Routes(app *fiber.App) {
	// User controller routes
	userController := controllers.NewUserController(
		services.NewUserService(
			repository.NewUserRepository(),
		),
	)

	// Authentication routes
	authController := controllers.NewAuthController(
		services.NewUserService(
			repository.NewUserRepository(),
		),
	)

	//Public routes
	app.Post("/api/public/register", authController.Register)
	app.Post("/api/public/login", authController.Login)
	app.Get("/api/public/healthcheck", controllers.HealthCheck)

	//Private routes
	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/private/logged", controllers.Logged)
	app.Post("/api/private/logout", controllers.Logout)

	app.Get("/api/private/users", userController.GetUsers)

	app.Get("/api/private/users/:id", userController.GetUser)

	app.Post("/api/private/users", userController.CreateUser)

	app.Put("/api/private/users", userController.UpdateUser)

	app.Delete("/api/private/users/:id", userController.DeleteUser)
}
