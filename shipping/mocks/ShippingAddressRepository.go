// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
	errs "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"

	mock "github.com/stretchr/testify/mock"
)

// ShippingAddressRepository is an autogenerated mock type for the ShippingAddressRepository type
type ShippingAddressRepository struct {
	mock.Mock
}

// DeleteShippingAddressById provides a mock function with given fields: _a0
func (_m *ShippingAddressRepository) DeleteShippingAddressById(_a0 int) *errs.AppError {
	ret := _m.Called(_a0)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(int) *errs.AppError); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// FindShippingAddressById provides a mock function with given fields: _a0
func (_m *ShippingAddressRepository) FindShippingAddressById(_a0 int) (*domain.ShippingAddress, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.ShippingAddress
	if rf, ok := ret.Get(0).(func(int) *domain.ShippingAddress); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ShippingAddress)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// InsertShippingAddress provides a mock function with given fields: _a0
func (_m *ShippingAddressRepository) InsertShippingAddress(_a0 domain.ShippingAddress) (domain.ShippingAddress, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 domain.ShippingAddress
	if rf, ok := ret.Get(0).(func(domain.ShippingAddress) domain.ShippingAddress); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.ShippingAddress)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(domain.ShippingAddress) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// UpdateShippingAddress provides a mock function with given fields: _a0
func (_m *ShippingAddressRepository) UpdateShippingAddress(_a0 domain.ShippingAddress) (*domain.ShippingAddress, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.ShippingAddress
	if rf, ok := ret.Get(0).(func(domain.ShippingAddress) *domain.ShippingAddress); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ShippingAddress)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(domain.ShippingAddress) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}