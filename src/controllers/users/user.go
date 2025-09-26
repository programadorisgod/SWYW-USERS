package userController

import (
	"log"
	"net/http"
	user "swyw-users/src/models/users"
	usersServices "swyw-users/src/services/users"
	passwordHashing "swyw-users/src/utils/crypto"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {

	var req user.UserRegister

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	//TODO: create a middleware to validation
	if req.Email == "" || req.Pass == "" || req.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Email, password and name are required",
		})
	}

	u, err := usersServices.FindUser(req.Email)

	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	if u != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	id, error := usersServices.SaveUser(&req)

	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create user",
			"e":     err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func AuthenticateUser(c *fiber.Ctx) error {
	var req user.UserLogin
	//TODO: create a middleware to validation

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	u, err := usersServices.FindUser(req.Email)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "error searching user",
		})
	}

	if u == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	if !passwordHashing.VerifyPassword(req.Pass, u.Pass) {
		log.Printf("Intento de login fallido para el email: %s. Email enviado: %s, Email esperado: %s, Pass enviada: %s, Pass esperada: %s", req.Email, req.Email, u.Email, req.Pass, u.Pass)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	userResp := user.UserResponse{
		Id:        u.Id,
		Name:      u.Name,
		Email:     u.Email,
		Create_at: u.Create_at,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"user": userResp,
	})

}
