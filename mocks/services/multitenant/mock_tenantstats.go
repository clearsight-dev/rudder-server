// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/services/multitenant (interfaces: MultiTenantI)

// Package mock_tenantstats is a generated GoMock package.
package mock_tenantstats

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockMultiTenantI is a mock of MultiTenantI interface.
type MockMultiTenantI struct {
	ctrl     *gomock.Controller
	recorder *MockMultiTenantIMockRecorder
}

// MockMultiTenantIMockRecorder is the mock recorder for MockMultiTenantI.
type MockMultiTenantIMockRecorder struct {
	mock *MockMultiTenantI
}

// NewMockMultiTenantI creates a new mock instance.
func NewMockMultiTenantI(ctrl *gomock.Controller) *MockMultiTenantI {
	mock := &MockMultiTenantI{ctrl: ctrl}
	mock.recorder = &MockMultiTenantIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMultiTenantI) EXPECT() *MockMultiTenantIMockRecorder {
	return m.recorder
}

// CalculateSuccessFailureCounts mocks base method.
func (m *MockMultiTenantI) CalculateSuccessFailureCounts(arg0, arg1 string, arg2, arg3 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CalculateSuccessFailureCounts", arg0, arg1, arg2, arg3)
}

// CalculateSuccessFailureCounts indicates an expected call of CalculateSuccessFailureCounts.
func (mr *MockMultiTenantIMockRecorder) CalculateSuccessFailureCounts(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateSuccessFailureCounts", reflect.TypeOf((*MockMultiTenantI)(nil).CalculateSuccessFailureCounts), arg0, arg1, arg2, arg3)
}

// GetRouterPickupJobs mocks base method.
func (m *MockMultiTenantI) GetRouterPickupJobs(arg0 string, arg1 int, arg2 time.Duration, arg3 int) map[string]int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouterPickupJobs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(map[string]int)
	return ret0
}

// GetRouterPickupJobs indicates an expected call of GetRouterPickupJobs.
func (mr *MockMultiTenantIMockRecorder) GetRouterPickupJobs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouterPickupJobs", reflect.TypeOf((*MockMultiTenantI)(nil).GetRouterPickupJobs), arg0, arg1, arg2, arg3)
}

// ReportProcLoopAddStats mocks base method.
func (m *MockMultiTenantI) ReportProcLoopAddStats(arg0 map[string]map[string]int, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportProcLoopAddStats", arg0, arg1)
}

// ReportProcLoopAddStats indicates an expected call of ReportProcLoopAddStats.
func (mr *MockMultiTenantIMockRecorder) ReportProcLoopAddStats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportProcLoopAddStats", reflect.TypeOf((*MockMultiTenantI)(nil).ReportProcLoopAddStats), arg0, arg1)
}

// Start mocks base method.
func (m *MockMultiTenantI) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockMultiTenantIMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockMultiTenantI)(nil).Start))
}

// Stop mocks base method.
func (m *MockMultiTenantI) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockMultiTenantIMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockMultiTenantI)(nil).Stop))
}

// UpdateWorkspaceLatencyMap mocks base method.
func (m *MockMultiTenantI) UpdateWorkspaceLatencyMap(arg0, arg1 string, arg2 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateWorkspaceLatencyMap", arg0, arg1, arg2)
}

// UpdateWorkspaceLatencyMap indicates an expected call of UpdateWorkspaceLatencyMap.
func (mr *MockMultiTenantIMockRecorder) UpdateWorkspaceLatencyMap(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkspaceLatencyMap", reflect.TypeOf((*MockMultiTenantI)(nil).UpdateWorkspaceLatencyMap), arg0, arg1, arg2)
}
