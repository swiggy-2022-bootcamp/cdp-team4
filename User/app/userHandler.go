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


func (h UserHandler) HandleGetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		records, err := h.userService.GetAllUsers()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"records": records})
	}
}


func (h UserHandler) HandleUpdateUserByID() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		var newUpdatedUser userDTO
		userId := ctx.Param("id")

		if err := ctx.BindJSON(&newUpdatedUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		role, err := domain.GetEnumByIndex(newUpdatedUser.Role) 
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		_, err1 := h.userService.UpdateUserById(
			userId,
			newUpdatedUser.FirstName, 
			newUpdatedUser.LastName, 
			newUpdatedUser.Username, 
			newUpdatedUser.Phone, 
			newUpdatedUser.Email, 
			newUpdatedUser.Password, 
			role,
		)
		
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, err1)
		} else {
			ctx.JSON(http.StatusAccepted, gin.H{"message": "user updated"})
		}
	}
}


func (h UserHandler) HandleDeleteUserByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		ok, err := h.userService.DeleteUserById(id)

		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": (*err).Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "user deleted"})
	}
}