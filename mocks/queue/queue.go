// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=../../../mocks/queue/queue.go
//

// Package mock_queue is a generated GoMock package.
package mock_queue

import (
	context "context"
	reflect "reflect"

	queue "github.com/talon-one/talon-backend-assingment/internal/domain/queue"
	gomock "go.uber.org/mock/gomock"
)

// MockMessageHandler is a mock of MessageHandler interface.
type MockMessageHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMessageHandlerMockRecorder
	isgomock struct{}
}

// MockMessageHandlerMockRecorder is the mock recorder for MockMessageHandler.
type MockMessageHandlerMockRecorder struct {
	mock *MockMessageHandler
}

// NewMockMessageHandler creates a new mock instance.
func NewMockMessageHandler(ctrl *gomock.Controller) *MockMessageHandler {
	mock := &MockMessageHandler{ctrl: ctrl}
	mock.recorder = &MockMessageHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageHandler) EXPECT() *MockMessageHandlerMockRecorder {
	return m.recorder
}

// HandleWithRetry mocks base method.
func (m *MockMessageHandler) HandleWithRetry(arg0 []byte, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleWithRetry", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleWithRetry indicates an expected call of HandleWithRetry.
func (mr *MockMessageHandlerMockRecorder) HandleWithRetry(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleWithRetry", reflect.TypeOf((*MockMessageHandler)(nil).HandleWithRetry), arg0, arg1)
}

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockInterface)(nil).Close))
}

// Consume mocks base method.
func (m *MockInterface) Consume(handler queue.MessageHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// Consume indicates an expected call of Consume.
func (mr *MockInterfaceMockRecorder) Consume(handler any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockInterface)(nil).Consume), handler)
}

// Produce mocks base method.
func (m *MockInterface) Produce(arg0 context.Context, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Produce indicates an expected call of Produce.
func (mr *MockInterfaceMockRecorder) Produce(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockInterface)(nil).Produce), arg0, arg1)
}
