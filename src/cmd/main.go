package main

import (
	userapis "cmn-express/src/apis/user"
	entity "cmn-express/src/domain/user/entity"
	usecase "cmn-express/src/domain/user/usecase"
	db "cmn-express/src/pkgs/database"

	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var enableMigration = true

func main() {
	slog.Info("service running on port 3000")
	// 1. init fiber instance
	app := fiber.New()

	// 2. connect to database
	var db = connectDB()

	// 3. init route
	var userHander = userapis.UserHandler{
		CreateUserUsecase: usecase.NewCreateUserUsecase(db),
	}
	userapis.SetupUserRoutes(app, userHander)

	app.Listen(":3000")
}

func connectDB() *gorm.DB {
	var conn = db.Connection{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	var gormDB, err = db.NewDB(conn)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		panic(err)
	}

	if enableMigration {
		err := gormDB.AutoMigrate(&entity.User{})
		if err != nil {
			slog.Error("failed to migrate database", "error", err)
			return nil
		}
	}

	return gormDB
}
