// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/common/time.go

// Package common is a generated GoMock package.
package common

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockTimeUtil is a mock of TimeUtil interface.
type MockTimeUtil struct {
	ctrl     *gomock.Controller
	recorder *MockTimeUtilMockRecorder
}

// MockTimeUtilMockRecorder is the mock recorder for MockTimeUtil.
type MockTimeUtilMockRecorder struct {
	mock *MockTimeUtil
}

// NewMockTimeUtil creates a new mock instance.
func NewMockTimeUtil(ctrl *gomock.Controller) *MockTimeUtil {
	mock := &MockTimeUtil{ctrl: ctrl}
	mock.recorder = &MockTimeUtilMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimeUtil) EXPECT() *MockTimeUtilMockRecorder {
	return m.recorder
}

// DaysBetween mocks base method.
func (m *MockTimeUtil) DaysBetween(start, end time.Time) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DaysBetween", start, end)
	ret0, _ := ret[0].(int)
	return ret0
}

// DaysBetween indicates an expected call of DaysBetween.
func (mr *MockTimeUtilMockRecorder) DaysBetween(start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DaysBetween", reflect.TypeOf((*MockTimeUtil)(nil).DaysBetween), start, end)
}

// GetCurrentTime mocks base method.
func (m *MockTimeUtil) GetCurrentTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetCurrentTime indicates an expected call of GetCurrentTime.
func (mr *MockTimeUtilMockRecorder) GetCurrentTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentTime", reflect.TypeOf((*MockTimeUtil)(nil).GetCurrentTime))
}
