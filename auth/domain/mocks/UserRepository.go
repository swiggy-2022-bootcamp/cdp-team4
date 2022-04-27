// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/swiggy-2022-bootcamp/cdp-team4/auth/domain"
	errs "github.com/swiggy-2022-bootcamp/cdp-team4/auth/utils/errs"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// FindByUsername provides a mock function with given fields: _a0
func (_m *UserRepository) FindByUsername(_a0 string) (*domain.UserModel, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.UserModel
	if rf, ok := ret.Get(0).(func(string) *domain.UserModel); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserModel)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t testing.TB) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
