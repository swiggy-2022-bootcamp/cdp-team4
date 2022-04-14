package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentService interface {
	CreateDynamoPaymentRecord(
		int16,
		string,
		string,
		string,
		string,
		string,
		string,
		string,
		[]string,
	) (bool, error)
	GetPaymentRecordById(string) (Payment, error)
	GetPaymentAllRecordsByUserId(string) ([]Payment, error)
	UpdatePaymentStatus(string, string) (bool, error)
	UpdatePaymentMethod(string, string) (bool, error)
	GetPaymentMethods(string) ([]string, error)
}

type paymentService struct {
	PaymentDynamoRepository PaymentDynamoRepository
}

func _generateUniqueId() string {
	return primitive.NewObjectID().String()
}

func (service paymentService) CreateDynamoPaymentRecord(
	amount int16,
	currency, status, order_id, user_id, method, description, vpa string,
	notes []string,
) (bool, error) {
	id := _generateUniqueId()
	paymentRecord := Payment{
		Id:          id,
		Amount:      amount,
		Currency:    currency,
		Status:      status,
		OrderID:     order_id,
		UserID:      user_id,
		Method:      method,
		Description: description,
		VPA:         vpa,
		Notes:       notes,
	}

	ok, err := service.PaymentDynamoRepository.Insert(paymentRecord)
	if !ok {
		return false, err
	}
	return true, nil
}

func (service paymentService) GetPaymentRecordById(id string) (Payment, error) {
	return Payment{}, nil
}

func (service paymentService) GetPaymentAllRecordsByUserId(id string) ([]Payment, error) {
	return []Payment{}, nil
}

func (service paymentService) UpdatePaymentMethod(id, method string) (bool, error) {
	return true, nil
}

func (service paymentService) GetPaymentMethods(id string) ([]string, error) {
	return []string{}, nil
}

func (service paymentService) UpdatePaymentStatus(
	paymentID string,
	paymentStatus string,
) (bool, error) {
	return true, nil
}

func NewPaymentService(paymentDynamoRepository PaymentDynamoRepository) PaymentService {
	return &paymentService{
		PaymentDynamoRepository: paymentDynamoRepository,
	}
}
