// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package mock_ssm is a generated GoMock package.
package mock_ssm

import (
	reflect "reflect"

	ssm "github.com/aws/aws-sdk-go/service/ssm"
	gomock "github.com/golang/mock/gomock"
)

// MockAWSIface is a mock of AWSIface interface
type MockAWSIface struct {
	ctrl     *gomock.Controller
	recorder *MockAWSIfaceMockRecorder
}

// MockAWSIfaceMockRecorder is the mock recorder for MockAWSIface
type MockAWSIfaceMockRecorder struct {
	mock *MockAWSIface
}

// NewMockAWSIface creates a new mock instance
func NewMockAWSIface(ctrl *gomock.Controller) *MockAWSIface {
	mock := &MockAWSIface{ctrl: ctrl}
	mock.recorder = &MockAWSIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAWSIface) EXPECT() *MockAWSIfaceMockRecorder {
	return m.recorder
}

// PutParameter mocks base method
func (m *MockAWSIface) PutParameter(arg0 *ssm.PutParameterInput) (*ssm.PutParameterOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutParameter", arg0)
	ret0, _ := ret[0].(*ssm.PutParameterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutParameter indicates an expected call of PutParameter
func (mr *MockAWSIfaceMockRecorder) PutParameter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutParameter", reflect.TypeOf((*MockAWSIface)(nil).PutParameter), arg0)
}

// GetParameter mocks base method
func (m *MockAWSIface) GetParameter(arg0 *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParameter", arg0)
	ret0 := CustomMockParameterOutput()
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParameter indicates an expected call of GetParameter
func (mr *MockAWSIfaceMockRecorder) GetParameter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParameter", reflect.TypeOf((*MockAWSIface)(nil).GetParameter), arg0)
}

// CustomMockParameterOutput mock the response of *ssm.GetParameterOutput
func CustomMockParameterOutput() *ssm.GetParameterOutput {
	v := "This is a secret"
	return &ssm.GetParameterOutput{
		Parameter: &ssm.Parameter{
			Value: &v,
		},
	}
}
