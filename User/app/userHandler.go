package app

import (
	"net/http"
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
	return func (ctx *gin.Context) {
		var newUser userDTO

		if err := ctx.BindJSON(&newUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		role, err := domain.GetEnumByIndex(newUser.Role) 
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		user, err1 := h.userService.CreateUserInDynamodb(
			newUser.FirstName, 
			newUser.LastName, 
			newUser.Username, 
			newUser.Phone, 
			newUser.Email, 
			newUser.Password, 
			role,
		)
		
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, err1)
		} else {
			data, _ := user.MarshalJSON()
			ctx.Data(http.StatusCreated, "application/json", data)
		}
	}
}


func (h UserHandler) HandleGetUserByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		record, err := h.userService.GetUserById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": record})
	}
}