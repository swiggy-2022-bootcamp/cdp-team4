// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain (interfaces: PaymentDynamoRepository)

// Package mock_domain is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
)

// MockPaymentDynamoRepository is a mock of PaymentDynamoRepository interface.
type MockPaymentDynamoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentDynamoRepositoryMockRecorder
}

// MockPaymentDynamoRepositoryMockRecorder is the mock recorder for MockPaymentDynamoRepository.
type MockPaymentDynamoRepositoryMockRecorder struct {
	mock *MockPaymentDynamoRepository
}

// NewMockPaymentDynamoRepository creates a new mock instance.
func NewMockPaymentDynamoRepository(ctrl *gomock.Controller) *MockPaymentDynamoRepository {
	mock := &MockPaymentDynamoRepository{ctrl: ctrl}
	mock.recorder = &MockPaymentDynamoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentDynamoRepository) EXPECT() *MockPaymentDynamoRepositoryMockRecorder {
	return m.recorder
}

// DeletePaymentRecordByID mocks base method.
func (m *MockPaymentDynamoRepository) DeletePaymentRecordByID(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePaymentRecordByID", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePaymentRecordByID indicates an expected call of DeletePaymentRecordByID.
func (mr *MockPaymentDynamoRepositoryMockRecorder) DeletePaymentRecordByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePaymentRecordByID", reflect.TypeOf((*MockPaymentDynamoRepository)(nil).DeletePaymentRecordByID), arg0)
}

// FindPaymentRecordById mocks base method.
func (m *MockPaymentDynamoRepository) FindPaymentRecordById(arg0 string) (*domain.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPaymentRecordById", arg0)
	ret0, _ := ret[0].(*domain.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPaymentRecordById indicates an expected call of FindPaymentRecordById.
func (mr *MockPaymentDynamoRepositoryMockRecorder) FindPaymentRecordById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPaymentRecordById", reflect.TypeOf((*MockPaymentDynamoRepository)(nil).FindPaymentRecordById), arg0)
}

// FindPaymentRecordByUserID mocks base method.
func (m *MockPaymentDynamoRepository) FindPaymentRecordByUserID(arg0 string) ([]domain.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPaymentRecordByUserID", arg0)
	ret0, _ := ret[0].([]domain.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPaymentRecordByUserID indicates an expected call of FindPaymentRecordByUserID.
func (mr *MockPaymentDynamoRepositoryMockRecorder) FindPaymentRecordByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPaymentRecordByUserID", reflect.TypeOf((*MockPaymentDynamoRepository)(nil).FindPaymentRecordByUserID), arg0)
}

// GetPaymentMethods mocks base method.
func (m *MockPaymentDynamoRepository) GetPaymentMethods(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentMethods", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentMethods indicates an expected call of GetPaymentMethods.
func (mr *MockPaymentDynamoRepositoryMockRecorder) GetPaymentMethods(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentMethods", reflect.TypeOf((*MockPaymentDynamoRepository)(nil).GetPaymentMethods), arg0)
}

// InsertPaymentMethod mocks base method.
func (m *MockPaymentDynamoRepository) InsertPaymentMethod(arg0 domain.PaymentMethod) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPaymentMethod", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertPaymentMethod indicates an expected call of InsertPaymentMethod.
func (mr *MockPaymentDynamoRepositoryMockRecorder) InsertPaymentMethod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPaymentMethod", reflect.TypeOf((*MockPaymentDynamoRepository)(nil).InsertPaymentMethod), arg0)
}

// InsertPaymentRecord mocks base method.
func (m *MockPaymentDynamoRepository) InsertPaymentRecord(arg0 domain.Payment) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPaymentRecord", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertPaymentRecord indicates an expected call of InsertPaymentRecord.
func (mr *MockPaymentDynamoRepositoryMockRecorder) InsertPaymentRecord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPaymentRecord", reflect.TypeOf((*MockPaymentDynamoRepository)(nil).InsertPaymentRecord), arg0)
}

// UpdatePaymentMethods mocks base method.
func (m *MockPaymentDynamoRepository) UpdatePaymentMethods(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePaymentMethods", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePaymentMethods indicates an expected call of UpdatePaymentMethods.
func (mr *MockPaymentDynamoRepositoryMockRecorder) UpdatePaymentMethods(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePaymentMethods", reflect.TypeOf((*MockPaymentDynamoRepository)(nil).UpdatePaymentMethods), arg0, arg1)
}

// UpdatePaymentRecord mocks base method.
func (m *MockPaymentDynamoRepository) UpdatePaymentRecord(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePaymentRecord", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePaymentRecord indicates an expected call of UpdatePaymentRecord.
func (mr *MockPaymentDynamoRepositoryMockRecorder) UpdatePaymentRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePaymentRecord", reflect.TypeOf((*MockPaymentDynamoRepository)(nil).UpdatePaymentRecord), arg0, arg1)
}