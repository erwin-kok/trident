// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/internal/nodeprep/nodeinfo (interfaces: Binary)
//
// Generated by this command:
//
//	mockgen -destination=../../../mocks/mock_internal/mock_nodeprep/mock_nodeinfo/mock_binary.go github.com/netapp/trident/internal/nodeprep/nodeinfo Binary
//

// Package mock_nodeinfo is a generated GoMock package.
package mock_nodeinfo

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBinary is a mock of Binary interface.
type MockBinary struct {
	ctrl     *gomock.Controller
	recorder *MockBinaryMockRecorder
}

// MockBinaryMockRecorder is the mock recorder for MockBinary.
type MockBinaryMockRecorder struct {
	mock *MockBinary
}

// NewMockBinary creates a new mock instance.
func NewMockBinary(ctrl *gomock.Controller) *MockBinary {
	mock := &MockBinary{ctrl: ctrl}
	mock.recorder = &MockBinaryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBinary) EXPECT() *MockBinaryMockRecorder {
	return m.recorder
}

// FindPath mocks base method.
func (m *MockBinary) FindPath(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPath", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FindPath indicates an expected call of FindPath.
func (mr *MockBinaryMockRecorder) FindPath(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPath", reflect.TypeOf((*MockBinary)(nil).FindPath), arg0)
}
