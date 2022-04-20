package app

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/infra/gokafka"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PayHandler struct {
	PaymentService domain.PaymentService
}

type PaymentRecordDTO struct {
	Amount      int16
	Currency    string
	Status      string
	OrderID     string
	UserID      string
	Method      string
	Description string
	VPA         string
	Notes       []string
}

type PaymentMethodDTO struct {
	Id      string
	Method  string
	Agree   string
	Comment string
}

type UpdatePayStatusDTO struct {
	Id     string
	Status string
}

func GenerateUniqueId() string {
	return primitive.NewObjectID().Hex()
}

func simulatePaymentDone(data interface{}) {
	time.Sleep(10 * time.Second)
	ok, err := gokafka.WriteMsgToKafka("payment", data)
	if !ok {
		log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
			Error("unable to write message to kafka")
	}
}

// Handle pay
// @Summary Initiate the payment process
// @Description Returns razorpay payment link with other details to the user
// @Tags Payment
// @Schemes
// @Accept json
// @Produce json
// @Param   req  body PaymentRecordDTO true "Payment details"
// @Success	200  string		Payment record added successfully
// @Failure 400  string   	Bad request
// @Failure 500  string		Internal Server Error
// @Router /pay/ [POST]
func (ph PayHandler) HandlePay() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paymentDto PaymentRecordDTO

		if err := ctx.BindJSON(&paymentDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("bind json")
			return
		}
		id := GenerateUniqueId()

		data, err := ph.PaymentService.CreateDynamoPaymentRecord(
			id,
			paymentDto.Amount,
			paymentDto.Currency,
			paymentDto.Status,
			paymentDto.OrderID,
			paymentDto.UserID,
			paymentDto.Method,
			paymentDto.Description,
			paymentDto.VPA,
			paymentDto.Notes,
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("create payment record")
			return
		}
		if data == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to create payment link"})
			log.WithFields(logrus.Fields{"message": err, "status": http.StatusBadRequest}).
				Error("unable create payment link")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "payment record added", "data": data})
		log.WithFields(logrus.Fields{"data": data, "status": http.StatusOK}).
			Info("payment record added")

		// simulatoin of successful payment
		go simulatePaymentDone(data)
	}
}

// Get payment record by Id
// @Summary get payment record by Id
// @Description Get payment record by Id
// @Tags Payment
// @Schemes
// @Accept json
// @Produce json
// @Param   req  query int true "Payment id"
// @Success	200  {object}		domain.Payment
// @Failure 400  string   	Bad request
// @Failure 500  string		Internal Server Error
// @Router /pay/ [GET]
func (ph PayHandler) HandleGetPayRecordByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		record, err := ph.PaymentService.GetPaymentRecordById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to fetch payment record")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"record": record})
		log.WithFields(logrus.Fields{"record": record, "status": http.StatusOK}).
			Info("payment record by id")
	}
}

func (ph PayHandler) handleGetPayRecordsByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// id := ctx.Param("user_id")
		// record, err := ph.PaymentService.GetPaymentAllRecordsByUserId(id)

		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		// 	return
		// }
		// ctx.JSON(http.StatusAccepted, gin.H{"record": record})
	}
}

// Update payment status
// @Summary update payment status
// @Description update payment status
// @Tags Payment
// @Schemes
// @Accept json
// @Produce json
// @Param   req  body UpdatePayStatusDTO true "Payment status"
// @Success	200  string		Payment status update successfully
// @Failure 400  string   	Bad request
// @Failure 500  string		Internal Server Error
// @Router /pay/ [PUT]
func (ph PayHandler) HandleUpdatePayStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updatePayStatusDTO UpdatePayStatusDTO

		if err := ctx.BindJSON(&updatePayStatusDTO); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to json bind")
			return
		}

		ok, err := ph.PaymentService.UpdatePaymentStatus(
			updatePayStatusDTO.Id,
			updatePayStatusDTO.Status,
		)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable update payment status")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "payment record updated"})
		log.WithFields(logrus.Fields{"status": http.StatusOK}).
			Info("payment record updated")
	}
}

// Add payment method
// @Summary add payment method
// @Description add new payment method
// @Tags PaymentMethod
// @Schemes
// @Accept json
// @Produce json
// @Param   req  body PaymentMethodDTO true "Payment method details"
// @Success	200  string		Payment method added successfully
// @Failure 400  string   	Bad request
// @Failure 500  string		Internal Server Error
// @Router /pay/paymentMethods [POST]
func (ph PayHandler) HandleAddPaymentMethods() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paymentMethodDTO PaymentMethodDTO
		if err := ctx.BindJSON(&paymentMethodDTO); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to json bind")
			return
		}

		ok, err := ph.PaymentService.AddPaymentMethod(
			paymentMethodDTO.Id,
			paymentMethodDTO.Method,
			paymentMethodDTO.Agree,
			paymentMethodDTO.Comment,
		)

		if !ok {
			if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "method already exists"})
				log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
					Debug("method already exists")
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to add payment method")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "payment method added"})
		log.WithFields(logrus.Fields{"status": http.StatusOK}).
			Info("payment method added")
	}
}

// Handle get payment methods
// @Summary Get supported payment methods for user
// @Description Get supported payment methods for user
// @Tags PaymentMethod
// @Schemes
// @Accept json
// @Produce json
// @Param   req  query 		int true "User Id"
// @Success	200  {object}	[]string
// @Failure 400  string   	Bad request
// @Failure 500  string		Internal Server Error
// @Router /pay/paymentMethods/:id [GET]
func (ph PayHandler) HandleGetPaymentMethods() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		methods, err := ph.PaymentService.GetPaymentMethods(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error("unable to fetch payment methods")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"methods": methods})
		log.WithFields(logrus.Fields{"methods": methods, "status": http.StatusOK}).
			Info("fetch payment methods")
	}
}

func NewPaymentHandler(paymentService domain.PaymentService) PayHandler {
	return PayHandler{
		PaymentService: paymentService,
	}
}
