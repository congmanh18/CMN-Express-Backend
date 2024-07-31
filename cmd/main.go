package main

import (
	user_api "cmn-express/apis/user"
	entity "cmn-express/domain/user/entity"
	"cmn-express/domain/user/usecase"
	db "cmn-express/pkgs/database"
	"log/slog"

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
	var userHander = user_api.UserHandler{
		CreateUserUsecase: usecase.NewCreateUserUsecase(db),
	}
	user_api.SetupUserRoutes(app, userHander)

	app.Listen(":3000")
}

func connectDB() *gorm.DB {
	var conn = db.Connection{
		Host:     "localhost",
		User:     "postgres",
		Password: "231002",
		DBName:   "postgres",
		Port:     "5432",
	}

	var gormDB, err = db.NewDB(conn)
	if err != nil {
		panic(err)
	}

	if enableMigration {
		//gormDB.Exec("CREATE DATABASE gorm_db")
		// err := gormDB.Debug().AutoMigrate(&entity.User{})
		err := gormDB.AutoMigrate(&entity.User{})
		if err != nil {
			return nil
		}
	}

	return gormDB
}
