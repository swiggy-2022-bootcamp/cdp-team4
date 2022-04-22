package domain_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/payment/mocks"
)

func TestGenerateUniqueId(t *testing.T) {
	var id interface{} = domain.GenerateUniqueId()
	_, ok := id.(string)

	assert.Equal(t, true, ok)
}

func TestGetRazorpayPaymentLink(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockDynamoRepo := mocks.NewMockPaymentDynamoRepository(mockCtrl)

	service := domain.NewPaymentService(mockDynamoRepo)

	response, err := service.GetRazorpayPaymentLink(domain.Payment{
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
		Id:          domain.GenerateUniqueId(),
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

	testcases := []struct {
		name       string
		createStub func(mocks.MockPaymentDynamoRepository)
		assertTest func(*testing.T, map[string]interface{}, error)
	}{
		{
			name: "FailCreateDynamoPaymentRecord",
			createStub: func(mpdr mocks.MockPaymentDynamoRepository) {
				mpdr.EXPECT().
					InsertPaymentRecord(payment).
					Return(false, fmt.Errorf("unable to insert record"))
			},
			assertTest: func(t *testing.T, m map[string]interface{}, err error) {
				assert.Nil(t, m)
				assert.NotNil(t, err)
			},
		},
		{
			name: "FailCreateDynamoPaymentRecordGenerateLink",
			createStub: func(mpdr mocks.MockPaymentDynamoRepository) {
				mpdr.EXPECT().InsertPaymentRecord(payment).Return(true, nil)
			},
			assertTest: func(t *testing.T, m map[string]interface{}, err error) {
				assert.Nil(t, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockPaymentDynamoRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewPaymentService(mockDynamoRepo)

			// mockDynamoRepo.On("InsertPaymentRecord", mock.Anything).Return(true, nil)
			data, err := service.CreateDynamoPaymentRecord(
				payment.Id,
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

			testcase.assertTest(t, data, err)
		})
	}

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

	testcases := []struct {
		name       string
		createStub func(mocks.MockPaymentDynamoRepository)
		assertTest func(*testing.T, *domain.Payment, error)
	}{
		{
			name: "SuccessGetPaymentRecordById",
			createStub: func(mpdr mocks.MockPaymentDynamoRepository) {
				mpdr.EXPECT().FindPaymentRecordById("xyx").Return(&payment, nil)
			},
			assertTest: func(t *testing.T, m *domain.Payment, err error) {
				assert.NotNil(t, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailGetPaymentRecordById",
			createStub: func(mpdr mocks.MockPaymentDynamoRepository) {
				mpdr.EXPECT().
					FindPaymentRecordById("xyx").
					Return(nil, fmt.Errorf("unable to find record"))
			},
			assertTest: func(t *testing.T, m *domain.Payment, err error) {
				assert.Nil(t, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockPaymentDynamoRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewPaymentService(mockDynamoRepo)

			data, err := service.GetPaymentRecordById("xyx")

			testcase.assertTest(t, data, err)
		})
	}

}

func TestGetPaymentMethods(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(mocks.MockPaymentDynamoRepository)
		assertTest func(*testing.T, []string, error)
	}{
		{
			name: "SuccessGetPaymentMethods",
			createStub: func(mpdr mocks.MockPaymentDynamoRepository) {
				mpdr.EXPECT().GetPaymentMethods("xyx").Return([]string{""}, nil)
			},
			assertTest: func(t *testing.T, m []string, err error) {
				assert.NotNil(t, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailGetPaymentMethods",
			createStub: func(mpdr mocks.MockPaymentDynamoRepository) {
				mpdr.EXPECT().
					GetPaymentMethods("xyx").
					Return(nil, fmt.Errorf("unable to find record"))
			},
			assertTest: func(t *testing.T, m []string, err error) {
				assert.Nil(t, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockPaymentDynamoRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewPaymentService(mockDynamoRepo)

			data, err := service.GetPaymentMethods("xyx")

			testcase.assertTest(t, data, err)
		})
	}

}

func TestUpdatePaymentStatus(t *testing.T) {
	t.Run("updatePaymentStatus", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		mockDynamoRepo := mocks.NewMockPaymentDynamoRepository(mockCtrl)

		service := domain.NewPaymentService(mockDynamoRepo)
		ok, err := service.UpdatePaymentStatus("xyx", "confirmed")

		assert.Equal(t, true, ok)
		assert.Nil(t, err)
	})
}

func TestAddPaymentMethod(t *testing.T) {
	var paymentRecord = domain.PaymentMethod{
		Id:      "id",
		Agree:   "agree",
		Comment: "comment",
		Method:  []string{"method"},
	}

	testcases := []struct {
		name       string
		createStub func(mocks.MockPaymentDynamoRepository)
		assertTest func(*testing.T, bool, error)
	}{
		{
			name: "SuccessAddPaymentMethodUpdateState",
			createStub: func(mpdr mocks.MockPaymentDynamoRepository) {
				mpdr.EXPECT().GetPaymentMethods("id").Return([]string{""}, nil)
				mpdr.EXPECT().UpdatePaymentMethods("id", "method").Return(true, nil)
			},
			assertTest: func(t *testing.T, m bool, err error) {
				assert.Equal(t, true, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailAddPaymentMethodUpdateState",
			createStub: func(mpdr mocks.MockPaymentDynamoRepository) {
				mpdr.EXPECT().GetPaymentMethods("id").Return([]string{""}, nil)
				mpdr.EXPECT().
					UpdatePaymentMethods("id", "method").
					Return(false, fmt.Errorf("unable to update methods"))
			},
			assertTest: func(t *testing.T, m bool, err error) {
				assert.Equal(t, false, m)
				assert.NotNil(t, err)
			},
		},
		{
			name: "SuccessAddPaymentMethodInsertState",
			createStub: func(mpdr mocks.MockPaymentDynamoRepository) {
				mpdr.EXPECT().
					GetPaymentMethods("id").
					Return([]string{""}, fmt.Errorf("method not found"))
				mpdr.EXPECT().InsertPaymentMethod(paymentRecord).Return(true, nil)
			},
			assertTest: func(t *testing.T, m bool, err error) {
				assert.Equal(t, true, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailAddPaymentMethodInsertState",
			createStub: func(mpdr mocks.MockPaymentDynamoRepository) {
				mpdr.EXPECT().
					GetPaymentMethods("id").
					Return([]string{""}, fmt.Errorf("method not found"))
				mpdr.EXPECT().
					InsertPaymentMethod(paymentRecord).
					Return(false, fmt.Errorf("unable to insert"))
			},
			assertTest: func(t *testing.T, m bool, err error) {
				assert.Equal(t, false, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockPaymentDynamoRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewPaymentService(mockDynamoRepo)

			data, err := service.AddPaymentMethod("id", "method", "agree", "comment")

			testcase.assertTest(t, data, err)
		})
	}
}
