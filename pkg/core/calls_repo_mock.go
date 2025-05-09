// Code generated by mockery v2.46.3. DO NOT EDIT.

//go:build !compile

package core

import mock "github.com/stretchr/testify/mock"

// MockCallsRepo is an autogenerated mock type for the CallsRepo type
type MockCallsRepo struct {
	mock.Mock
}

type MockCallsRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCallsRepo) EXPECT() *MockCallsRepo_Expecter {
	return &MockCallsRepo_Expecter{mock: &_m.Mock}
}

// GetCall provides a mock function with given fields: method
func (_m *MockCallsRepo) GetCall(method string) Handler {
	ret := _m.Called(method)

	if len(ret) == 0 {
		panic("no return value specified for GetCall")
	}

	var r0 Handler
	if rf, ok := ret.Get(0).(func(string) Handler); ok {
		r0 = rf(method)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Handler)
		}
	}

	return r0
}

// MockCallsRepo_GetCall_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCall'
type MockCallsRepo_GetCall_Call struct {
	*mock.Call
}

// GetCall is a helper method to define mock.On call
//   - method string
func (_e *MockCallsRepo_Expecter) GetCall(method interface{}) *MockCallsRepo_GetCall_Call {
	return &MockCallsRepo_GetCall_Call{Call: _e.mock.On("GetCall", method)}
}

func (_c *MockCallsRepo_GetCall_Call) Run(run func(method string)) *MockCallsRepo_GetCall_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockCallsRepo_GetCall_Call) Return(_a0 Handler) *MockCallsRepo_GetCall_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCallsRepo_GetCall_Call) RunAndReturn(run func(string) Handler) *MockCallsRepo_GetCall_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCalls provides a mock function with given fields: _a0
func (_m *MockCallsRepo) UpdateCalls(_a0 map[string]Handler) {
	_m.Called(_a0)
}

// MockCallsRepo_UpdateCalls_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCalls'
type MockCallsRepo_UpdateCalls_Call struct {
	*mock.Call
}

// UpdateCalls is a helper method to define mock.On call
//   - _a0 map[string]Handler
func (_e *MockCallsRepo_Expecter) UpdateCalls(_a0 interface{}) *MockCallsRepo_UpdateCalls_Call {
	return &MockCallsRepo_UpdateCalls_Call{Call: _e.mock.On("UpdateCalls", _a0)}
}

func (_c *MockCallsRepo_UpdateCalls_Call) Run(run func(_a0 map[string]Handler)) *MockCallsRepo_UpdateCalls_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]Handler))
	})
	return _c
}

func (_c *MockCallsRepo_UpdateCalls_Call) Return() *MockCallsRepo_UpdateCalls_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCallsRepo_UpdateCalls_Call) RunAndReturn(run func(map[string]Handler)) *MockCallsRepo_UpdateCalls_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCallsRepo creates a new instance of MockCallsRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCallsRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCallsRepo {
	mock := &MockCallsRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
