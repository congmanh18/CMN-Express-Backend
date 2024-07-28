package user_api

import (
	entity "cmn-express/domain/user/entity"
	"cmn-express/domain/user/usecase"
	"cmn-express/pkgs/jwtauth"
	"cmn-express/pkgs/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	CreateUserUsecase usecase.CreateUserUsecase
	LoginUserUsecase  usecase.LoginUserUsecase
	UpdateUserUsecase usecase.UpdateUserUsecase
}

func (u UserHandler) HandleCreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Bắt request người dùng gửi lên
		var createUserReq = new(CreateUserReq)
		if err := c.BodyParser(&createUserReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		// Hash the user's password
		hashedPassword, err := utils.HashPassword(createUserReq.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		// init user entity
		var userEntity = entity.User{
			ID:         createUserReq.ID,
			First_name: &createUserReq.First_name,
			Last_name:  &createUserReq.Last_name,
			Email:      &createUserReq.Email,
			Password:   &hashedPassword,
			Role:       &createUserReq.Role,
		}

		// execute usecase
		var useCaseErr = u.CreateUserUsecase.ExecuteCreateUser(c.Context(), userEntity)
		if useCaseErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": useCaseErr.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "OK",
		})
	}
}

func (u UserHandler) HandleLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var loginUserReq = new(LoginUserReq)
		if err := c.BodyParser(&loginUserReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		// Fetch user entity from database by email
		var userEntity entity.User
		err := u.LoginUserUsecase.ExecuteLoginUser(c.Context(), loginUserReq.Email, userEntity)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": "Invalid credentials",
			})
		}

		// Verify the user's password
		err = utils.VerifyPassword(*userEntity.Password, loginUserReq.Password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"error": "Invalid credentials",
			})
		}

		// Generate JWT tokens
		accessToken, refreshToken, err := jwtauth.GenerateAllToken(userEntity.ID, *userEntity.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		// *Edit later*
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message":       "Login successfully",
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}

func (u UserHandler) HandleUpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var createUserReq = new(CreateUserReq)
		if err := c.BodyParser(&createUserReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "OK",
		})
	}
}

func CheckUser(c *fiber.Ctx) error {
	// Example handler using utils
	var u entity.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	userType := utils.CheckUserType(u)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user_type": userType})
}

func MatchUser(c *fiber.Ctx) error {
	var u entity.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	id := c.Params("id")
	// Convert id to uint and handle errors
	match, err := utils.MatchUserToID(u, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"match": match})
}
