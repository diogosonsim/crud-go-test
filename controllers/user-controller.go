package controllers

import (
	"net/http"
	"strconv"

	"github.com/diogosonsim/crud-go-test/models"
	"github.com/diogosonsim/crud-go-test/services"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// Retrieves all users.
// @Description Retrieves all users.
// @Summary Retrieves all users.
// @Tags Private Routes
// @Produce json
// @Success 200 {array} models.User
// @Router /private/users [get]
func (ctrl *UserController) GetUsers(c *fiber.Ctx) error {
	if users, err := ctrl.service.GetUsers(); err != nil {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Invalid request",
		})
	} else {
		return c.JSON(users)
	}
}

// Retrieves user based on given ID
// @Description Retrieves user based on given ID
// @Summary Retrieves user based on given ID
// @Tags Private Routes
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /private/users/{id} [get]
func (ctrl *UserController) GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if user, err := ctrl.service.GetUser(uint(id)); err != nil {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	} else {
		return c.JSON(user)
	}
}

// Create a new user.
// @Description Create a new user.
// @Summary Create a new user.
// @Tags Private Routes
// @Produce json
// @Param data body models.UserRegister true "User Data"
// @Success 200 {object} models.User
// @Router /private/users [post]
func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	var userRegister *models.UserRegister

	if err := c.BodyParser(&userRegister); err != nil {
		return err
	}

	if err := userRegister.IsValidRegister(); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	user := models.User{
		Name:    userRegister.Name,
		Age:     userRegister.Age,
		Address: userRegister.Address,
		Email:   userRegister.Email,
	}
	user.SetPassword(userRegister.Password)

	if value, err := ctrl.service.CreateUser(&user); err != nil {
		c.Status(http.StatusConflict)
		return c.JSON(fiber.Map{
			"message": "Email already exist!",
		})
	} else {
		message := "User created! " + value.Email

		return c.JSON(fiber.Map{
			"message": message,
		})
	}

}

// Update selected user
// @Description Update selected user
// @Summary Update selected user
// @Tags Private Routes
// @Produce json
// @Param data body models.User true "User Data"
// @Success 200 {object} models.User
// @Router /private/users [put]
func (ctrl *UserController) UpdateUser(c *fiber.Ctx) error {
	var user *models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if err := user.IsValid(); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	if value, err := ctrl.service.UpdateUser(user); err != nil {
		c.Status(http.StatusConflict)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	} else {
		return c.JSON(value)
	}

}

// Delete user based on given ID
// @Description Delete user based on given ID
// @Summary Delete user based on given ID
// @Tags Private Routes
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /private/users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if msg, err := ctrl.service.DeleteUser(uint(id)); err != nil {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": msg,
		})
	}
}
