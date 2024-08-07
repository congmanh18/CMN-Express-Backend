package userapis

import (
	entity "cmn-express/src/domain/user/entity"
	usecase "cmn-express/src/domain/user/usecase"
	"cmn-express/src/pkgs/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	CreateUserUsecase usecase.CreateUserUsecase
	// LoginUserUsecase  usecase.LoginUserUsecase
	// UpdateUserUsecase usecase.UpdateUserUsecase
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
			Phone:      &createUserReq.Phone,
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
