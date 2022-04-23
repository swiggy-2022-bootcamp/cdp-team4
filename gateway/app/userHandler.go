package app

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type userHandler struct {
}

type ResponseDTO struct {
	Data interface{}
}

type authModel struct {
	UserId string
	Role   int
}

func ValidateToken(authorizationHeader string) authModel {
	authServiceUri := "localhost:8881/api/v1/validate"
	req, err := http.NewRequest("GET", authServiceUri, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", authorizationHeader)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var resDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&resDTO)

	if err != nil {

	}

	model := resDTO.Data.(authModel)

	return authModel{
		UserId: model.UserId,
		Role:   model.Role,
	}
}

func (h userHandler) ValidateAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := ValidateToken(c.Request.Header.Get("Authorization"))
		c.JSON(http.StatusOK, res)
	}
}
