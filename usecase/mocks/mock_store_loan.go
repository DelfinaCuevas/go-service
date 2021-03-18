// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eiizu/go-service/usecase (interfaces: StoreLoan)

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "github.com/eiizu/go-service/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStoreLoan is a mock of StoreLoan interface
type MockStoreLoan struct {
	ctrl     *gomock.Controller
	recorder *MockStoreLoanMockRecorder
}

// MockStoreLoanMockRecorder is the mock recorder for MockStoreLoan
type MockStoreLoanMockRecorder struct {
	mock *MockStoreLoan
}

// NewMockStoreLoan creates a new mock instance
func NewMockStoreLoan(ctrl *gomock.Controller) *MockStoreLoan {
	mock := &MockStoreLoan{ctrl: ctrl}
	mock.recorder = &MockStoreLoanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStoreLoan) EXPECT() *MockStoreLoanMockRecorder {
	return m.recorder
}

// CreateLoan mocks base method
func (m *MockStoreLoan) CreateLoan(arg0 entity.Loan) (*entity.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoan", arg0)
	ret0, _ := ret[0].(*entity.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoan indicates an expected call of CreateLoan
func (mr *MockStoreLoanMockRecorder) CreateLoan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoan", reflect.TypeOf((*MockStoreLoan)(nil).CreateLoan), arg0)
}

// GetLoan mocks base method
func (m *MockStoreLoan) GetLoan(arg0 map[string]string) (map[string]entity.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoan", arg0)
	ret0, _ := ret[0].(map[string]entity.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoan indicates an expected call of GetLoan
func (mr *MockStoreLoanMockRecorder) GetLoan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoan", reflect.TypeOf((*MockStoreLoan)(nil).GetLoan), arg0)
}

// GetLoan_ mocks base method
func (m *MockStoreLoan) GetLoan_(arg0 map[string]string) (map[string]entity.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoan_", arg0)
	ret0, _ := ret[0].(map[string]entity.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoan_ indicates an expected call of GetLoan_
func (mr *MockStoreLoanMockRecorder) GetLoan_(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoan_", reflect.TypeOf((*MockStoreLoan)(nil).GetLoan_), arg0)
}

// GetLoans mocks base method
func (m *MockStoreLoan) GetLoans() (map[string]entity.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoans")
	ret0, _ := ret[0].(map[string]entity.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoans indicates an expected call of GetLoans
func (mr *MockStoreLoanMockRecorder) GetLoans() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoans", reflect.TypeOf((*MockStoreLoan)(nil).GetLoans))
}

// UpdateLoan mocks base method
func (m *MockStoreLoan) UpdateLoan(arg0 entity.Loan) (*entity.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLoan", arg0)
	ret0, _ := ret[0].(*entity.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLoan indicates an expected call of UpdateLoan
func (mr *MockStoreLoanMockRecorder) UpdateLoan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLoan", reflect.TypeOf((*MockStoreLoan)(nil).UpdateLoan), arg0)
}