package domain_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/mocks"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/swiggy-2022-bootcamp/cdp-team4/cart/utils/errs"
)


func TestShouldReturnNewService(t *testing.T) {
	newService := domain.NewCartService(nil)
	assert.NotNil(t, newService)
}

func TestCreateCartRecord(t *testing.T) {
	itemList := map[string]domain.Item{
		"1": domain.Item{Name: "Pen", Cost: 10, Quantity: 1},
		"2": domain.Item{Name: "Pencil", Cost: 5, Quantity: 2},
	}
	cart := domain.Cart{
		//Id:primitive.NewObjectID().Hex(),
		UserID: "randomUserId",
		Items:  itemList,
	}
	testcases := []struct {
		name       string
		createStub func(mocks.MockCartRepository)
		assertTest func(*testing.T, string, *errs.AppError)
	}{
		{
			name: "FailCreateCartRecord",
			createStub: func(mrr mocks.MockCartRepository) {
				errstring := "unable to insert record"
				mrr.EXPECT().
					InsertCart(cart).
					Return("", &errs.AppError{Message: errstring})
			},
			assertTest: func(t *testing.T, m string, err *errs.AppError) {
				assert.Equal(t, "", m)
				assert.NotNil(t, err)
			},
		},
		{
			name: "SuccessCreateCartRecord",
			createStub: func(mrr mocks.MockCartRepository) {
				mrr.EXPECT().InsertCart(cart).Return(cart.Id, nil)
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
			mockDynamoRepo := mocks.NewMockCartRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewCartService(mockDynamoRepo)

			data, err := service.CreateCart(
				cart.UserID,
				cart.Items,
			)

			testcase.assertTest(t, data, err)
		})
	}

}

