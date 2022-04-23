package app

import "github.com/joho/godotenv"

import (
	"fmt"
	"log"
	"os"
)

func Start() {
	RegisterUserRoutes()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	PORT := os.Getenv("PORT")

	err = Router.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		return
	}
}
