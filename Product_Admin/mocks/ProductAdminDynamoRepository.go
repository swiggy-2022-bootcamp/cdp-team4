// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin/domain"
)

// ProductAdminDynamoRepository is an autogenerated mock type for the ProductAdminDynamoRepository type
type ProductAdminDynamoRepository struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: _a0
func (_m *ProductAdminDynamoRepository) DeleteByID(_a0 string) (bool, error) {
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

// Find provides a mock function with given fields:
func (_m *ProductAdminDynamoRepository) Find() ([]domain.Product, error) {
	ret := _m.Called()

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func() []domain.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByCategoryID provides a mock function with given fields: _a0
func (_m *ProductAdminDynamoRepository) FindByCategoryID(_a0 string) ([]domain.Product, error) {
	ret := _m.Called(_a0)

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func(string) []domain.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
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

// FindByID provides a mock function with given fields: _a0
func (_m *ProductAdminDynamoRepository) FindByID(_a0 string) (domain.Product, error) {
	ret := _m.Called(_a0)

	var r0 domain.Product
	if rf, ok := ret.Get(0).(func(string) domain.Product); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByKeyword provides a mock function with given fields: _a0
func (_m *ProductAdminDynamoRepository) FindByKeyword(_a0 string) ([]domain.Product, error) {
	ret := _m.Called(_a0)

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func(string) []domain.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
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

// FindByManufacturerID provides a mock function with given fields: _a0
func (_m *ProductAdminDynamoRepository) FindByManufacturerID(_a0 string) ([]domain.Product, error) {
	ret := _m.Called(_a0)

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func(string) []domain.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
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

// GetProductAvailability provides a mock function with given fields: _a0, _a1
func (_m *ProductAdminDynamoRepository) GetProductAvailability(_a0 string, _a1 int64) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, int64) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0
func (_m *ProductAdminDynamoRepository) Insert(_a0 domain.Product) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.Product) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Product) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateItem provides a mock function with given fields: _a0
func (_m *ProductAdminDynamoRepository) UpdateItem(_a0 domain.Product) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.Product) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Product) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateQuantity provides a mock function with given fields: _a0, _a1
func (_m *ProductAdminDynamoRepository) UpdateQuantity(_a0 string, _a1 int64) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, int64) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}