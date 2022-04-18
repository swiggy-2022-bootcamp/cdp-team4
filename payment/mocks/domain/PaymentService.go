// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
)

// PaymentService is an autogenerated mock type for the PaymentService type
type PaymentService struct {
	mock.Mock
}

// AddPaymentMethod provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *PaymentService) AddPaymentMethod(_a0 string, _a1 string, _a2 string, _a3 string) (bool, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string, string) bool); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDynamoPaymentRecord provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8
func (_m *PaymentService) CreateDynamoPaymentRecord(_a0 int16, _a1 string, _a2 string, _a3 string, _a4 string, _a5 string, _a6 string, _a7 string, _a8 []string) (bool, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int16, string, string, string, string, string, string, string, []string) bool); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int16, string, string, string, string, string, string, string, []string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentAllRecordsByUserId provides a mock function with given fields: _a0
func (_m *PaymentService) GetPaymentAllRecordsByUserId(_a0 string) ([]domain.Payment, error) {
	ret := _m.Called(_a0)

	var r0 []domain.Payment
	if rf, ok := ret.Get(0).(func(string) []domain.Payment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentMethods provides a mock function with given fields: _a0
func (_m *PaymentService) GetPaymentMethods(_a0 string) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentRecordById provides a mock function with given fields: _a0
func (_m *PaymentService) GetPaymentRecordById(_a0 string) (*domain.Payment, error) {
	ret := _m.Called(_a0)

	var r0 *domain.Payment
	if rf, ok := ret.Get(0).(func(string) *domain.Payment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePaymentMethod provides a mock function with given fields: _a0, _a1
func (_m *PaymentService) UpdatePaymentMethod(_a0 string, _a1 string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePaymentStatus provides a mock function with given fields: _a0, _a1
func (_m *PaymentService) UpdatePaymentStatus(_a0 string, _a1 string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}