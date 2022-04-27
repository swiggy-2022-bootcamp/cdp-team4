package app_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/app"
// 	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/domain"
// 	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/Product_FrontStore/mocks"
// )

// func TestHandleGetProductRecordByID(t *testing.T) {

// 	testcases := []struct {
// 		name       string
// 		createStub func(*mocks.MockProductFrontStoreService)
// 		expected   int
// 	}{
// 		{
// 			name: "SuccessGetProductRecordByID",
// 			createStub: func(mps *mocks.MockProductFrontStoreService) {
// 				mps.EXPECT().GetProductById("xyx" /* id */).Return(domain.Product{}, nil)
// 			},
// 			expected: 200,
// 		},
// 		// {
// 		// 	name: "FailGetProductRecordByID",
// 		// 	createStub: func(mps *mocks.MockProductFrontStoreService) {
// 		// 		mps.EXPECT().GetProductById("xyx" /* id */).Return(nil, &errs.AppError{Message: "errstring"})
// 		// 	},
// 		// 	expected: 500,
// 		// },
// 	}

// 	for _, testcase := range testcases {
// 		t.Run(testcase.name, func(t *testing.T) {
// 			mockCtrl := gomock.NewController(t)
// 			defer mockCtrl.Finish()

// 			mockService := mocks.NewMockProductFrontStoreService(mockCtrl)
// 			testcase.createStub(mockService)

// 			router := app.SetupRouter(app.ProductFrontStoreHandler{
// 				ProductFrontStoreService: mockService,
// 			})

// 			recorder := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodGet, "/products/xyx", nil)
// 			router.ServeHTTP(recorder, req)

// 			assert.Equal(t, testcase.expected, recorder.Code)
// 		})
// 	}
// }
