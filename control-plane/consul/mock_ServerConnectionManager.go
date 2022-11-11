// Code generated by mockery v2.14.0. DO NOT EDIT.

package consul

import (
	discovery "github.com/hashicorp/consul-server-connection-manager/discovery"
	mock "github.com/stretchr/testify/mock"
)

// MockServerConnectionManager is an autogenerated mock type for the ServerConnectionManager type
type MockServerConnectionManager struct {
	mock.Mock
}

// Run provides a mock function with given fields:
func (_m *MockServerConnectionManager) Run() {
	_m.Called()
}

// State provides a mock function with given fields:
func (_m *MockServerConnectionManager) State() (discovery.State, error) {
	ret := _m.Called()

	var r0 discovery.State
	if rf, ok := ret.Get(0).(func() discovery.State); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(discovery.State)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stop provides a mock function with given fields:
func (_m *MockServerConnectionManager) Stop() {
	_m.Called()
}

type mockConstructorTestingTNewMockServerConnectionManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockServerConnectionManager creates a new instance of MockServerConnectionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockServerConnectionManager(t mockConstructorTestingTNewMockServerConnectionManager) *MockServerConnectionManager {
	mock := &MockServerConnectionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}