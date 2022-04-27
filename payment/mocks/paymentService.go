// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain (interfaces: PaymentService)

// Package mock_domain is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
)

// MockPaymentService is a mock of PaymentService interface.
type MockPaymentService struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentServiceMockRecorder
}

// MockPaymentServiceMockRecorder is the mock recorder for MockPaymentService.
type MockPaymentServiceMockRecorder struct {
	mock *MockPaymentService
}

// NewMockPaymentService creates a new mock instance.
func NewMockPaymentService(ctrl *gomock.Controller) *MockPaymentService {
	mock := &MockPaymentService{ctrl: ctrl}
	mock.recorder = &MockPaymentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentService) EXPECT() *MockPaymentServiceMockRecorder {
	return m.recorder
}

// AddPaymentMethod mocks base method.
func (m *MockPaymentService) AddPaymentMethod(arg0, arg1, arg2, arg3 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPaymentMethod", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPaymentMethod indicates an expected call of AddPaymentMethod.
func (mr *MockPaymentServiceMockRecorder) AddPaymentMethod(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPaymentMethod", reflect.TypeOf((*MockPaymentService)(nil).AddPaymentMethod), arg0, arg1, arg2, arg3)
}

// CreateDynamoPaymentRecord mocks base method.
func (m *MockPaymentService) CreateDynamoPaymentRecord(arg0 string, arg1 int16, arg2, arg3, arg4, arg5, arg6, arg7, arg8 string, arg9 []string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDynamoPaymentRecord", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDynamoPaymentRecord indicates an expected call of CreateDynamoPaymentRecord.
func (mr *MockPaymentServiceMockRecorder) CreateDynamoPaymentRecord(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDynamoPaymentRecord", reflect.TypeOf((*MockPaymentService)(nil).CreateDynamoPaymentRecord), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// GetPaymentMethods mocks base method.
func (m *MockPaymentService) GetPaymentMethods(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentMethods", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentMethods indicates an expected call of GetPaymentMethods.
func (mr *MockPaymentServiceMockRecorder) GetPaymentMethods(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentMethods", reflect.TypeOf((*MockPaymentService)(nil).GetPaymentMethods), arg0)
}

// GetPaymentRecordById mocks base method.
func (m *MockPaymentService) GetPaymentRecordById(arg0 string) (*domain.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentRecordById", arg0)
	ret0, _ := ret[0].(*domain.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentRecordById indicates an expected call of GetPaymentRecordById.
func (mr *MockPaymentServiceMockRecorder) GetPaymentRecordById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentRecordById", reflect.TypeOf((*MockPaymentService)(nil).GetPaymentRecordById), arg0)
}

// GetRazorpayPaymentLink mocks base method.
func (m *MockPaymentService) GetRazorpayPaymentLink(arg0 domain.Payment) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRazorpayPaymentLink", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRazorpayPaymentLink indicates an expected call of GetRazorpayPaymentLink.
func (mr *MockPaymentServiceMockRecorder) GetRazorpayPaymentLink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRazorpayPaymentLink", reflect.TypeOf((*MockPaymentService)(nil).GetRazorpayPaymentLink), arg0)
}

// UpdatePaymentStatus mocks base method.
func (m *MockPaymentService) UpdatePaymentStatus(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePaymentStatus", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePaymentStatus indicates an expected call of UpdatePaymentStatus.
func (mr *MockPaymentServiceMockRecorder) UpdatePaymentStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePaymentStatus", reflect.TypeOf((*MockPaymentService)(nil).UpdatePaymentStatus), arg0, arg1)
}