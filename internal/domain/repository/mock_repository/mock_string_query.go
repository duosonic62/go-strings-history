// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/repository/string_query.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	output "github.com/duosonic62/go-strings-history/pkg/usecase/output"
	gomock "github.com/golang/mock/gomock"
	null "github.com/volatiletech/null"
	reflect "reflect"
)

// MockStringQueryRepository is a mock of StringQueryRepository interface
type MockStringQueryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStringQueryRepositoryMockRecorder
}

// MockStringQueryRepositoryMockRecorder is the mock recorder for MockStringQueryRepository
type MockStringQueryRepositoryMockRecorder struct {
	mock *MockStringQueryRepository
}

// NewMockStringQueryRepository creates a new mock instance
func NewMockStringQueryRepository(ctrl *gomock.Controller) *MockStringQueryRepository {
	mock := &MockStringQueryRepository{ctrl: ctrl}
	mock.recorder = &MockStringQueryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStringQueryRepository) EXPECT() *MockStringQueryRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockStringQueryRepository) Find(stringID string) (*output.GuitarStringOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", stringID)
	ret0, _ := ret[0].(*output.GuitarStringOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockStringQueryRepositoryMockRecorder) Find(stringID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockStringQueryRepository)(nil).Find), stringID)
}

// Search mocks base method
func (m *MockStringQueryRepository) Search(name, maker null.String, thinGauge, thickGauge null.Int) (*[]output.GuitarStringOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", name, maker, thinGauge, thickGauge)
	ret0, _ := ret[0].(*[]output.GuitarStringOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockStringQueryRepositoryMockRecorder) Search(name, maker, thinGauge, thickGauge interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockStringQueryRepository)(nil).Search), name, maker, thinGauge, thickGauge)
}
