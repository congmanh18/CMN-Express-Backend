package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GORM *gorm.DB

// func NewDB(conn Connection) (*gorm.DB, error) {
// 	fmt.Println("start connect to database...")
// 	if err := conn.HasError(); err != nil {
// 		return nil, fmt.Errorf("connection error: %v", err)
// 	}

// 	gormDB, err := gorm.Open(postgres.Open(conn.String()), &gorm.Config{})
// 	if err != nil {
// 		return nil, fmt.Errorf("error connect database: %v", err)
// 	}

// 	fmt.Println("connection database successfuly!")
// 	GORM = gormDB
// 	return gormDB, nil
// }

func NewDB(conn Connection) (*gorm.DB, error) {
	fmt.Println("start connect to database...")

	if err := conn.HasError(); err != nil {
		return nil, fmt.Errorf("connection error: %v", err)
	}

	var gormDB *gorm.DB
	var err error
	for i := 0; i < 10; i++ { // thử lại 10 lần
		gormDB, err = gorm.Open(postgres.Open(conn.String()), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Printf("failed to connect to database (attempt %d/10): %v\n", i+1, err)
		time.Sleep(2 * time.Second) // đợi 2 giây trước khi thử lại
	}

	if err != nil {
		return nil, fmt.Errorf("error connect database: %v", err)
	}

	fmt.Println("connection database successfuly!")
	GORM = gormDB
	return gormDB, nil
}
