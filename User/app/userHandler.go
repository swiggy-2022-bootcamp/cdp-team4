package app

import (
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"
)

type UserHandler struct {
	userService domain.UserService
}

type userDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
}

// Create User
// @Summary      Create Customer
// @Description  to create a customere
// @Tags         Customer
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /    [get]
func (h UserHandler) HandleUserCreation() gin.HandlerFunc {
	return func (c *gin.Context) {
		var newUser userDTO
		err := json.NewDecoder(c.Request.Body).Decode(&newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			role, err := domain.GetEnumByIndex(newUser.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
			}
			user, err1 := h.userService.CreateUserInDynamodb(newUser.FirstName, newUser.LastName, newUser.Username, newUser.Phone, newUser.Email, newUser.Password, role)
			if err1 != nil {
				c.JSON(http.StatusInternalServerError, err1)
			} else {
				data, _ := user.MarshalJSON()
				c.Data(http.StatusCreated, "application/json", data)
			}
		}
	}
}