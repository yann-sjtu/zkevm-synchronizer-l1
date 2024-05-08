// Code generated by mockery. DO NOT EDIT.

package mock_l1_check_block

import (
	context "context"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"

	mock "github.com/stretchr/testify/mock"
)

// StatePreCheckInterfacer is an autogenerated mock type for the StatePreCheckInterfacer type
type StatePreCheckInterfacer struct {
	mock.Mock
}

type StatePreCheckInterfacer_Expecter struct {
	mock *mock.Mock
}

func (_m *StatePreCheckInterfacer) EXPECT() *StatePreCheckInterfacer_Expecter {
	return &StatePreCheckInterfacer_Expecter{mock: &_m.Mock}
}

// GetUncheckedBlocks provides a mock function with given fields: ctx, fromBlockNumber, toBlockNumber, dbTx
func (_m *StatePreCheckInterfacer) GetUncheckedBlocks(ctx context.Context, fromBlockNumber uint64, toBlockNumber uint64, dbTx entities.Tx) ([]*entities.L1Block, error) {
	ret := _m.Called(ctx, fromBlockNumber, toBlockNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetUncheckedBlocks")
	}

	var r0 []*entities.L1Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, entities.Tx) ([]*entities.L1Block, error)); ok {
		return rf(ctx, fromBlockNumber, toBlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, entities.Tx) []*entities.L1Block); ok {
		r0 = rf(ctx, fromBlockNumber, toBlockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.L1Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, entities.Tx) error); ok {
		r1 = rf(ctx, fromBlockNumber, toBlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatePreCheckInterfacer_GetUncheckedBlocks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUncheckedBlocks'
type StatePreCheckInterfacer_GetUncheckedBlocks_Call struct {
	*mock.Call
}

// GetUncheckedBlocks is a helper method to define mock.On call
//   - ctx context.Context
//   - fromBlockNumber uint64
//   - toBlockNumber uint64
//   - dbTx entities.Tx
func (_e *StatePreCheckInterfacer_Expecter) GetUncheckedBlocks(ctx interface{}, fromBlockNumber interface{}, toBlockNumber interface{}, dbTx interface{}) *StatePreCheckInterfacer_GetUncheckedBlocks_Call {
	return &StatePreCheckInterfacer_GetUncheckedBlocks_Call{Call: _e.mock.On("GetUncheckedBlocks", ctx, fromBlockNumber, toBlockNumber, dbTx)}
}

func (_c *StatePreCheckInterfacer_GetUncheckedBlocks_Call) Run(run func(ctx context.Context, fromBlockNumber uint64, toBlockNumber uint64, dbTx entities.Tx)) *StatePreCheckInterfacer_GetUncheckedBlocks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(uint64), args[3].(entities.Tx))
	})
	return _c
}

func (_c *StatePreCheckInterfacer_GetUncheckedBlocks_Call) Return(_a0 []*entities.L1Block, _a1 error) *StatePreCheckInterfacer_GetUncheckedBlocks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StatePreCheckInterfacer_GetUncheckedBlocks_Call) RunAndReturn(run func(context.Context, uint64, uint64, entities.Tx) ([]*entities.L1Block, error)) *StatePreCheckInterfacer_GetUncheckedBlocks_Call {
	_c.Call.Return(run)
	return _c
}

// NewStatePreCheckInterfacer creates a new instance of StatePreCheckInterfacer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStatePreCheckInterfacer(t interface {
	mock.TestingT
	Cleanup(func())
}) *StatePreCheckInterfacer {
	mock := &StatePreCheckInterfacer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
