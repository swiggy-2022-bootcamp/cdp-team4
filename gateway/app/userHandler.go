package app

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/gateway/utils/errs"
	"log"
	"net/http"
	"time"
)

type userHandler struct {
}

type AuthModel struct {
	UserId string
	Role   int
}

type ValidationDTO struct {
	UserId string `json:"user_id"`
	Role   int    `json:"role"`
}

func ValidateToken(authorizationHeader string) (*AuthModel, *errs.AppError) {
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
		return nil, errs.NewAuthenticationError(err.Error())
	}

	var resDTO ResponseDTO
	err = json.NewDecoder(res.Body).Decode(&resDTO)

	if err != nil {
		//log.Fatalf("unable to decode response %v", err)
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if resDTO.Status == http.StatusUnauthorized {
		return nil, errs.NewAuthenticationError(resDTO.Message)
	}

	var valDTO ValidationDTO
	marshalledData, _ := json.Marshal(resDTO.Data)
	err = json.Unmarshal(marshalledData, &valDTO)
	if err != nil {
		return nil, errs.NewAuthenticationError(err.Error())
	}

	return &AuthModel{
		UserId: valDTO.UserId,
		Role:   valDTO.Role,
	}, nil
}

func (h userHandler) ValidateAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := ValidateToken(c.Request.Header.Get("Authorization"))

		if err != nil {
			response := ResponseDTO{
				Status:  err.Code,
				Message: err.Message,
			}
			c.JSON(response.Status, response)
			c.Abort()
			return
		}

		response := ResponseDTO{
			Status:  http.StatusOK,
			Data:    res,
			Message: "User authenticated successfully",
		}
		c.JSON(response.Status, response)
	}
}
