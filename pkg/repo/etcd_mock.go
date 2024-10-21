// Code generated by mockery v2.45.1. DO NOT EDIT.

//go:build !compile

package repo

import (
	context "context"

	clientv3 "go.etcd.io/etcd/clientv3"

	mock "github.com/stretchr/testify/mock"
)

// MockEtcd is an autogenerated mock type for the Etcd type
type MockEtcd struct {
	mock.Mock
}

type MockEtcd_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEtcd) EXPECT() *MockEtcd_Expecter {
	return &MockEtcd_Expecter{mock: &_m.Mock}
}

// Client provides a mock function with given fields: cfg
func (_m *MockEtcd) Client(cfg EtcdConfig) (*clientv3.Client, error) {
	ret := _m.Called(cfg)

	if len(ret) == 0 {
		panic("no return value specified for Client")
	}

	var r0 *clientv3.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(EtcdConfig) (*clientv3.Client, error)); ok {
		return rf(cfg)
	}
	if rf, ok := ret.Get(0).(func(EtcdConfig) *clientv3.Client); ok {
		r0 = rf(cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientv3.Client)
		}
	}

	if rf, ok := ret.Get(1).(func(EtcdConfig) error); ok {
		r1 = rf(cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEtcd_Client_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Client'
type MockEtcd_Client_Call struct {
	*mock.Call
}

// Client is a helper method to define mock.On call
//   - cfg EtcdConfig
func (_e *MockEtcd_Expecter) Client(cfg interface{}) *MockEtcd_Client_Call {
	return &MockEtcd_Client_Call{Call: _e.mock.On("Client", cfg)}
}

func (_c *MockEtcd_Client_Call) Run(run func(cfg EtcdConfig)) *MockEtcd_Client_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EtcdConfig))
	})
	return _c
}

func (_c *MockEtcd_Client_Call) Return(_a0 *clientv3.Client, _a1 error) *MockEtcd_Client_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEtcd_Client_Call) RunAndReturn(run func(EtcdConfig) (*clientv3.Client, error)) *MockEtcd_Client_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: ctx, cli, key, value
func (_m *MockEtcd) Put(ctx context.Context, cli *clientv3.Client, key string, value string) error {
	ret := _m.Called(ctx, cli, key, value)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *clientv3.Client, string, string) error); ok {
		r0 = rf(ctx, cli, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEtcd_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockEtcd_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - ctx context.Context
//   - cli *clientv3.Client
//   - key string
//   - value string
func (_e *MockEtcd_Expecter) Put(ctx interface{}, cli interface{}, key interface{}, value interface{}) *MockEtcd_Put_Call {
	return &MockEtcd_Put_Call{Call: _e.mock.On("Put", ctx, cli, key, value)}
}

func (_c *MockEtcd_Put_Call) Run(run func(ctx context.Context, cli *clientv3.Client, key string, value string)) *MockEtcd_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*clientv3.Client), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockEtcd_Put_Call) Return(_a0 error) *MockEtcd_Put_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEtcd_Put_Call) RunAndReturn(run func(context.Context, *clientv3.Client, string, string) error) *MockEtcd_Put_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEtcd creates a new instance of MockEtcd. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEtcd(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEtcd {
	mock := &MockEtcd{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}