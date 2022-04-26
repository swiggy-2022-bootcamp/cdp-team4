package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/infra"
	"log"
	"os"
)

func Start(testMode bool) {

	healthHandler := HealthHandler{}

	userRepo := infra.NewUserRepository()
	authRepo := infra.NewAuthRepository()

	authService := domain.NewAuthService(userRepo, authRepo)
	authHandler := AuthHandler{
		AuthService: authService,
	}
	RegisterHealthStatusRoute(healthHandler)
	RegisterAuthHandlerRoute(authHandler)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		return
	}

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	PORT := os.Getenv("PORT")

	if !testMode {
		err = Router.Run(fmt.Sprintf(":%s", PORT))
		if err != nil {
			return
		}
	}
}
