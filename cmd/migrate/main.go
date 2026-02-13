package main

import (
	"log"

	"abhedya_2026/internal/configs"
	"abhedya_2026/internal/models"
)

func init() {
	configs.LoadEnv()
	configs.ConnectDB()
}

func main() {
	err := configs.DB.AutoMigrate(
		&models.User{},
		&models.Question{},
		&models.Submission{},
	)
	if err != nil {
		log.Fatal("Migrations failed", err)
	}
	log.Println("Migrations completed successfully")
}
