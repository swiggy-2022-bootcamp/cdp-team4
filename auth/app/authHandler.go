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
	AuthService domain.AuthService
}

type LoginDTO struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// GetAuthToken @Schemes
// @Description Returns auth token upon login
// @Tags         auth
// @Produce json
// @Accept json
// @Param        login-credentials  body LoginDTO true "User login"
// @Success 200 {object} ResponseDTO
// @Router /login [post]
func (ah AuthHandler) GetAuthToken(c *gin.Context) {

	var credentials LoginDTO
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

	authToken, _ := ah.AuthService.GenerateAuthToken(username, password)
	responseDto := ResponseDTO{
		Status: http.StatusOK,
		Data:   authToken,
	}

	c.JSON(responseDto.Status, responseDto)
}

// ValidateAuthToken @Schemes
// @Description validates auth token
// @Tags         auth
// @Produce json
// @Accept json
// @Success 200 {object} ResponseDTO
// @Router /validate [get]
// @Security ApiKeyAuth
func (ah AuthHandler) ValidateAuthToken(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	authToken := splitToken[1]
	authModel, err := ah.AuthService.ValidateAuthToken(authToken)

	if err != nil {
		customErr := errs.NewAuthenticationError("Invalid token, Access Denied")
		responseDto := ResponseDTO{
			Status:  customErr.Code,
			Message: customErr.Message,
		}
		c.JSON(responseDto.Status, responseDto)
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

// InvalidateAuthToken @Schemes
// @Description Invalidates auth token
// @Tags         auth
// @Produce json
// @Accept json
// @Success 200 {object} ResponseDTO
// @Router /logout [post]
// @Security ApiKeyAuth
func (ah AuthHandler) InvalidateAuthToken(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	authToken := splitToken[1]
	_, err := ah.AuthService.ValidateAuthToken(authToken)

	if err != nil {
		customErr := errs.NewAuthenticationError("Invalid token, Access Denied")
		responseDto := ResponseDTO{
			Status:  customErr.Code,
			Message: customErr.Message,
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}
	err = ah.AuthService.InvalidateAuthToken(authToken)
	if err != nil {
		customErr := errs.NewValidationError(err.Message)
		responseDto := ResponseDTO{
			Status:  customErr.Code,
			Message: customErr.Message,
		}
		c.JSON(responseDto.Status, responseDto)
		c.Abort()
		return
	}

	responseDto := ResponseDTO{
		Status:  http.StatusOK,
		Message: "Logged out successfully",
	}

	c.JSON(responseDto.Status, responseDto)

}
