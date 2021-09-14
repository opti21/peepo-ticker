package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func initDB() {
	var err error
	dsn := os.Getenv("DB_DSN")
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
	fmt.Println("Database connected")

	DBConn.AutoMigrate(&Twitch_cred{})
}
