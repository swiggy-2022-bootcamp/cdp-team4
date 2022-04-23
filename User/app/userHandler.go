package app

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"
	pb "github.com/swiggy-2022-bootcamp/cdp-team4/user/app/protobuf"
	"fmt"
	"google.golang.org/grpc"
	"context"

)

type UserHandler struct {
	UserService domain.UserService
}

type ShippingAddressDTO struct {
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	City      string    `json:"city"`
	Address1  string    `json:"address_1"`
	Address2  string    `json:"address_2"`
	CountryID uint32    `json:"country_id"`
	PostCode  uint32    `json:"postcode"`
}

type UserDTO struct {
	FirstName string 				`json:"first_name"`
	LastName  string 				`json:"last_name"`
	Username  string 				`json:"username"`
	Password  string 				`json:"password"`
	Phone     string 				`json:"phone"`
	Email     string 				`json:"email"`
	Role      int    				`json:"role"`
	Address   ShippingAddressDTO 	`json:"address"`
	Fax		  string 				`json:"fax"`
}

type GrpcHelper interface {
	GetShippingAddressId(ShippingAddressDTO) (string)
}

// @Summary Create User
// @Description To register a new user for the app.
// @Tags User
// @Schemes
// @Accept json
// @Produce json
// @Param        user	body	UserDTO  true  "User structure"
// @Success	201  {string} 	http.StatusCreated
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Router /user [POST]
func (h UserHandler) HandleUserCreation() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		var newUser UserDTO

		if err := ctx.BindJSON(&newUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("bind json")
			return
		}

		role, err := domain.GetEnumByIndex(newUser.Role) 
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("unable to get role")
			return
		}

		fmt.Println("role: ", role, err)

		shippingAddressId := GetShippingAddressId(newUser.Address)

		fmt.Println("shippingAddressId: ", shippingAddressId)

		user, err1 := h.UserService.CreateUserInDynamodb(
			newUser.FirstName, 
			newUser.LastName, 
			newUser.Username, 
			newUser.Phone, 
			newUser.Email, 
			newUser.Password, 
			role,
			// shippingAddressId,
			"abcid",
			newUser.Fax,
		)
		
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, err1)
			log.WithFields(logrus.Fields{"message": err1.Error(), "status": http.StatusInternalServerError}).
				Error("unable to create user")
		} else {
			data, _ := user.MarshalJSON()
			log.WithFields(logrus.Fields{"data": data, "status": http.StatusCreated}).
				Info("User Added")
			ctx.Data(http.StatusCreated, "application/json", data)
		}
	}
}

// @Summary Get User
// @Description To get user details.
// @Tags User
// @Schemes
// @Accept json
// @Param id path string true "User Name"
// @Produce json
// @Success	202  {object} 	domain.User
// @Failure	400  {number} 	http.StatusBadRequest
// @Security Bearer Token
// @Router /user/{id} [GET]
func (h UserHandler) HandleGetUserByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		record, err := h.UserService.GetUserById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to get user by id")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": record})
		log.WithFields(logrus.Fields{"data": record, "status": http.StatusAccepted}).
				Info("User fetched")
	}
}


// @Summary Get all User details
// @Description To get every user detail.
// @Tags User
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {array} 	domain.User
// @Failure	400  {number} 	http.StatusBadRequest
// @Router /users [GET]
func (h UserHandler) HandleGetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		records, err := h.UserService.GetAllUsers()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to get users")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"records": records})
		log.WithFields(logrus.Fields{"data": records, "status": http.StatusAccepted}).
				Info("Fetched users")
	}
}


