// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/analyze/fetchers/fetcher.go

package fetchers

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	executors "github.com/golangci/golangci-worker/app/analyze/executors"
)

// MockFetcher is a mock of Fetcher interface
type MockFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockFetcherMockRecorder
}

// MockFetcherMockRecorder is the mock recorder for MockFetcher
type MockFetcherMockRecorder struct {
	mock *MockFetcher
}

// NewMockFetcher creates a new mock instance
func NewMockFetcher(ctrl *gomock.Controller) *MockFetcher {
	mock := &MockFetcher{ctrl: ctrl}
	mock.recorder = &MockFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockFetcher) EXPECT() *MockFetcherMockRecorder {
	return _m.recorder
}

// Fetch mocks base method
func (_m *MockFetcher) Fetch(ctx context.Context, url string, ref string, destDir string, exec executors.Executor) error {
	ret := _m.ctrl.Call(_m, "Fetch", ctx, url, ref, destDir, exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fetch indicates an expected call of Fetch
func (_mr *MockFetcherMockRecorder) Fetch(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Fetch", reflect.TypeOf((*MockFetcher)(nil).Fetch), arg0, arg1, arg2, arg3, arg4)
}
