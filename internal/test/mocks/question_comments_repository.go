// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/forum/application/repositories/question_comments_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	enterprise "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

// MockQuestionCommentsRepositoryInterface is a mock of QuestionCommentsRepositoryInterface interface.
type MockQuestionCommentsRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockQuestionCommentsRepositoryInterfaceMockRecorder
}

// MockQuestionCommentsRepositoryInterfaceMockRecorder is the mock recorder for MockQuestionCommentsRepositoryInterface.
type MockQuestionCommentsRepositoryInterfaceMockRecorder struct {
	mock *MockQuestionCommentsRepositoryInterface
}

// NewMockQuestionCommentsRepositoryInterface creates a new mock instance.
func NewMockQuestionCommentsRepositoryInterface(ctrl *gomock.Controller) *MockQuestionCommentsRepositoryInterface {
	mock := &MockQuestionCommentsRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockQuestionCommentsRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuestionCommentsRepositoryInterface) EXPECT() *MockQuestionCommentsRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockQuestionCommentsRepositoryInterface) Create(questionComment *enterprise.QuestionComment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", questionComment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockQuestionCommentsRepositoryInterfaceMockRecorder) Create(questionComment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockQuestionCommentsRepositoryInterface)(nil).Create), questionComment)
}

// DeleteByID mocks base method.
func (m *MockQuestionCommentsRepositoryInterface) DeleteByID(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockQuestionCommentsRepositoryInterfaceMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockQuestionCommentsRepositoryInterface)(nil).DeleteByID), id)
}

// GetByID mocks base method.
func (m *MockQuestionCommentsRepositoryInterface) GetByID(id string) (enterprise.QuestionComment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(enterprise.QuestionComment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockQuestionCommentsRepositoryInterfaceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockQuestionCommentsRepositoryInterface)(nil).GetByID), id)
}
