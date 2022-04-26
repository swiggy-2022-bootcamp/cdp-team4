package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/domain"
)

type TransactionHandler struct {
	TransactionService domain.TransactionService
}

type TransactionRecordDTO struct {
	UserID            string `json:"user_id"`
	TransactionPoints int    `json:"transaction_points"`
}

// Get Transaction by userID
// @Summary      Get Transaction by userId
// @Description  This Handle returns Transaction given userId
// @Tags         Transaction
// @Produce      json
// @Param   req  query int true "User id"
// @Success      200  {object}  domain.Transaction
// @Failure  400 string   	Bad request
// @Failure  500  string		Internal Server Error
// @Router       /transaction/:userId    [get]
func (th TransactionHandler) HandleGetTransactionRecordByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("userId")
		if id == "" {
			log.WithFields(logrus.Fields{"message": "Invalid User ID", "status": http.StatusBadRequest}).
				Error("Error while Getting transaction by transaction id")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid User ID"})
			return
		}
		res, err := th.TransactionService.GetTransactionByUserId(id)

		if err != nil {
			log.WithFields(logrus.Fields{"message": err.Message, "status": http.StatusBadRequest}).
				Error("Record not found")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found: " + err.Message})
			return
		}
		transactiondto := convertTransactionModeltoTransactionDTO(*res)
		ctx.JSON(http.StatusAccepted, gin.H{"record": transactiondto})
	}
}

// Update transaction for a userId
// @Summary      Update transaction points for a userId
// @Description  This Handle Update transaction given user id
// @Tags         Transaction
// @Schemes
// @Accept json
// @Produce json
// @Param   req  body TransactionRecordDTO true "Transaction details"
// @Success      200  string  transaction record updated
// @Failure  400 string   	Bad request
// @Failure 500  string		Internal Server Error
// @Router       /transaction/:userId   [put]
func (th TransactionHandler) HandleUpdateTransactionByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestDTO struct {
			UserID            string `json:"user_id"`
			TransactionPoints int    `json:"transaction_points"`
		}
		if err := ctx.BindJSON(&requestDTO); err != nil {
			log.WithFields(logrus.Fields{"message": err.Error(), "status": http.StatusBadRequest}).
				Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ok, err := th.TransactionService.UpdateTransactionByUserId(requestDTO.UserID, requestDTO.TransactionPoints)
		if !ok {
			log.WithFields(logrus.Fields{"message": err.Message, "status": http.StatusBadRequest}).
				Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Message})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "transaction record updated"})
	}
}

func convertTransactionModeltoTransactionDTO(transaction domain.Transaction) TransactionRecordDTO {
	return TransactionRecordDTO{
		UserID:            transaction.UserID,
		TransactionPoints: transaction.TransactionPoints,
	}
}
func NewTransactionHandler(transactionService domain.TransactionService) TransactionHandler {
	return TransactionHandler{
		TransactionService: transactionService,
	}
}
