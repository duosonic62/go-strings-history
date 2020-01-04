// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/repository/user_command.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	entity "github.com/duosonic62/go-strings-history/internal/domain/entity"
	valueobject "github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserCommandRepository is a mock of UserCommandRepository interface
type MockUserCommandRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserCommandRepositoryMockRecorder
}

// MockUserCommandRepositoryMockRecorder is the mock recorder for MockUserCommandRepository
type MockUserCommandRepositoryMockRecorder struct {
	mock *MockUserCommandRepository
}

// NewMockUserCommandRepository creates a new mock instance
func NewMockUserCommandRepository(ctrl *gomock.Controller) *MockUserCommandRepository {
	mock := &MockUserCommandRepository{ctrl: ctrl}
	mock.recorder = &MockUserCommandRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserCommandRepository) EXPECT() *MockUserCommandRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockUserCommandRepository) Find(arg0 valueobject.AuthorizationToken) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockUserCommandRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserCommandRepository)(nil).Find), arg0)
}

// Save mocks base method
func (m *MockUserCommandRepository) Save(arg0 *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockUserCommandRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserCommandRepository)(nil).Save), arg0)
}

// Edit mocks base method
func (m *MockUserCommandRepository) Edit(user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit
func (mr *MockUserCommandRepositoryMockRecorder) Edit(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockUserCommandRepository)(nil).Edit), user)
}

// Delete mocks base method
func (m *MockUserCommandRepository) Delete(token valueobject.AuthorizationToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockUserCommandRepositoryMockRecorder) Delete(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserCommandRepository)(nil).Delete), token)
}
