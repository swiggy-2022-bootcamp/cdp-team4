package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/auth/utils/errs"
	"net/http"
	"strings"
)

type AuthHandler struct {
	authService domain.AuthService
}

type loginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetAuthToken @Schemes
// @Description Creates a auth token upon login
// @Tags users
// @Produce json
// @Accept json
// @Param        login-credentials  body loginDTO true "User login"
// @Success 200 {object} userResponseDTO
// @Router /login [post]
func (ah AuthHandler) GetAuthToken(c *gin.Context) {

	var credentials loginDTO
	err := json.NewDecoder(c.Request.Body).Decode(&credentials)

	if err != nil {
		customErr := errs.NewValidationError("Invalid request paylaod")
		responseDto := ResponseDTO{
			Status:  customErr.Code,
			Message: customErr.Message,
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	var username = credentials.Username
	var password = credentials.Password

	authToken, _ := ah.authService.GenerateAuthToken(username, password)
	responseDto := ResponseDTO{
		Status: http.StatusOK,
		Data:   authToken,
	}

	c.JSON(responseDto.Status, responseDto)
}

func (ah AuthHandler) ValidateAuthToken(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	authToken := splitToken[1]
	authModel, err := ah.authService.ValidateAuthToken(authToken)

	if err != nil {
		customErr := errs.NewAuthenticationError("Invalid token, Access Denied")
		c.JSON(http.StatusUnauthorized, customErr)
		c.Abort()
		return
	}

	responseDto := ResponseDTO{
		Status:  http.StatusOK,
		Message: "Access Granted",
		Data: ValidationDTO{
			UserId: authModel.UserId,
			Role:   authModel.Role,
		},
	}

	c.JSON(responseDto.Status, responseDto)
}

func (ah AuthHandler) InvalidateAuthToken(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	authToken := splitToken[1]
	_, err := ah.authService.ValidateAuthToken(authToken)

	if err != nil {
		customErr := errs.NewValidationError("Invalid token, Access Denied")
		c.JSON(http.StatusUnauthorized, customErr)
		c.Abort()
		return
	}
	err = ah.authService.InvalidateAuthToken(authToken)
	if err != nil {
		customErr := errs.NewValidationError(err.Message)
		c.JSON(http.StatusInternalServerError, customErr)
		c.Abort()
		return
	}

	responseDto := ResponseDTO{
		Status:  http.StatusOK,
		Message: "Logged out successfully",
	}

	c.JSON(responseDto.Status, responseDto)

}