func TestGetCartRecordByuserId(t *testing.T) {
	itemList := map[string]domain.Item{
		"1": domain.Item{Name: "Pen", Cost: 10, Quantity: 1},
		"2": domain.Item{Name: "Pencil", Cost: 5, Quantity: 2},
	}
	cart := domain.Cart{
		//Id:primitive.NewObjectID().Hex(),
		UserID: "randomUserId",
		Items:  itemList,
	}

	testcases := []struct {
		name       string
		createStub func(mocks.MockCartRepository)
		assertTest func(*testing.T, *domain.Cart, *errs.AppError)
	}{
		{
			name: "SuccessGetCartRecordByUserId",
			createStub: func(mpdr mocks.MockCartRepository) {
				mpdr.EXPECT().FindCartByUserId("randomUserId").Return(&cart, nil)
			},
			assertTest: func(t *testing.T, m *domain.Cart, err *errs.AppError) {
				assert.NotNil(t, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailGetCartRecordByUserId",
			createStub: func(mpdr mocks.MockCartRepository) {
				errstring := "unable to find record"
				mpdr.EXPECT().
					FindCartByUserId("randomUserId").
					Return(nil, &errs.AppError{Message: errstring})
			},
			assertTest: func(t *testing.T, m *domain.Cart, err *errs.AppError) {
				assert.Nil(t, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockCartRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewCartService(mockDynamoRepo)

			data, err := service.GetCartByUserId("randomUserId")

			testcase.assertTest(t, data, err)
		})
	}

}

func TestAllGetCartRecordByuserId(t *testing.T) {
	itemList := map[string]domain.Item{
		"1": domain.Item{Name: "Pen", Cost: 10, Quantity: 1},
		"2": domain.Item{Name: "Pencil", Cost: 5, Quantity: 2},
	}
	cart := domain.Cart{
		//Id:primitive.NewObjectID().Hex(),
		UserID: "randomUserId",
		Items:  itemList,
	}

	carts:=[]domain.Cart{cart}
	

	testcases := []struct {
		name       string
		createStub func(mocks.MockCartRepository)
		assertTest func(*testing.T, []domain.Cart, *errs.AppError)
	}{
		{
			name: "SuccessGetAllCartRecordByUserId",
			createStub: func(mpdr mocks.MockCartRepository) {
				mpdr.EXPECT().FindAllCarts().Return(carts, nil)
			},
			assertTest: func(t *testing.T, m []domain.Cart, err *errs.AppError) {
				assert.NotNil(t, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailGetAllCartRecordByUserId",
			createStub: func(mpdr mocks.MockCartRepository) {
				errstring := "unable to find record"
				mpdr.EXPECT().
				FindAllCarts().
					Return(nil, &errs.AppError{Message: errstring})
			},
			assertTest: func(t *testing.T, m []domain.Cart, err *errs.AppError) {
				assert.Nil(t, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockCartRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewCartService(mockDynamoRepo)

			data, err := service.GetAllCarts()

			testcase.assertTest(t, data, err)
		})
	}

}

func TestUpdateCart(t *testing.T) {
	itemList := map[string]domain.Item{
		"1": domain.Item{Name: "Pen", Cost: 10, Quantity: 1},
		"2": domain.Item{Name: "Pencil", Cost: 5, Quantity: 2},
	}

	testcases := []struct {
		name       string
		createStub func(mocks.MockCartRepository)
		assertTest func(*testing.T, bool, *errs.AppError)
	}{
		{
			name: "SuccessUpdateCart",
			createStub: func(mpdr mocks.MockCartRepository) {
				mpdr.EXPECT().UpdateCartByUserId("randomUserId", itemList).Return(true, nil)
			},
			assertTest: func(t *testing.T, m bool, err *errs.AppError) {
				assert.Equal(t, true, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailUpdateCart",
			createStub: func(mpdr mocks.MockCartRepository) {
				errstring := "unable to find record"
				mpdr.EXPECT().
					UpdateCartByUserId("randomUserId", itemList).
					Return(false, &errs.AppError{Message: errstring})
			},
			assertTest: func(t *testing.T, m bool, err *errs.AppError) {
				assert.Equal(t, false, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockCartRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewCartService(mockDynamoRepo)

			data, err := service.UpdateCartItemsByUserId("randomUserId", itemList)

			testcase.assertTest(t, data, err)
		})
	}

}

func TestDeleteCartRecordByuserId(t *testing.T) {

	testcases := []struct {
		name       string
		createStub func(mocks.MockCartRepository)
		assertTest func(*testing.T, bool, *errs.AppError)
	}{
		{
			name: "SuccessDeleteCartRecordByUserId",
			createStub: func(mpdr mocks.MockCartRepository) {
				mpdr.EXPECT().DeleteCartByUserId("randomUserId").Return(true, nil)
			},
			assertTest: func(t *testing.T, m bool, err *errs.AppError) {
				assert.NotNil(t, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailDeleteCartRecordByUserId",
			createStub: func(mpdr mocks.MockCartRepository) {
				errstring := "unable to find record"
				mpdr.EXPECT().
					DeleteCartByUserId("randomUserId").
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
			mockDynamoRepo := mocks.NewMockCartRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewCartService(mockDynamoRepo)

			data, err := service.DeleteCartByUserId("randomUserId")

			testcase.assertTest(t, data, err)
		})
	}

}

func TestDeleteCartItems(t *testing.T) {
	
	productList:= []string{"1","2"}

	testcases := []struct {
		name       string
		createStub func(mocks.MockCartRepository)
		assertTest func(*testing.T, bool, *errs.AppError)
	}{
		{
			name: "SuccessDeleteCartItems",
			createStub: func(mpdr mocks.MockCartRepository) {
				mpdr.EXPECT().DeleteCartItemByUserId("randomUserId", productList).Return(true, nil)
			},
			assertTest: func(t *testing.T, m bool, err *errs.AppError) {
				assert.Equal(t, true, m)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailDeleteCartItems",
			createStub: func(mpdr mocks.MockCartRepository) {
				errstring := "unable to find record"
				mpdr.EXPECT().
				DeleteCartItemByUserId("randomUserId", productList).
					Return(false, &errs.AppError{Message: errstring})
			},
			assertTest: func(t *testing.T, m bool, err *errs.AppError) {
				assert.Equal(t, false, m)
				assert.NotNil(t, err)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockDynamoRepo := mocks.NewMockCartRepository(mockCtrl)
			testcase.createStub(*mockDynamoRepo)

			service := domain.NewCartService(mockDynamoRepo)

			data, err := service.DeleteCartItemByUserId("randomUserId", productList)

			testcase.assertTest(t, data, err)
		})
	}

}