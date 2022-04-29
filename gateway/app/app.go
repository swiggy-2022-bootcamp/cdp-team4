package app

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Start() {
	RegisterUserRoutes()
	RegisterOrderRoutes()
	RegisterShippingRoutes()
	RegisterProductAdminRoutes()
	RegisterProductFrontStoreRoutes()
	RegisterCategoryRoutes()
	RegisterPaymentRoutes()
	RegisterCartRouter()
	RegisterRewardRouter()
	RegisterTransactionRouter()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	fmt.Print(PORT)

	err = Router.Run(":8000")
	if err != nil {
		return
	}
}
