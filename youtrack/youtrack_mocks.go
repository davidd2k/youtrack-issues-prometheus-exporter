// Code generated by MockGen. DO NOT EDIT.
// Source: youtrack.go

// Package youtrack is a generated GoMock package.
package youtrack

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockmakeRequester is a mock of makeRequester interface
type MockmakeRequester struct {
	ctrl     *gomock.Controller
	recorder *MockmakeRequesterMockRecorder
}

// MockmakeRequesterMockRecorder is the mock recorder for MockmakeRequester
type MockmakeRequesterMockRecorder struct {
	mock *MockmakeRequester
}

// NewMockmakeRequester creates a new mock instance
func NewMockmakeRequester(ctrl *gomock.Controller) *MockmakeRequester {
	mock := &MockmakeRequester{ctrl: ctrl}
	mock.recorder = &MockmakeRequesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockmakeRequester) EXPECT() *MockmakeRequesterMockRecorder {
	return m.recorder
}

// MakeRequest mocks base method
func (m *MockmakeRequester) MakeRequest(url string, headers map[string]string) ([]byte, error) {
	ret := m.ctrl.Call(m, "MakeRequest", url, headers)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeRequest indicates an expected call of MakeRequest
func (mr *MockmakeRequesterMockRecorder) MakeRequest(url, headers interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRequest", reflect.TypeOf((*MockmakeRequester)(nil).MakeRequest), url, headers)
}
