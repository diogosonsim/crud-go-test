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

// GetUsers func gets all exists users.
// @Description Get all exists users.
// @Summary get all exists uses
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func (ctrl *UserController) GetUsers(c *fiber.Ctx) error {
	if users, err := ctrl.service.GetUsers(); err != nil {
		c.Status(http.StatusConflict)
		return c.JSON(fiber.Map{
			"message": "Email already exist!",
		})
	} else {
		return c.JSON(users)
	}
}

// GetUser func return user by Id.
// @Description Get user by Id.
// @Summary Retrieves user based on given ID
// @Tags Users
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
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

// CreateUser create user.
// @Description Create user.
// @Summary Create user
// @Tags Users
// @Produce json
// @Param data body models.UserRegister true "User Data"
// @Success 200 {object} models.User
// @Router /users [post]
func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	var userRegister *models.UserRegister

	if err := c.BodyParser(&userRegister); err != nil {
		return err
	}

	if userRegister.Password == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password required",
		})
	}

	if userRegister.ConfirmPassword == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "confirm_password required",
		})
	}

	if userRegister.Email == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "email required",
		})
	}

	if userRegister.Password != userRegister.ConfirmPassword {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
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

// UpdateUser update user.
// @Description Update user.
// @Summary Update user
// @Tags Users
// @Produce json
// @Param data body models.User true "User Data"
// @Success 200 {object} models.User
// @Router /users [put]
func (ctrl *UserController) UpdateUser(c *fiber.Ctx) error {
	var user *models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if user.Email == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "email required",
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

// DeleteUser func delete the user id.
// @Description Delete user by Id.
// @Summary Delete user based on given ID
// @Tags Users
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [delete]
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
