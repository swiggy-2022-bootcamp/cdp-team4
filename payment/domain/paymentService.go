package domain

import "errors"

type PaymentService interface {
	CreateDynamoPaymentRecord(Payment) (bool, error)
	GetPaymentRecordById(string) (*Payment, error)
	GetPaymentAllRecordsByUserId(string) ([]*Payment, error)
	UpdatePaymentStatus(string, string) (*Payment, error)
}

type paymentService struct {
	PaymentDynamoRepository PaymentDynamoRepository
}

func (p paymentService) CreateDynamoPaymentRecord(payment Payment) (bool, error) {
	return true, errors.New("dummy!")
}

func (p paymentService) GetPaymentRecordById(id string) (*Payment, error) {
	return &Payment{}, errors.New("dummy!")

}

func (p paymentService) GetPaymentAllRecordsByUserId(id string) ([]*Payment, error) {
	return []*Payment{}, errors.New("dummy!")

}

func (p paymentService) UpdatePaymentStatus(
	paymentID string,
	paymentStatus string,
) (*Payment, error) {
	return &Payment{}, errors.New("dummy!")
}

func NewPaymentService(paymentDynamoRepository PaymentDynamoRepository) PaymentService {
	return &paymentService{
		PaymentDynamoRepository: paymentDynamoRepository,
	}
}
