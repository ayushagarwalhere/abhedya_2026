// Package configs contain all the initial configrations to start with
package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}
}
