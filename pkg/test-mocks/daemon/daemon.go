// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/replicatedhq/ship/pkg/lifecycle/daemon (interfaces: Daemon)

// Package daemon is a generated GoMock package.
package daemon

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/replicatedhq/ship/pkg/api"
	daemon "github.com/replicatedhq/ship/pkg/lifecycle/daemon"
)

// MockDaemon is a mock of Daemon interface
type MockDaemon struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonMockRecorder
}

// MockDaemonMockRecorder is the mock recorder for MockDaemon
type MockDaemonMockRecorder struct {
	mock *MockDaemon
}

// NewMockDaemon creates a new mock instance
func NewMockDaemon(ctrl *gomock.Controller) *MockDaemon {
	mock := &MockDaemon{ctrl: ctrl}
	mock.recorder = &MockDaemonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDaemon) EXPECT() *MockDaemonMockRecorder {
	return m.recorder
}

// AllStepsDone mocks base method
func (m *MockDaemon) AllStepsDone(arg0 context.Context) {
	m.ctrl.Call(m, "AllStepsDone", arg0)
}

// AllStepsDone indicates an expected call of AllStepsDone
func (mr *MockDaemonMockRecorder) AllStepsDone(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllStepsDone", reflect.TypeOf((*MockDaemon)(nil).AllStepsDone), arg0)
}

// CleanPreviousStep mocks base method
func (m *MockDaemon) CleanPreviousStep() {
	m.ctrl.Call(m, "CleanPreviousStep")
}

// CleanPreviousStep indicates an expected call of CleanPreviousStep
func (mr *MockDaemonMockRecorder) CleanPreviousStep() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanPreviousStep", reflect.TypeOf((*MockDaemon)(nil).CleanPreviousStep))
}

// ClearProgress mocks base method
func (m *MockDaemon) ClearProgress() {
	m.ctrl.Call(m, "ClearProgress")
}

// ClearProgress indicates an expected call of ClearProgress
func (mr *MockDaemonMockRecorder) ClearProgress() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearProgress", reflect.TypeOf((*MockDaemon)(nil).ClearProgress))
}

// ConfigSavedChan mocks base method
func (m *MockDaemon) ConfigSavedChan() chan interface{} {
	ret := m.ctrl.Call(m, "ConfigSavedChan")
	ret0, _ := ret[0].(chan interface{})
	return ret0
}

// ConfigSavedChan indicates an expected call of ConfigSavedChan
func (mr *MockDaemonMockRecorder) ConfigSavedChan() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigSavedChan", reflect.TypeOf((*MockDaemon)(nil).ConfigSavedChan))
}

// EnsureStarted mocks base method
func (m *MockDaemon) EnsureStarted(arg0 context.Context, arg1 *api.Release) chan error {
	ret := m.ctrl.Call(m, "EnsureStarted", arg0, arg1)
	ret0, _ := ret[0].(chan error)
	return ret0
}

// EnsureStarted indicates an expected call of EnsureStarted
func (mr *MockDaemonMockRecorder) EnsureStarted(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureStarted", reflect.TypeOf((*MockDaemon)(nil).EnsureStarted), arg0, arg1)
}

// GetCurrentConfig mocks base method
func (m *MockDaemon) GetCurrentConfig() map[string]interface{} {
	ret := m.ctrl.Call(m, "GetCurrentConfig")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// GetCurrentConfig indicates an expected call of GetCurrentConfig
func (mr *MockDaemonMockRecorder) GetCurrentConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentConfig", reflect.TypeOf((*MockDaemon)(nil).GetCurrentConfig))
}

// KubectlConfirmedChan mocks base method
func (m *MockDaemon) KubectlConfirmedChan() chan bool {
	ret := m.ctrl.Call(m, "KubectlConfirmedChan")
	ret0, _ := ret[0].(chan bool)
	return ret0
}

// KubectlConfirmedChan indicates an expected call of KubectlConfirmedChan
func (mr *MockDaemonMockRecorder) KubectlConfirmedChan() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubectlConfirmedChan", reflect.TypeOf((*MockDaemon)(nil).KubectlConfirmedChan))
}

// KustomizeSavedChan mocks base method
func (m *MockDaemon) KustomizeSavedChan() chan interface{} {
	ret := m.ctrl.Call(m, "KustomizeSavedChan")
	ret0, _ := ret[0].(chan interface{})
	return ret0
}

// KustomizeSavedChan indicates an expected call of KustomizeSavedChan
func (mr *MockDaemonMockRecorder) KustomizeSavedChan() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KustomizeSavedChan", reflect.TypeOf((*MockDaemon)(nil).KustomizeSavedChan))
}

// MessageConfirmedChan mocks base method
func (m *MockDaemon) MessageConfirmedChan() chan string {
	ret := m.ctrl.Call(m, "MessageConfirmedChan")
	ret0, _ := ret[0].(chan string)
	return ret0
}

