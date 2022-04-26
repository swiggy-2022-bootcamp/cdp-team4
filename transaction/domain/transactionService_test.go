package domain_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/mocks"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/swiggy-2022-bootcamp/cdp-team4/transaction/utils/errs"
)

// func TestGetRazorpayTransactionLink(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	mockDynamoRepo := mocks.NewMockTransactionRepository(mockCtrl)

// 	service := domain.NewTransactionService(mockDynamoRepo)
// }

func TestShouldReturnNewService(t *testing.T) {
	newService := domain.NewTransactionService(nil)
	assert.NotNil(t, newService)
}

func TestCreateTransactionRecord(t *testing.T) {
	transaction := domain.Transaction{
		//Id:primitive.NewObjectID().Hex(),
		UserID:       "randomUserId",
		TransactionPoints: 10,
	}
	testcases := []struct {
		name       string
		createStub func(mocks.MockTransactionRepository)
		assertTest func(*testing.T, string, *errs.AppError)
	}{
		{
			name: "FailCreateTransactionRecord",
			createStub: func(mrr mocks.MockTransactionRepository) {
				errstring := "unable to insert record"
				mrr.EXPECT().
					InsertTransaction(transaction).
					Return("", &errs.AppError{Message: errstring})
			},
			assertTest: func(t *testing.T, m string, err *errs.AppError) {
				assert.Equal(t, "", m)
				assert.NotNil(t, err)
			},
		},
		{
			name: "SuccessCreateTransactionRecord",
			createStub: func(mrr mocks.MockTransactionRepository) {
				mrr.EXPECT().InsertTransaction(transaction).Return(transaction.Id, nil)
			},
			assertTest: func(t *testing.T, m string, err *errs.AppError) {
				assert.NotNil(t, m)
				assert.Nil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockTransactionRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewTransactionService(mockDynamoRepo)

			data, err := service.CreateTransaction(
				transaction.UserID,
				transaction.TransactionPoints,
			)

			testcase.assertTest(t, data, err)
		})
	}

}

func TestGetTransactionRecordByuserId(t *testing.T) {
	transaction := domain.Transaction{
		//Id:primitive.NewObjectID().Hex(),
		UserID:       "randomUserId",
		TransactionPoints: 10,
	}

	testcases := []struct {
		name       string
		createStub func(mocks.MockTransactionRepository)
		assertTest func(*testing.T, *domain.Transaction, *errs.AppError)
	}{
		{
			name: "SuccessGetTransactionRecordByUserId",
			createStub: func(mpdr mocks.MockTransactionRepository) {
				mpdr.EXPECT().FindTransactionByUserId("randomUserId").Return(&transaction, nil)
			},
			assertTest: func(t *testing.T, m *domain.Transaction, err *errs.AppError) {
				assert.NotNil(t, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailGetTransactionRecordByUserId",
			createStub: func(mpdr mocks.MockTransactionRepository) {
				errstring := "unable to find record"
				mpdr.EXPECT().
					FindTransactionByUserId("randomUserId").
					Return(nil, &errs.AppError{Message: errstring})
			},
			assertTest: func(t *testing.T, m *domain.Transaction, err *errs.AppError) {
				assert.Nil(t, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockTransactionRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewTransactionService(mockDynamoRepo)

			data, err := service.GetTransactionByUserId("randomUserId")

			testcase.assertTest(t, data, err)
		})
	}

}

func TestUpdateTransactionPoints(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(mocks.MockTransactionRepository)
		assertTest func(*testing.T, bool, *errs.AppError)
	}{
		{
			name: "SuccessUpdateTransaction",
			createStub: func(mpdr mocks.MockTransactionRepository) {
				mpdr.EXPECT().UpdateTransactionByUserId("randomUserId",20).Return(true, nil)
			},
			assertTest: func(t *testing.T, m bool, err *errs.AppError) {
				assert.Equal(t,true, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailUpdateTransaction",
			createStub: func(mpdr mocks.MockTransactionRepository) {
				errstring := "unable to find record"
				mpdr.EXPECT().
					UpdateTransactionByUserId("randomUserId",20).
					Return(false, &errs.AppError{Message: errstring})
			},
			assertTest: func(t *testing.T, m bool, err *errs.AppError) {
				assert.Equal(t,false, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockTransactionRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewTransactionService(mockDynamoRepo)

			data, err := service.UpdateTransactionByUserId("randomUserId", 20)

			testcase.assertTest(t, data, err)
		})
	}

}
