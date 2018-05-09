// Code generated by MockGen. DO NOT EDIT.
// Source: gerrit.pb.go

// Package gerrit is a generated GoMock package.
package gerrit

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockGerritClient is a mock of GerritClient interface
type MockGerritClient struct {
	ctrl     *gomock.Controller
	recorder *MockGerritClientMockRecorder
}

// MockGerritClientMockRecorder is the mock recorder for MockGerritClient
type MockGerritClientMockRecorder struct {
	mock *MockGerritClient
}

// NewMockGerritClient creates a new mock instance
func NewMockGerritClient(ctrl *gomock.Controller) *MockGerritClient {
	mock := &MockGerritClient{ctrl: ctrl}
	mock.recorder = &MockGerritClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGerritClient) EXPECT() *MockGerritClientMockRecorder {
	return m.recorder
}

// GetChange mocks base method
func (m *MockGerritClient) GetChange(ctx context.Context, in *GetChangeRequest, opts ...grpc.CallOption) (*ChangeInfo, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChange", varargs...)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChange indicates an expected call of GetChange
func (mr *MockGerritClientMockRecorder) GetChange(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChange", reflect.TypeOf((*MockGerritClient)(nil).GetChange), varargs...)
}

// CheckAccess mocks base method
func (m *MockGerritClient) CheckAccess(ctx context.Context, in *CheckAccessRequest, opts ...grpc.CallOption) (*CheckAccessResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckAccess", varargs...)
	ret0, _ := ret[0].(*CheckAccessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAccess indicates an expected call of CheckAccess
func (mr *MockGerritClientMockRecorder) CheckAccess(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccess", reflect.TypeOf((*MockGerritClient)(nil).CheckAccess), varargs...)
}

// MockGerritServer is a mock of GerritServer interface
type MockGerritServer struct {
	ctrl     *gomock.Controller
	recorder *MockGerritServerMockRecorder
}

// MockGerritServerMockRecorder is the mock recorder for MockGerritServer
type MockGerritServerMockRecorder struct {
	mock *MockGerritServer
}

// NewMockGerritServer creates a new mock instance
func NewMockGerritServer(ctrl *gomock.Controller) *MockGerritServer {
	mock := &MockGerritServer{ctrl: ctrl}
	mock.recorder = &MockGerritServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGerritServer) EXPECT() *MockGerritServerMockRecorder {
	return m.recorder
}

// GetChange mocks base method
func (m *MockGerritServer) GetChange(arg0 context.Context, arg1 *GetChangeRequest) (*ChangeInfo, error) {
	ret := m.ctrl.Call(m, "GetChange", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChange indicates an expected call of GetChange
func (mr *MockGerritServerMockRecorder) GetChange(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChange", reflect.TypeOf((*MockGerritServer)(nil).GetChange), arg0, arg1)
}

// CheckAccess mocks base method
func (m *MockGerritServer) CheckAccess(arg0 context.Context, arg1 *CheckAccessRequest) (*CheckAccessResponse, error) {
	ret := m.ctrl.Call(m, "CheckAccess", arg0, arg1)
	ret0, _ := ret[0].(*CheckAccessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAccess indicates an expected call of CheckAccess
func (mr *MockGerritServerMockRecorder) CheckAccess(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccess", reflect.TypeOf((*MockGerritServer)(nil).CheckAccess), arg0, arg1)
}