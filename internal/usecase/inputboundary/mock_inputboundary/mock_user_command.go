// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecase/inputboundary/user_command.go

// Package mock_inputboundary is a generated GoMock package.
package mock_inputboundary

import (
	input "github.com/duosonic62/go-strings-history/pkg/usecase/input"
	command "github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserCommandUseCase is a mock of UserCommandUseCase interface
type MockUserCommandUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserCommandUseCaseMockRecorder
}

// MockUserCommandUseCaseMockRecorder is the mock recorder for MockUserCommandUseCase
type MockUserCommandUseCaseMockRecorder struct {
	mock *MockUserCommandUseCase
}

// NewMockUserCommandUseCase creates a new mock instance
func NewMockUserCommandUseCase(ctrl *gomock.Controller) *MockUserCommandUseCase {
	mock := &MockUserCommandUseCase{ctrl: ctrl}
	mock.recorder = &MockUserCommandUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserCommandUseCase) EXPECT() *MockUserCommandUseCaseMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockUserCommandUseCase) Add(data command.UserAddInputData, ctx input.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", data, ctx)
}

// Add indicates an expected call of Add
func (mr *MockUserCommandUseCaseMockRecorder) Add(data, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockUserCommandUseCase)(nil).Add), data, ctx)
}

// Edit mocks base method
func (m *MockUserCommandUseCase) Edit(data command.UserEditInputData, ctx input.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Edit", data, ctx)
}

// Edit indicates an expected call of Edit
func (mr *MockUserCommandUseCaseMockRecorder) Edit(data, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockUserCommandUseCase)(nil).Edit), data, ctx)
}
