// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\DELL\Desktop\final_project_S\cdp-team4\user\domain\userService.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/swiggy-2022-bootcamp/cdp-team4/user/domain"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUserInDynamodb mocks base method.
func (m *MockUserService) CreateUserInDynamodb(arg0, arg1, arg2, arg3, arg4, arg5 string, arg6 domain.Role, arg7, arg8 string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserInDynamodb", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserInDynamodb indicates an expected call of CreateUserInDynamodb.
func (mr *MockUserServiceMockRecorder) CreateUserInDynamodb(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserInDynamodb", reflect.TypeOf((*MockUserService)(nil).CreateUserInDynamodb), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// DeleteUserById mocks base method.
func (m *MockUserService) DeleteUserById(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserById", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserById indicates an expected call of DeleteUserById.
func (mr *MockUserServiceMockRecorder) DeleteUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockUserService)(nil).DeleteUserById), arg0)
}

// GetAllUsers mocks base method.
func (m *MockUserService) GetAllUsers() ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserServiceMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserService)(nil).GetAllUsers))
}

// GetUserById mocks base method.
func (m *MockUserService) GetUserById(arg0 string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserServiceMockRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserService)(nil).GetUserById), arg0)
}

// UpdateUserById mocks base method.
func (m *MockUserService) UpdateUserById(user_id, firstName, lastName, username, phone, email, password string, role domain.Role, addressId, fax string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserById", user_id, firstName, lastName, username, phone, email, password, role, addressId, fax)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserById indicates an expected call of UpdateUserById.
func (mr *MockUserServiceMockRecorder) UpdateUserById(user_id, firstName, lastName, username, phone, email, password, role, addressId, fax interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserById", reflect.TypeOf((*MockUserService)(nil).UpdateUserById), user_id, firstName, lastName, username, phone, email, password, role, addressId, fax)
}
