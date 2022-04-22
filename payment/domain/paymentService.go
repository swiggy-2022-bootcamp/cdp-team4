package domain

import (
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	razorpay "github.com/razorpay/razorpay-go"
)

type PaymentService interface {
	CreateDynamoPaymentRecord(
		string,
		int16,
		string,
		string,
		string,
		string,
		string,
		string,
		string,
		[]string,
	) (map[string]interface{}, error)
	GetPaymentRecordById(string) (*Payment, error)
	// GetPaymentAllRecordsByUserId(string) ([]Payment, error)
	UpdatePaymentStatus(string, string) (bool, error)
	// UpdatePaymentMethod(string, string) (bool, error)
	GetPaymentMethods(string) ([]string, error)
	AddPaymentMethod(string, string, string, string) (bool, error)
	GetRazorpayPaymentLink(Payment) (map[string]interface{}, error)
}

type paymentService struct {
	PaymentDynamoRepository PaymentDynamoRepository
}

func GenerateUniqueId() string {
	return primitive.NewObjectID().Hex()
}

func (service paymentService) GetRazorpayPaymentLink(
	p Payment,
) (map[string]interface{}, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}

	key_id := os.Getenv("RAZORPAY_KEY_ID")
	key_secret := os.Getenv("RAZORPAY_KEY_SECRET")
	client := razorpay.NewClient(key_id, key_secret)

	data := gin.H{
		"amount":       p.Amount,
		"currency":     p.Currency,
		"reference_id": GenerateUniqueId(),
		// "customer": struct {
		// 	userId  string
		// 	orderId string
		// }{
		// 	userId:  p.UserID,
		// 	orderId: p.OrderID,
		// },
		// "notes": p.Notes,
	}
	body, err := client.PaymentLink.Create(data, nil)

	if err != nil {
		return nil, err
	}
	return body, err
}

func (service paymentService) CreateDynamoPaymentRecord(
	id string,
	amount int16,
	currency, status, order_id, user_id, method, description, vpa string,
	notes []string,
) (map[string]interface{}, error) {
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

	ok, err := service.PaymentDynamoRepository.InsertPaymentRecord(paymentRecord)
	if !ok {
		return nil, err
	}

	data, err := service.GetRazorpayPaymentLink(paymentRecord)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service paymentService) GetPaymentRecordById(id string) (*Payment, error) {
	paymentRecord, err := service.PaymentDynamoRepository.FindPaymentRecordById(id)
	if err != nil {
		return nil, err
	}
	return paymentRecord, nil
}

// func (service paymentService) GetPaymentAllRecordsByUserId(id string) ([]Payment, error) {
// 	return []Payment{}, nil
// }

// func (service paymentService) UpdatePaymentMethod(id, method string) (bool, error) {
// 	return true, nil
// }

func (service paymentService) GetPaymentMethods(id string) ([]string, error) {
	methods, err := service.PaymentDynamoRepository.GetPaymentMethods(id)
	if err != nil {
		return nil, err
	}

	return methods, nil
}

func (service paymentService) UpdatePaymentStatus(
	paymentID string,
	paymentStatus string,
) (bool, error) {
	return true, nil
}

func (service paymentService) AddPaymentMethod(
	id, method, agree, comment string,
) (bool, error) {
	var paymentRecord = PaymentMethod{
		Id:      id,
		Agree:   agree,
		Comment: comment,
		Method:  []string{method},
	}
	_, err := service.PaymentDynamoRepository.GetPaymentMethods(id)

	if err != nil {
		ok, err := service.PaymentDynamoRepository.InsertPaymentMethod(paymentRecord)
		if !ok {
			return false, err
		}
		return ok, nil
	}

	ok, err := service.PaymentDynamoRepository.UpdatePaymentMethods(id, method)
	if !ok {
		return ok, err
	}

	return true, nil
}

func NewPaymentService(paymentDynamoRepository PaymentDynamoRepository) PaymentService {
	return &paymentService{
		PaymentDynamoRepository: paymentDynamoRepository,
	}
}
