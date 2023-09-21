// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/forum/application/repositories/answer_comments_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	enterprise "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

// MockAnswerCommentsRepositoryInterface is a mock of AnswerCommentsRepositoryInterface interface.
type MockAnswerCommentsRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAnswerCommentsRepositoryInterfaceMockRecorder
}

// MockAnswerCommentsRepositoryInterfaceMockRecorder is the mock recorder for MockAnswerCommentsRepositoryInterface.
type MockAnswerCommentsRepositoryInterfaceMockRecorder struct {
	mock *MockAnswerCommentsRepositoryInterface
}

// NewMockAnswerCommentsRepositoryInterface creates a new mock instance.
func NewMockAnswerCommentsRepositoryInterface(ctrl *gomock.Controller) *MockAnswerCommentsRepositoryInterface {
	mock := &MockAnswerCommentsRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockAnswerCommentsRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnswerCommentsRepositoryInterface) EXPECT() *MockAnswerCommentsRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAnswerCommentsRepositoryInterface) Create(answerComment *enterprise.AnswerComment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", answerComment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAnswerCommentsRepositoryInterfaceMockRecorder) Create(answerComment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAnswerCommentsRepositoryInterface)(nil).Create), answerComment)
}

// DeleteByID mocks base method.
func (m *MockAnswerCommentsRepositoryInterface) DeleteByID(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockAnswerCommentsRepositoryInterfaceMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockAnswerCommentsRepositoryInterface)(nil).DeleteByID), id)
}

// GetByID mocks base method.
func (m *MockAnswerCommentsRepositoryInterface) GetByID(id string) (enterprise.AnswerComment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(enterprise.AnswerComment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockAnswerCommentsRepositoryInterfaceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAnswerCommentsRepositoryInterface)(nil).GetByID), id)
}

// GetManyByID mocks base method.
func (m *MockAnswerCommentsRepositoryInterface) GetManyByID(page int64, id string) ([]enterprise.AnswerComment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManyByID", page, id)
	ret0, _ := ret[0].([]enterprise.AnswerComment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManyByID indicates an expected call of GetManyByID.
func (mr *MockAnswerCommentsRepositoryInterfaceMockRecorder) GetManyByID(page, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManyByID", reflect.TypeOf((*MockAnswerCommentsRepositoryInterface)(nil).GetManyByID), page, id)
}