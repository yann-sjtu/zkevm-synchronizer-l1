// Code generated by mockery. DO NOT EDIT.

package mock_synchronizer

import (
	context "context"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"

	mock "github.com/stretchr/testify/mock"
)

// stateInterface is an autogenerated mock type for the stateInterface type
type stateInterface struct {
	mock.Mock
}

type stateInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *stateInterface) EXPECT() *stateInterface_Expecter {
	return &stateInterface_Expecter{mock: &_m.Mock}
}

// BeginTransaction provides a mock function with given fields: ctx
func (_m *stateInterface) BeginTransaction(ctx context.Context) (entities.Tx, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BeginTransaction")
	}

	var r0 entities.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (entities.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) entities.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entities.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// stateInterface_BeginTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeginTransaction'
type stateInterface_BeginTransaction_Call struct {
	*mock.Call
}

// BeginTransaction is a helper method to define mock.On call
//   - ctx context.Context
func (_e *stateInterface_Expecter) BeginTransaction(ctx interface{}) *stateInterface_BeginTransaction_Call {
	return &stateInterface_BeginTransaction_Call{Call: _e.mock.On("BeginTransaction", ctx)}
}

func (_c *stateInterface_BeginTransaction_Call) Run(run func(ctx context.Context)) *stateInterface_BeginTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *stateInterface_BeginTransaction_Call) Return(_a0 entities.Tx, _a1 error) *stateInterface_BeginTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *stateInterface_BeginTransaction_Call) RunAndReturn(run func(context.Context) (entities.Tx, error)) *stateInterface_BeginTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// GetPreviousBlock provides a mock function with given fields: ctx, depth, tx
func (_m *stateInterface) GetPreviousBlock(ctx context.Context, depth uint64, tx entities.Tx) (*entities.L1Block, error) {
	ret := _m.Called(ctx, depth, tx)

	if len(ret) == 0 {
		panic("no return value specified for GetPreviousBlock")
	}

	var r0 *entities.L1Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)); ok {
		return rf(ctx, depth, tx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) *entities.L1Block); ok {
		r0 = rf(ctx, depth, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.L1Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, entities.Tx) error); ok {
		r1 = rf(ctx, depth, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// stateInterface_GetPreviousBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPreviousBlock'
type stateInterface_GetPreviousBlock_Call struct {
	*mock.Call
}

// GetPreviousBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - depth uint64
//   - tx entities.Tx
func (_e *stateInterface_Expecter) GetPreviousBlock(ctx interface{}, depth interface{}, tx interface{}) *stateInterface_GetPreviousBlock_Call {
	return &stateInterface_GetPreviousBlock_Call{Call: _e.mock.On("GetPreviousBlock", ctx, depth, tx)}
}

func (_c *stateInterface_GetPreviousBlock_Call) Run(run func(ctx context.Context, depth uint64, tx entities.Tx)) *stateInterface_GetPreviousBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *stateInterface_GetPreviousBlock_Call) Return(_a0 *entities.L1Block, _a1 error) *stateInterface_GetPreviousBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *stateInterface_GetPreviousBlock_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)) *stateInterface_GetPreviousBlock_Call {
	_c.Call.Return(run)
	return _c
}

// newStateInterface creates a new instance of stateInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newStateInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *stateInterface {
	mock := &stateInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
