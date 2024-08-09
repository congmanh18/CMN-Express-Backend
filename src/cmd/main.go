package main

import (
	userapis "cmn-express/src/apis/user"
	usecase "cmn-express/src/internal/domain/user/usecase"
	entity "cmn-express/src/internal/domain/user/entity"
	db "cmn-express/src/pkgs/database"
	"log"

	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
	// Load các biến môi trường từ file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var conn = db.Connection{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	var gormDB *gorm.DB
	gormDB, err = db.NewDB(conn)
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
