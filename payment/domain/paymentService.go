package domain

import (
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	razorpay "github.com/razorpay/razorpay-go"
)

// interface wraps all the methods should be implemented on service layer
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
	UpdatePaymentStatus(string, string) (bool, error)
	GetPaymentMethods(string) ([]string, error)
	AddPaymentMethod(string, string, string, string) (bool, error)
	GetRazorpayPaymentLink(Payment) (map[string]interface{}, error)
}

type paymentService struct {
	PaymentDynamoRepository PaymentDynamoRepository
}

// function to generate unique id which internally uses the primitive's Object id
// that is used in MongoDb to automatically create an ID.
func GenerateUniqueId() string {
	return primitive.NewObjectID().Hex()
}

// function that returns Razorpay's payment link on the basis of the
// user details
//
// https://razorpay.com/docs/api/payments/payment-links/#create-payment-link
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

	// to get razopay payment link, we have to send request data to the api
	/*
		data format:
				{
		  "amount": 1000,
		  "currency": "INR",
		  "accept_partial": true,
		  "first_min_partial_amount": 100,
		  "expire_by": 1691097057,
		  "reference_id": "TS1989",
		  "description": "Payment for policy no #23456",
		  "customer": {
		    "name": "Gaurav Kumar",
		    "contact": "+919999999999",
		    "email": "gaurav.kumar@example.com"
		  },
		  "notify": {
		    "sms": true,
		    "email": true
		  },
		  "reminder_enable": true,
		  "notes": {
		    "policy_name": "Jeevan Bima"
		  },
		  "callback_url": "https://example-callback-url.com/",
		  "callback_method": "get"
		}
	*/
	data := gin.H{
		"amount":       p.Amount,
		"currency":     p.Currency,
		"reference_id": GenerateUniqueId(),
		"customer": struct {
			userId  string
			orderId string
		}{
			userId:  p.UserID,
			orderId: p.OrderID,
		},
	}
	body, err := client.PaymentLink.Create(data, nil)

	if err != nil {
		return nil, err
	}
	return body, err
}

// method to insert the payment record and return the razorpay payment link
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

// function to get payment record using record id given in request
func (service paymentService) GetPaymentRecordById(id string) (*Payment, error) {
	paymentRecord, err := service.PaymentDynamoRepository.FindPaymentRecordById(id)
	if err != nil {
		return nil, err
	}
	return paymentRecord, nil
}

// function to return all the payment methods supported for particular user using give
// user id in request
func (service paymentService) GetPaymentMethods(id string) ([]string, error) {
	methods, err := service.PaymentDynamoRepository.GetPaymentMethods(id)
	if err != nil {
		return nil, err
	}

	return methods, nil
}

// function to update payment status that can be confirm, pending or failed.
func (service paymentService) UpdatePaymentStatus(
	paymentID string,
	paymentStatus string,
) (bool, error) {
	return true, nil
}

// function to add payment method that is going to support for particular user
// ideally, this thing we be done after user is verfied for that method.
//
// if record is already present, just need to update the method array using UpdatePaymentMethods
// otherwise,
// insert a new record by calling InsertPaymentMethod.
func (service paymentService) AddPaymentMethod(
	id, method, agree, comment string,
) (bool, error) {
	var paymentRecord = PaymentMethod{
		Id:      id,
		Agree:   agree,
		Comment: comment,
		Method:  []string{method},
	}
	// get all the current payment methods supported for the user using id
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

// constructor method of payment service
func NewPaymentService(paymentDynamoRepository PaymentDynamoRepository) PaymentService {
	return &paymentService{
		PaymentDynamoRepository: paymentDynamoRepository,
	}
}
