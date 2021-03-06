// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/utils/httputils/client.go

package httputils

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockClient) EXPECT() *MockClientMockRecorder {
	return _m.recorder
}

// Get mocks base method
func (_m *MockClient) Get(ctx context.Context, url string) (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "Get", ctx, url)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (_mr *MockClientMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), arg0, arg1)
}
