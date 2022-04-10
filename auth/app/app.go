package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Start() {

	healthHandler := HealthHandler{}

	RegisterHealthStatusRoute(healthHandler)

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
