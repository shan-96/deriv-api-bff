// Code generated by mockery v2.45.1. DO NOT EDIT.

//go:build !compile

package handler

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockWaitComposer is an autogenerated mock type for the WaitComposer type
type MockWaitComposer struct {
	mock.Mock
}

type MockWaitComposer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWaitComposer) EXPECT() *MockWaitComposer_Expecter {
	return &MockWaitComposer_Expecter{mock: &_m.Mock}
}

// Compose provides a mock function with given fields:
func (_m *MockWaitComposer) Compose() (map[string]interface{}, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Compose")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]interface{}, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWaitComposer_Compose_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Compose'
type MockWaitComposer_Compose_Call struct {
	*mock.Call
}

// Compose is a helper method to define mock.On call
func (_e *MockWaitComposer_Expecter) Compose() *MockWaitComposer_Compose_Call {
	return &MockWaitComposer_Compose_Call{Call: _e.mock.On("Compose")}
}

func (_c *MockWaitComposer_Compose_Call) Run(run func()) *MockWaitComposer_Compose_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWaitComposer_Compose_Call) Return(_a0 map[string]interface{}, _a1 error) *MockWaitComposer_Compose_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWaitComposer_Compose_Call) RunAndReturn(run func() (map[string]interface{}, error)) *MockWaitComposer_Compose_Call {
	_c.Call.Return(run)
	return _c
}

// Prepare provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockWaitComposer) Prepare(_a0 context.Context, _a1 string, _a2 Parser) (int64, map[string]interface{}, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Prepare")
	}

	var r0 int64
	var r1 map[string]interface{}
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, Parser) (int64, map[string]interface{}, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, Parser) int64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, Parser) map[string]interface{}); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, Parser) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockWaitComposer_Prepare_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prepare'
type MockWaitComposer_Prepare_Call struct {
	*mock.Call
}

// Prepare is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 Parser
func (_e *MockWaitComposer_Expecter) Prepare(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockWaitComposer_Prepare_Call {
	return &MockWaitComposer_Prepare_Call{Call: _e.mock.On("Prepare", _a0, _a1, _a2)}
}

func (_c *MockWaitComposer_Prepare_Call) Run(run func(_a0 context.Context, _a1 string, _a2 Parser)) *MockWaitComposer_Prepare_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(Parser))
	})
	return _c
}

func (_c *MockWaitComposer_Prepare_Call) Return(_a0 int64, _a1 map[string]interface{}, _a2 error) *MockWaitComposer_Prepare_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockWaitComposer_Prepare_Call) RunAndReturn(run func(context.Context, string, Parser) (int64, map[string]interface{}, error)) *MockWaitComposer_Prepare_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWaitComposer creates a new instance of MockWaitComposer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWaitComposer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWaitComposer {
	mock := &MockWaitComposer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
