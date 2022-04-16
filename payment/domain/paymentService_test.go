package domain_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/payment/mocks/domain"
)

var mockDynamoRepo = mocks.PaymentDynamoRepository{}
var service = domain.NewPaymentService(&mockDynamoRepo)

func TestGenerateUniqueId(t *testing.T) {
	var id interface{} = domain.GenerateUniqueId()
	_, ok := id.(string)

	assert.Equal(t, true, ok)
}

func TestGetRazorpayPaymentLink(t *testing.T) {
	response, err := domain.GetRazorpayPaymentLink(domain.Payment{
		Amount:   45,
		Currency: "INR",
		UserID:   "adf",
		OrderID:  "asg",
		Notes:    []string{""},
	})

	assert.Nil(t, nil, response)
	assert.NotNil(t, err)
}

func TestShouldReturnNewUserService(t *testing.T) {
	userService := domain.NewPaymentService(nil)
	assert.NotNil(t, userService)
}

func TestCreateDynamoPaymentRecord(t *testing.T) {
	payment := domain.Payment{
		Amount:      54,
		Currency:    "INR",
		Status:      "pending",
		Method:      "upi",
		Description: "description",
		VPA:         "asdf",
		UserID:      "dfa",
		OrderID:     "isf",
		Notes:       []string{""},
	}

	mockDynamoRepo.On("InsertPaymentRecord", mock.Anything).Return(true, nil)
	service.CreateDynamoPaymentRecord(
		payment.Amount,
		payment.Currency,
		payment.Status,
		payment.OrderID,
		payment.UserID,
		payment.Method,
		payment.Description,
		payment.VPA,
		payment.Notes,
	)
	mockDynamoRepo.AssertNumberOfCalls(t, "InsertPaymentRecord", 1)
}

func TestGetPaymentRecordById(t *testing.T) {
	payment := domain.Payment{
		Amount:      54,
		Currency:    "INR",
		Status:      "pending",
		Method:      "upi",
		Description: "description",
		VPA:         "asdf",
		UserID:      "dfa",
		OrderID:     "isf",
		Notes:       []string{""},
	}
	mockDynamoRepo.On("FindPaymentRecordById", "abc").Return(&payment, nil)
	service.GetPaymentRecordById("abc")

	mockDynamoRepo.AssertNumberOfCalls(t, "FindPaymentRecordById", 1)
}

func TestFailGetPaymentRecordById(t *testing.T) {
	mockDynamoRepo.On("FindPaymentRecordById", "abcd").Return(nil, fmt.Errorf("element id not found"))
	service.GetPaymentRecordById("abcd")

	mockDynamoRepo.AssertNumberOfCalls(t, "FindPaymentRecordById", 2)
}

func TestGetPaymentMethods(t *testing.T) {
	methodList := []string{"upi"}
	mockDynamoRepo.On("GetPaymentMethods", "abd").Return(methodList, nil)
	service.GetPaymentMethods("abd")

	mockDynamoRepo.AssertNumberOfCalls(t, "GetPaymentMethods", 1)
}

func TestFailGetPaymentMethods(t *testing.T) {
	mockDynamoRepo.On("GetPaymentMethods", "abdc").Return(nil, fmt.Errorf("element id not found"))
	service.GetPaymentMethods("abdc")

	mockDynamoRepo.AssertNumberOfCalls(t, "GetPaymentMethods", 2)
}

func TestUpdatePaymentStatus(t *testing.T) {
	ok, _ := service.UpdatePaymentStatus("abc", "pending")
	assert.Equal(t, true, ok)
}

func TestAddPaymentMethod(t *testing.T) {
	mockDynamoRepo.On("GetPaymentMethods", "abc").Return(nil, fmt.Errorf("element id not found"))
	mockDynamoRepo.On("InsertPaymentMethod", mock.Anything).Return(true, nil)
	mockDynamoRepo.On("UpdatePaymentMethods", "abc", "upi").Return(true, nil)

	service.AddPaymentMethod("abc", "upi", "1", "none")
	// mockDynamoRepo.AssertNumberOfCalls(t, "UpdatePaymentMethods", 1)
	mockDynamoRepo.AssertNumberOfCalls(t, "InsertPaymentMethod", 1)
}

func TestAddNextPaymentMethod(t *testing.T) {
	mockDynamoRepo.On("GetPaymentMethods", "xyz").Return(nil, nil)
	mockDynamoRepo.On("InsertPaymentMethod", mock.Anything).Return(true, nil)
	mockDynamoRepo.On("UpdatePaymentMethods", "xyz", "upi").Return(true, nil)

	service.AddPaymentMethod("xyz", "upi", "1", "none")
	mockDynamoRepo.AssertNumberOfCalls(t, "UpdatePaymentMethods", 1)
}
