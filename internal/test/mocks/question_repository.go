// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/forum/application/repositories/question_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	enterprise "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

// MockQuestionRepositoryInterface is a mock of QuestionRepositoryInterface interface.
type MockQuestionRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockQuestionRepositoryInterfaceMockRecorder
}

// MockQuestionRepositoryInterfaceMockRecorder is the mock recorder for MockQuestionRepositoryInterface.
type MockQuestionRepositoryInterfaceMockRecorder struct {
	mock *MockQuestionRepositoryInterface
}

// NewMockQuestionRepositoryInterface creates a new mock instance.
func NewMockQuestionRepositoryInterface(ctrl *gomock.Controller) *MockQuestionRepositoryInterface {
	mock := &MockQuestionRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockQuestionRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuestionRepositoryInterface) EXPECT() *MockQuestionRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockQuestionRepositoryInterface) Create(question *enterprise.Question) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", question)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockQuestionRepositoryInterfaceMockRecorder) Create(question interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockQuestionRepositoryInterface)(nil).Create), question)
}

// DeleteByID mocks base method.
func (m *MockQuestionRepositoryInterface) DeleteByID(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockQuestionRepositoryInterfaceMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockQuestionRepositoryInterface)(nil).DeleteByID), id)
}

// GetByID mocks base method.
func (m *MockQuestionRepositoryInterface) GetByID(id string) (enterprise.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(enterprise.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockQuestionRepositoryInterfaceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockQuestionRepositoryInterface)(nil).GetByID), id)
}

// GetBySlug mocks base method.
func (m *MockQuestionRepositoryInterface) GetBySlug(slug string) (enterprise.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySlug", slug)
	ret0, _ := ret[0].(enterprise.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySlug indicates an expected call of GetBySlug.
func (mr *MockQuestionRepositoryInterfaceMockRecorder) GetBySlug(slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySlug", reflect.TypeOf((*MockQuestionRepositoryInterface)(nil).GetBySlug), slug)
}

// GetManyRecent mocks base method.
func (m *MockQuestionRepositoryInterface) GetManyRecent(page int64) ([]enterprise.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManyRecent", page)
	ret0, _ := ret[0].([]enterprise.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManyRecent indicates an expected call of GetManyRecent.
func (mr *MockQuestionRepositoryInterfaceMockRecorder) GetManyRecent(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManyRecent", reflect.TypeOf((*MockQuestionRepositoryInterface)(nil).GetManyRecent), page)
}

// Save mocks base method.
func (m *MockQuestionRepositoryInterface) Save(question *enterprise.Question) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", question)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockQuestionRepositoryInterfaceMockRecorder) Save(question interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockQuestionRepositoryInterface)(nil).Save), question)
}
