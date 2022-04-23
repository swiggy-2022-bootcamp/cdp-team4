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

type ValidationDTO struct {
	UserId string      `json:"user_id"`
	Role   int         `json:"role"`
}


func ValidateToken(authorizationHeader string) authModel {
	authServiceUri := "http://localhost:8881/api/v1/validate"
	req, err := http.NewRequest("GET", authServiceUri, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Minute)
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
		log.Fatalf("unable to decode response %v", err)
	}

	var valDTO ValidationDTO
	jsonbytes, _ := json.Marshal(resDTO.Data)
	json.Unmarshal(jsonbytes, &valDTO)

	return authModel{
		UserId: valDTO.UserId,
		Role:   valDTO.Role,
	}
}

func (h userHandler) ValidateAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := ValidateToken(c.Request.Header.Get("Authorization"))

		c.JSON(http.StatusOK, res)
	}
}
