// Code generated by mockery. DO NOT EDIT.

package mock_synchronizer

import (
	context "context"

	l1_check_block "github.com/0xPolygonHermez/zkevm-synchronizer-l1/synchronizer/l1_check_block"
	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"
)

// StateInterfacer is an autogenerated mock type for the StateInterfacer type
type StateInterfacer struct {
	mock.Mock
}

type StateInterfacer_Expecter struct {
	mock *mock.Mock
}

func (_m *StateInterfacer) EXPECT() *StateInterfacer_Expecter {
	return &StateInterfacer_Expecter{mock: &_m.Mock}
}

// GetFirstUncheckedBlock provides a mock function with given fields: ctx, fromBlockNumber, dbTx
func (_m *StateInterfacer) GetFirstUncheckedBlock(ctx context.Context, fromBlockNumber uint64, dbTx pgx.Tx) (*l1_check_block.L1Block, error) {
	ret := _m.Called(ctx, fromBlockNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetFirstUncheckedBlock")
	}

	var r0 *l1_check_block.L1Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*l1_check_block.L1Block, error)); ok {
		return rf(ctx, fromBlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *l1_check_block.L1Block); ok {
		r0 = rf(ctx, fromBlockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*l1_check_block.L1Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, fromBlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateInterfacer_GetFirstUncheckedBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFirstUncheckedBlock'
type StateInterfacer_GetFirstUncheckedBlock_Call struct {
	*mock.Call
}

// GetFirstUncheckedBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - fromBlockNumber uint64
//   - dbTx pgx.Tx
func (_e *StateInterfacer_Expecter) GetFirstUncheckedBlock(ctx interface{}, fromBlockNumber interface{}, dbTx interface{}) *StateInterfacer_GetFirstUncheckedBlock_Call {
	return &StateInterfacer_GetFirstUncheckedBlock_Call{Call: _e.mock.On("GetFirstUncheckedBlock", ctx, fromBlockNumber, dbTx)}
}

func (_c *StateInterfacer_GetFirstUncheckedBlock_Call) Run(run func(ctx context.Context, fromBlockNumber uint64, dbTx pgx.Tx)) *StateInterfacer_GetFirstUncheckedBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *StateInterfacer_GetFirstUncheckedBlock_Call) Return(_a0 *l1_check_block.L1Block, _a1 error) *StateInterfacer_GetFirstUncheckedBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StateInterfacer_GetFirstUncheckedBlock_Call) RunAndReturn(run func(context.Context, uint64, pgx.Tx) (*l1_check_block.L1Block, error)) *StateInterfacer_GetFirstUncheckedBlock_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCheckedBlockByNumber provides a mock function with given fields: ctx, blockNumber, newCheckedStatus, dbTx
func (_m *StateInterfacer) UpdateCheckedBlockByNumber(ctx context.Context, blockNumber uint64, newCheckedStatus bool, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, blockNumber, newCheckedStatus, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCheckedBlockByNumber")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, bool, pgx.Tx) error); ok {
		r0 = rf(ctx, blockNumber, newCheckedStatus, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StateInterfacer_UpdateCheckedBlockByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCheckedBlockByNumber'
type StateInterfacer_UpdateCheckedBlockByNumber_Call struct {
	*mock.Call
}

// UpdateCheckedBlockByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - blockNumber uint64
//   - newCheckedStatus bool
//   - dbTx pgx.Tx
func (_e *StateInterfacer_Expecter) UpdateCheckedBlockByNumber(ctx interface{}, blockNumber interface{}, newCheckedStatus interface{}, dbTx interface{}) *StateInterfacer_UpdateCheckedBlockByNumber_Call {
	return &StateInterfacer_UpdateCheckedBlockByNumber_Call{Call: _e.mock.On("UpdateCheckedBlockByNumber", ctx, blockNumber, newCheckedStatus, dbTx)}
}

func (_c *StateInterfacer_UpdateCheckedBlockByNumber_Call) Run(run func(ctx context.Context, blockNumber uint64, newCheckedStatus bool, dbTx pgx.Tx)) *StateInterfacer_UpdateCheckedBlockByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(bool), args[3].(pgx.Tx))
	})
	return _c
}

func (_c *StateInterfacer_UpdateCheckedBlockByNumber_Call) Return(_a0 error) *StateInterfacer_UpdateCheckedBlockByNumber_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StateInterfacer_UpdateCheckedBlockByNumber_Call) RunAndReturn(run func(context.Context, uint64, bool, pgx.Tx) error) *StateInterfacer_UpdateCheckedBlockByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// NewStateInterfacer creates a new instance of StateInterfacer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStateInterfacer(t interface {
	mock.TestingT
	Cleanup(func())
}) *StateInterfacer {
	mock := &StateInterfacer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
