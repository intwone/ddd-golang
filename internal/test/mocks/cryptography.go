// Code generated by MockGen. DO NOT EDIT.
// Source: internal/infra/cryptography/cryptography.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCryptographyInterface is a mock of CryptographyInterface interface.
type MockCryptographyInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCryptographyInterfaceMockRecorder
}

// MockCryptographyInterfaceMockRecorder is the mock recorder for MockCryptographyInterface.
type MockCryptographyInterfaceMockRecorder struct {
	mock *MockCryptographyInterface
}

// NewMockCryptographyInterface creates a new mock instance.
func NewMockCryptographyInterface(ctrl *gomock.Controller) *MockCryptographyInterface {
	mock := &MockCryptographyInterface{ctrl: ctrl}
	mock.recorder = &MockCryptographyInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptographyInterface) EXPECT() *MockCryptographyInterfaceMockRecorder {
	return m.recorder
}

// Decrypt mocks base method.
func (m *MockCryptographyInterface) Decrypt(token string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", token)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockCryptographyInterfaceMockRecorder) Decrypt(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockCryptographyInterface)(nil).Decrypt), token)
}

// Encrypt mocks base method.
func (m *MockCryptographyInterface) Encrypt(value string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", value)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockCryptographyInterfaceMockRecorder) Encrypt(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockCryptographyInterface)(nil).Encrypt), value)
}
