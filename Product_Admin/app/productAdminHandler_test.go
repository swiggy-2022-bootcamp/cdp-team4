package app_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/app"
// 	"github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/domain"
// 	mocks "github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/mocks"
// )

// func TestHandleGetProductByID(t *testing.T) {

// 	testcases := []struct {
// 		name       string
// 		createStub func(*mocks.MockProductAdminService)
// 		expected   int
// 	}{
// 		{
// 			name: "SuccessGetProductRecordByID",
// 			createStub: func(mps *mocks.MockProductAdminService) {
// 				mps.EXPECT().
// 					GetProductById(gomock.Any() /* id */).
// 					Return(domain.Product{}, nil)
// 			},
// 			expected: 200,
// 		},
// 		// {
// 		// 	name: "FailGetProductRecordByID",
// 		// 	createStub: func(mps *mocks.MockProductAdminService) {
// 		// 		mps.EXPECT().
// 		// 			GetProductById("xyx" /* id */).
// 		// 			Return(domain.Product{}, fmt.Errorf("record not found"))
// 		// 	},
// 		// 	expected: 400,
// 		// },
// 	}

// 	for _, testcase := range testcases {
// 		t.Run(testcase.name, func(t *testing.T) {
// 			mockCtrl := gomock.NewController(t)
// 			defer mockCtrl.Finish()

// 			mockService := mocks.NewMockProductAdminService(mockCtrl)
// 			testcase.createStub(mockService)

// 			router := app.SetupRouter(app.ProductAdminHandler{
// 				ProductAdminService: mockService,
// 			})

// 			recorder := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodGet, "/products/xyx", nil)
// 			router.ServeHTTP(recorder, req)

// 			assert.Equal(t, testcase.expected, recorder.Code)
// 		})
// 	}
// }
