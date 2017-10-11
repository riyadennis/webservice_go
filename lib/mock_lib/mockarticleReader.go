// Code generated by MockGen. DO NOT EDIT.
// Source: articleReader.go

// Package mock_lib is a generated GoMock package.
package mock_lib

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockReader is a mock of Reader interface
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockReader) Read() {
	m.ctrl.Call(m, "Read")
}

// Read indicates an expected call of Read
func (mr *MockReaderMockRecorder) Read() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReader)(nil).Read))
}