package userController

import (
	"net/http"
	user "swyw-users/src/models/users"
	usersServices "swyw-users/src/services/users"
	messageError "swyw-users/src/utils/Error"
	passwordHashing "swyw-users/src/utils/crypto"
	logger "swyw-users/src/utils/logs"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func CreateUser(c *fiber.Ctx) error {

	var req user.UserRegister

	if err := c.BodyParser(&req); err != nil {
		logger.Log.Warn("Invalid request body", zap.Error(err))
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
		logger.Log.Error("Error searching for user",
			zap.Error(err))
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"error": messageError.ErrSearchingForUser,
			})
	}

	if u != nil {
		logger.Log.Warn("Error saving user, user exits", zap.String("userEmail", req.Email))
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	id, saveError := usersServices.SaveUser(&req)

	if saveError != nil {
		logger.Log.Error("Error saving user",
			zap.Error(err))
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": messageError.ErrInsertUser,
		})
	}

	logger.Log.Info("User created", zap.Int("userId", id), zap.String("email", req.Email))
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
		logger.Log.Error("Error searching for user",
			zap.Error(err))
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"error": messageError.ErrSearchingForUser,
			})
	}

	if u == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	if !passwordHashing.VerifyPassword(req.Pass, u.Pass) {
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

	logger.Log.Info("User login", zap.String("userEmail", userResp.Email))
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"user": userResp,
	})

}

func GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is required",
		})
	}

	u, err := usersServices.FindUser(email)

	if err != nil {
		logger.Log.Error("Error searching for user",
			zap.Error(err))
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"error": messageError.ErrSearchingForUser,
			})
	}

	if u == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"msg": "user not found",
		})
	}

	userResp := user.UserResponse{
		Id:        u.Id,
		Name:      u.Name,
		Email:     u.Email,
		Create_at: u.Create_at,
	}

	logger.Log.Info("User get by email", zap.String("userEmail", userResp.Email))
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"user": userResp,
	})
}
