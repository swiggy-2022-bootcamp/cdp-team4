package domain

import "github.com/swiggy-2022-bootcamp/cdp-team4/transaction/utils/errs"

type TransactionService interface {
	CreateTransaction(string, int) (string, *errs.AppError)
	GetTransactionById(string) (*Transaction, *errs.AppError)
	GetTransactionByUserId(string) (*Transaction, *errs.AppError)
	UpdateTransactionByUserId(string, int) (bool, *errs.AppError)
}

type service struct {
	transactionRepository TransactionRepository
}

func (s service) CreateTransaction(userId string, transactionPoints int) (string, *errs.AppError) {
	transaction := NewTransaction(userId, transactionPoints)
	resultId, err := s.transactionRepository.InsertTransaction(*transaction)
	if err != nil {
		return "", err
	}
	return resultId, nil
}

func (s service) GetTransactionById(transactionId string) (*Transaction, *errs.AppError) {
	res, err := s.transactionRepository.FindTransactionById(transactionId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) GetTransactionByUserId(userId string) (*Transaction, *errs.AppError) {
	res, err := s.transactionRepository.FindTransactionByUserId(userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) UpdateTransactionByUserId(userId string, points int) (bool, *errs.AppError) {
	_, err := s.transactionRepository.UpdateTransactionByUserId(userId, points)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewTransactionService(transactionRepository TransactionRepository) TransactionService {
	return &service{
		transactionRepository: transactionRepository,
	}
}
