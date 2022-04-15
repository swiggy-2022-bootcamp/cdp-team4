// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/swiggy-2022-bootcamp/cdp-team4/payment/domain"
)

// PaymentDynamoRepository is an autogenerated mock type for the PaymentDynamoRepository type
type PaymentDynamoRepository struct {
	mock.Mock
}

// DeletePaymentRecordByID provides a mock function with given fields: _a0
func (_m *PaymentDynamoRepository) DeletePaymentRecordByID(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPaymentRecordById provides a mock function with given fields: _a0
func (_m *PaymentDynamoRepository) FindPaymentRecordById(_a0 string) (*domain.Payment, error) {
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

// FindPaymentRecordByUserID provides a mock function with given fields: _a0
func (_m *PaymentDynamoRepository) FindPaymentRecordByUserID(_a0 string) ([]domain.Payment, error) {
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
func (_m *PaymentDynamoRepository) GetPaymentMethods(_a0 string) ([]string, error) {
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

// InsertPaymentMethod provides a mock function with given fields: _a0
func (_m *PaymentDynamoRepository) InsertPaymentMethod(_a0 domain.PaymentMethod) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.PaymentMethod) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.PaymentMethod) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertPaymentRecord provides a mock function with given fields: _a0
func (_m *PaymentDynamoRepository) InsertPaymentRecord(_a0 domain.Payment) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.Payment) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Payment) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePaymentMethods provides a mock function with given fields: _a0, _a1
func (_m *PaymentDynamoRepository) UpdatePaymentMethods(_a0 string, _a1 string) (bool, error) {
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

// UpdatePaymentRecord provides a mock function with given fields: _a0, _a1
func (_m *PaymentDynamoRepository) UpdatePaymentRecord(_a0 string, _a1 string) (bool, error) {
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