// MessageConfirmedChan indicates an expected call of MessageConfirmedChan
func (mr *MockDaemonMockRecorder) MessageConfirmedChan() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageConfirmedChan", reflect.TypeOf((*MockDaemon)(nil).MessageConfirmedChan))
}

// PushHelmIntroStep mocks base method
func (m *MockDaemon) PushHelmIntroStep(arg0 context.Context, arg1 daemon.HelmIntro, arg2 []daemon.Action) {
	m.ctrl.Call(m, "PushHelmIntroStep", arg0, arg1, arg2)
}

// PushHelmIntroStep indicates an expected call of PushHelmIntroStep
func (mr *MockDaemonMockRecorder) PushHelmIntroStep(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushHelmIntroStep", reflect.TypeOf((*MockDaemon)(nil).PushHelmIntroStep), arg0, arg1, arg2)
}

// PushHelmValuesStep mocks base method
func (m *MockDaemon) PushHelmValuesStep(arg0 context.Context, arg1 daemon.HelmValues, arg2 []daemon.Action) {
	m.ctrl.Call(m, "PushHelmValuesStep", arg0, arg1, arg2)
}

// PushHelmValuesStep indicates an expected call of PushHelmValuesStep
func (mr *MockDaemonMockRecorder) PushHelmValuesStep(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushHelmValuesStep", reflect.TypeOf((*MockDaemon)(nil).PushHelmValuesStep), arg0, arg1, arg2)
}

// PushKustomizeStep mocks base method
func (m *MockDaemon) PushKustomizeStep(arg0 context.Context, arg1 daemon.Kustomize) {
	m.ctrl.Call(m, "PushKustomizeStep", arg0, arg1)
}

// PushKustomizeStep indicates an expected call of PushKustomizeStep
func (mr *MockDaemonMockRecorder) PushKustomizeStep(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushKustomizeStep", reflect.TypeOf((*MockDaemon)(nil).PushKustomizeStep), arg0, arg1)
}

// PushMessageStep mocks base method
func (m *MockDaemon) PushMessageStep(arg0 context.Context, arg1 daemon.Message, arg2 []daemon.Action) {
	m.ctrl.Call(m, "PushMessageStep", arg0, arg1, arg2)
}

// PushMessageStep indicates an expected call of PushMessageStep
func (mr *MockDaemonMockRecorder) PushMessageStep(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushMessageStep", reflect.TypeOf((*MockDaemon)(nil).PushMessageStep), arg0, arg1, arg2)
}

// PushRenderStep mocks base method
func (m *MockDaemon) PushRenderStep(arg0 context.Context, arg1 daemon.Render) {
	m.ctrl.Call(m, "PushRenderStep", arg0, arg1)
}

// PushRenderStep indicates an expected call of PushRenderStep
func (mr *MockDaemonMockRecorder) PushRenderStep(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushRenderStep", reflect.TypeOf((*MockDaemon)(nil).PushRenderStep), arg0, arg1)
}

// PushStreamStep mocks base method
func (m *MockDaemon) PushStreamStep(arg0 context.Context, arg1 <-chan daemon.Message) {
	m.ctrl.Call(m, "PushStreamStep", arg0, arg1)
}

// PushStreamStep indicates an expected call of PushStreamStep
func (mr *MockDaemonMockRecorder) PushStreamStep(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushStreamStep", reflect.TypeOf((*MockDaemon)(nil).PushStreamStep), arg0, arg1)
}

// SetProgress mocks base method
func (m *MockDaemon) SetProgress(arg0 daemon.Progress) {
	m.ctrl.Call(m, "SetProgress", arg0)
}

// SetProgress indicates an expected call of SetProgress
func (mr *MockDaemonMockRecorder) SetProgress(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProgress", reflect.TypeOf((*MockDaemon)(nil).SetProgress), arg0)
}

// SetStepName mocks base method
func (m *MockDaemon) SetStepName(arg0 context.Context, arg1 string) {
	m.ctrl.Call(m, "SetStepName", arg0, arg1)
}

// SetStepName indicates an expected call of SetStepName
func (mr *MockDaemonMockRecorder) SetStepName(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStepName", reflect.TypeOf((*MockDaemon)(nil).SetStepName), arg0, arg1)
}

// TerraformConfirmedChan mocks base method
func (m *MockDaemon) TerraformConfirmedChan() chan bool {
	ret := m.ctrl.Call(m, "TerraformConfirmedChan")
	ret0, _ := ret[0].(chan bool)
	return ret0
}

// TerraformConfirmedChan indicates an expected call of TerraformConfirmedChan
func (mr *MockDaemonMockRecorder) TerraformConfirmedChan() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerraformConfirmedChan", reflect.TypeOf((*MockDaemon)(nil).TerraformConfirmedChan))
}
