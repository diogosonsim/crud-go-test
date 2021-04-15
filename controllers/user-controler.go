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

func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	var user *models.User
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if data["password"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password required",
		})
	}

	if data["confirm_password"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "confirm_password required",
		})
	}

	if user.Email == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "email required",
		})
	}

	if data["password"] != data["confirm_password"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	user.SetPassword(data["password"])

	if value, err := ctrl.service.CreateUser(user); err != nil {
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
