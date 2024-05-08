// Code generated by mockery. DO NOT EDIT.

package mock_synchronizer

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// EthermanReorgManager is an autogenerated mock type for the EthermanReorgManager type
type EthermanReorgManager struct {
	mock.Mock
}

type EthermanReorgManager_Expecter struct {
	mock *mock.Mock
}

func (_m *EthermanReorgManager) EXPECT() *EthermanReorgManager_Expecter {
	return &EthermanReorgManager_Expecter{mock: &_m.Mock}
}

// EthBlockByNumber provides a mock function with given fields: ctx, blockNumber
func (_m *EthermanReorgManager) EthBlockByNumber(ctx context.Context, blockNumber uint64) (*types.Block, error) {
	ret := _m.Called(ctx, blockNumber)

	if len(ret) == 0 {
		panic("no return value specified for EthBlockByNumber")
	}

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*types.Block, error)); ok {
		return rf(ctx, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *types.Block); ok {
		r0 = rf(ctx, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthermanReorgManager_EthBlockByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EthBlockByNumber'
type EthermanReorgManager_EthBlockByNumber_Call struct {
	*mock.Call
}

// EthBlockByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - blockNumber uint64
func (_e *EthermanReorgManager_Expecter) EthBlockByNumber(ctx interface{}, blockNumber interface{}) *EthermanReorgManager_EthBlockByNumber_Call {
	return &EthermanReorgManager_EthBlockByNumber_Call{Call: _e.mock.On("EthBlockByNumber", ctx, blockNumber)}
}

func (_c *EthermanReorgManager_EthBlockByNumber_Call) Run(run func(ctx context.Context, blockNumber uint64)) *EthermanReorgManager_EthBlockByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *EthermanReorgManager_EthBlockByNumber_Call) Return(_a0 *types.Block, _a1 error) *EthermanReorgManager_EthBlockByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthermanReorgManager_EthBlockByNumber_Call) RunAndReturn(run func(context.Context, uint64) (*types.Block, error)) *EthermanReorgManager_EthBlockByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// NewEthermanReorgManager creates a new instance of EthermanReorgManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEthermanReorgManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *EthermanReorgManager {
	mock := &EthermanReorgManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
