// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/replicatedhq/ship/pkg/lifecycle/render/helm (interfaces: Templater)

// Package helm is a generated GoMock package.
package helm

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	libyaml "github.com/replicatedhq/libyaml"
	api "github.com/replicatedhq/ship/pkg/api"
	root "github.com/replicatedhq/ship/pkg/lifecycle/render/root"
)

// MockTemplater is a mock of Templater interface
type MockTemplater struct {
	ctrl     *gomock.Controller
	recorder *MockTemplaterMockRecorder
}

// MockTemplaterMockRecorder is the mock recorder for MockTemplater
type MockTemplaterMockRecorder struct {
	mock *MockTemplater
}

// NewMockTemplater creates a new mock instance
func NewMockTemplater(ctrl *gomock.Controller) *MockTemplater {
	mock := &MockTemplater{ctrl: ctrl}
	mock.recorder = &MockTemplaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTemplater) EXPECT() *MockTemplaterMockRecorder {
	return m.recorder
}

// Template mocks base method
func (m *MockTemplater) Template(arg0 string, arg1 root.Fs, arg2 api.HelmAsset, arg3 api.ReleaseMetadata, arg4 []libyaml.ConfigGroup, arg5 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// Template indicates an expected call of Template
func (mr *MockTemplaterMockRecorder) Template(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockTemplater)(nil).Template), arg0, arg1, arg2, arg3, arg4, arg5)
}
