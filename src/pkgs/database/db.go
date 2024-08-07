package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GORM *gorm.DB

func NewDB(conn Connection) (*gorm.DB, error) {
	fmt.Println("start connect to database...")
	if err := conn.HasError(); err != nil {
		return nil, fmt.Errorf("connection error: %v", err)
	}

	gormDB, err := gorm.Open(postgres.Open(conn.String()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connect database: %v", err)
	}

	fmt.Println("connection database successfuly!")
	GORM = gormDB
	return gormDB, nil
}
