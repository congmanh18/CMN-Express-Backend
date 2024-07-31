package user_api

import "github.com/gofiber/fiber/v2"

func SetupUserRoutes(app *fiber.App, userHandler UserHandler) {
	app.Post("/user/register", userHandler.HandleCreateUser())
	app.Post("/user/login", userHandler.HandleLogin())
	// app.Put()
}
