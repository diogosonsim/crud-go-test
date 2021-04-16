package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/diogosonsim/crud-go-test/database"
	"github.com/diogosonsim/crud-go-test/models"
	"github.com/diogosonsim/crud-go-test/services"
	"github.com/diogosonsim/crud-go-test/utils"
	"github.com/gofiber/fiber/v2"
)

func NewAuthController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// Register User.
// @Description Public route to create a new user.
// @Summary Public route to create a new user
// @Tags Public Routes
// @Produce json
// @Param data body models.UserRegister true "User Data"
// @Success 200 {object} models.User
// @Router /public/register [post]
func (ctrl *UserController) Register(c *fiber.Ctx) error {
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
		Email:   userRegister.Email,
		Address: userRegister.Address,
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

// Login.
// @Description Public route to authenticate a user.
// @Summary Public route to authenticate a user
// @Tags Public Routes
// @Produce json
// @Param data body models.LoginRequest true "User Data"
// @Success 200 {object} models.User
// @Router /public/login [post]
func (ctrl *UserController) Login(c *fiber.Ctx) error {
	loginRequest := models.LoginRequest{}

	if err := c.BodyParser(&loginRequest); err != nil {
		return err
	}

	if err := loginRequest.IsValidLogin(); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	user := models.User{}

	if values, err := ctrl.service.GetUserByEmail(loginRequest.Email); err != nil {
		c.Status(http.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	} else {
		user.Password = values.Password
	}

	if err := user.ComparePassword(loginRequest.Password); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// Get logged user.
// @Description Retrieves user based on access token
// @Summary Retrieves user based on access token
// @Tags Private Routes
// @Produce json
// @Success 200 {object} models.User
// @Router /private/logged [get]
func Logged(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJwt(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

// Logout user.
// @Description Remove access token from cookie
// @Summary Remove access token from cookie
// @Tags Private Routes
// @Produce json
// @Success 200 {object} models.User
// @Router /private/logout [post]
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// Healthcheck.
// @Description Healthcheck
// @Summary Healthcheck
// @Tags Public Routes
// @Produce json
// @Success 200 string message
// @Router /public/healthCheck [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "The application is running.",
	})
}
