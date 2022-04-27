package domain

import (
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/utils/errs"
)

type Transaction struct {
	Id                string `json:"id"`
	UserID            string `json:"user_id"`
	TransactionPoints int    `json:"transaction_points"`
}

type TransactionRepository interface {
	InsertTransaction(Transaction) (string, *errs.AppError)
	FindTransactionById(string) (*Transaction, *errs.AppError)
	FindTransactionByUserId(string) (*Transaction, *errs.AppError)
	UpdateTransactionByUserId(string, int) (bool, *errs.AppError)
}

func NewTransaction(userId string, transactionpoints int) *Transaction {
	return &Transaction{
		UserID:            userId,
		TransactionPoints: transactionpoints,
	}
}
