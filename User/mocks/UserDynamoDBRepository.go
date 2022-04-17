// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"
)

// UserDynamoDBRepository is an autogenerated mock type for the UserDynamoDBRepository type
type UserDynamoDBRepository struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: _a0
func (_m *UserDynamoDBRepository) DeleteByID(_a0 string) (bool, error) {
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

// FindAll provides a mock function with given fields:
func (_m *UserDynamoDBRepository) FindAll() ([]domain.User, error) {
	ret := _m.Called()

	var r0 []domain.User
	if rf, ok := ret.Get(0).(func() []domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
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

// FindByID provides a mock function with given fields: _a0
func (_m *UserDynamoDBRepository) FindByID(_a0 string) (*domain.User, error) {
	ret := _m.Called(_a0)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
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

// Save provides a mock function with given fields: _a0
func (_m *UserDynamoDBRepository) Save(_a0 domain.User) (domain.User, error) {
	ret := _m.Called(_a0)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateById provides a mock function with given fields: _a0
func (_m *UserDynamoDBRepository) UpdateById(_a0 domain.User) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.User) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