// @Summary Update User
// @Description To update user
// @Tags User
// @Schemes
// @Accept json
// @Param id path string true "User Name"
// @Param        user	body	domain.User  true  "User structure"
// @Produce json
// @Success	202  {string} 	domain.User
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /user/{id} [PATCH]
func (h UserHandler) HandleUpdateUserByID() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		var newUpdatedUser UserDTO
		userId := ctx.Param("id")

		if err := ctx.BindJSON(&newUpdatedUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to get user by id")
			return
		}

		role, err := domain.GetEnumByIndex(newUpdatedUser.Role) 
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("unable to get role")
			return
		} 

		oldUser, err2 := h.UserService.GetUserById(userId)

		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err2.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("unable to user with given id")
			return
		}

		isAddressUpdated := UpdateShippingAddressId(newUpdatedUser.Address, oldUser.AddressID)

		if isAddressUpdated != true {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to update address"})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusInternalServerError}).
				Error("unable to update address")
			return
		}

		ok, err1 := h.UserService.UpdateUserById(
			userId,
			newUpdatedUser.FirstName, 
			newUpdatedUser.LastName, 
			newUpdatedUser.Username, 
			newUpdatedUser.Phone, 
			newUpdatedUser.Email, 
			newUpdatedUser.Password, 
			role,
			oldUser.AddressID,
			newUpdatedUser.Fax,
		)
		
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, err1)
			log.WithFields(logrus.Fields{"message": err1.Error(), "status": http.StatusInternalServerError}).
				Error("unable to update user")
		} else {
			ctx.JSON(http.StatusAccepted, gin.H{"message": "user updated"})
			log.WithFields(logrus.Fields{"is_updated": ok, "status": http.StatusAccepted}).
				Info("User updated")
		}
	}
}


// @Summary Delete User
// @Description To remove a particular user.
// @Tags User
// @Schemes
// @Accept json
// @Param id path string true "User Name"
// @Produce json
// @Success	202  {string} 	http.StatusAccepted
// @Failure	400  {number} 	http.StatusBadRequest
// @Router /user/{id} [DELETE]
func (h UserHandler) HandleDeleteUserByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		ok, err := h.UserService.DeleteUserById(id)

		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to delete user")
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "user deleted"})
		log.WithFields(logrus.Fields{"is_deleted": ok, "status": http.StatusAccepted}).
				Info("User deleted")
	}
}


func GetShippingAddressId(address ShippingAddressDTO) (string){
	// Set up connection with the grpc server

	conn, err := grpc.Dial("localhost:7776", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error while making connection, %v\n", err)
	}

	// Create a client instance
	c := pb.NewShippingClient(conn)

	// Lets invoke the remote function from client on the server
	resp, err1 := c.AddShippingAddress(
		context.Background(),
		&pb.ShippingAddressAddRequest{
			Firstname: address.FirstName,
			Lastname:  address.LastName,
			City:      address.City,
			Address1:  address.Address1,
			Address2:  address.Address2,
			Countryid: address.CountryID,
			Postcode:  address.PostCode,
		},
	)
	
	if err1 != nil {
		log.WithFields(logrus.Fields{"message": err1.Error(), "status": http.StatusInternalServerError}).
				Error("unable to add address")
		return "abcid"
	}

	return resp.ShippingAddressID
}


func UpdateShippingAddressId(address ShippingAddressDTO, id string) (bool){
	// Set up connection with the grpc server

	conn, err := grpc.Dial("localhost:7776", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error while making connection, %v\n", err)
	}

	// Create a client instance
	c := pb.NewShippingClient(conn)

	// Lets invoke the remote function from client on the server
	_, err1 := c.UpdateShippingAddress(
		context.Background(),
		&pb.ShippingAddressUpdateRequest{
			ShippingAddressID: id,
			Firstname: address.FirstName,
			Lastname:  address.LastName,
			City:      address.City,
			Address1:  address.Address1,
			Address2:  address.Address2,
			Countryid: address.CountryID,
			Postcode:  address.PostCode,
		},
	)
	
	if err1 != nil {
		log.WithFields(logrus.Fields{"message": err1.Error(), "status": http.StatusInternalServerError}).
				Error("unable to update address")
		return false
	}

	return true
}
