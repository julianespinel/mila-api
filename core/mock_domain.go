// Code generated by MockGen. DO NOT EDIT.
// Source: core/domain.go

// Package core is a generated GoMock package.
package core

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/julianespinel/mila-api/models"
	reflect "reflect"
	time "time"
)

// MockMilaDomain is a mock of MilaDomain interface
type MockMilaDomain struct {
	ctrl     *gomock.Controller
	recorder *MockMilaDomainMockRecorder
}

// MockMilaDomainMockRecorder is the mock recorder for MockMilaDomain
type MockMilaDomainMockRecorder struct {
	mock *MockMilaDomain
}

// NewMockMilaDomain creates a new mock instance
func NewMockMilaDomain(ctrl *gomock.Controller) *MockMilaDomain {
	mock := &MockMilaDomain{ctrl: ctrl}
	mock.recorder = &MockMilaDomainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMilaDomain) EXPECT() *MockMilaDomainMockRecorder {
	return m.recorder
}

// updateDailyStocks mocks base method
func (m *MockMilaDomain) updateDailyStocks(date time.Time) error {
	ret := m.ctrl.Call(m, "updateDailyStocks", date)
	ret0, _ := ret[0].(error)
	return ret0
}

// updateDailyStocks indicates an expected call of updateDailyStocks
func (mr *MockMilaDomainMockRecorder) updateDailyStocks(date interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateDailyStocks", reflect.TypeOf((*MockMilaDomain)(nil).updateDailyStocks), date)
}

// getCurrentDayStocks mocks base method
func (m *MockMilaDomain) getCurrentDayStocks(country string) []models.Stock {
	ret := m.ctrl.Call(m, "getCurrentDayStocks", country)
	ret0, _ := ret[0].([]models.Stock)
	return ret0
}

// getCurrentDayStocks indicates an expected call of getCurrentDayStocks
func (mr *MockMilaDomainMockRecorder) getCurrentDayStocks(country interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getCurrentDayStocks", reflect.TypeOf((*MockMilaDomain)(nil).getCurrentDayStocks), country)
}