// Code generated by mockery. DO NOT EDIT.

package mock_synchronizer

import (
	context "context"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	mock "github.com/stretchr/testify/mock"

	pgstorage "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/storage/pgstorage"
)

// storageSyncQueries is an autogenerated mock type for the storageSyncQueries type
type storageSyncQueries struct {
	mock.Mock
}

type storageSyncQueries_Expecter struct {
	mock *mock.Mock
}

func (_m *storageSyncQueries) EXPECT() *storageSyncQueries_Expecter {
	return &storageSyncQueries_Expecter{mock: &_m.Mock}
}

// AddBlock provides a mock function with given fields: ctx, block, dbTx
func (_m *storageSyncQueries) AddBlock(ctx context.Context, block *entities.L1Block, dbTx entities.Tx) error {
	ret := _m.Called(ctx, block, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for AddBlock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.L1Block, entities.Tx) error); ok {
		r0 = rf(ctx, block, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// storageSyncQueries_AddBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBlock'
type storageSyncQueries_AddBlock_Call struct {
	*mock.Call
}

// AddBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - block *entities.L1Block
//   - dbTx entities.Tx
func (_e *storageSyncQueries_Expecter) AddBlock(ctx interface{}, block interface{}, dbTx interface{}) *storageSyncQueries_AddBlock_Call {
	return &storageSyncQueries_AddBlock_Call{Call: _e.mock.On("AddBlock", ctx, block, dbTx)}
}

func (_c *storageSyncQueries_AddBlock_Call) Run(run func(ctx context.Context, block *entities.L1Block, dbTx entities.Tx)) *storageSyncQueries_AddBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.L1Block), args[2].(entities.Tx))
	})
	return _c
}

func (_c *storageSyncQueries_AddBlock_Call) Return(_a0 error) *storageSyncQueries_AddBlock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *storageSyncQueries_AddBlock_Call) RunAndReturn(run func(context.Context, *entities.L1Block, entities.Tx) error) *storageSyncQueries_AddBlock_Call {
	_c.Call.Return(run)
	return _c
}

// AddSequencedBatches provides a mock function with given fields: ctx, sequence, dbTx
func (_m *storageSyncQueries) AddSequencedBatches(ctx context.Context, sequence *entities.SequencedBatches, dbTx entities.Tx) error {
	ret := _m.Called(ctx, sequence, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for AddSequencedBatches")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.SequencedBatches, entities.Tx) error); ok {
		r0 = rf(ctx, sequence, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// storageSyncQueries_AddSequencedBatches_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSequencedBatches'
type storageSyncQueries_AddSequencedBatches_Call struct {
	*mock.Call
}

// AddSequencedBatches is a helper method to define mock.On call
//   - ctx context.Context
//   - sequence *entities.SequencedBatches
//   - dbTx entities.Tx
func (_e *storageSyncQueries_Expecter) AddSequencedBatches(ctx interface{}, sequence interface{}, dbTx interface{}) *storageSyncQueries_AddSequencedBatches_Call {
	return &storageSyncQueries_AddSequencedBatches_Call{Call: _e.mock.On("AddSequencedBatches", ctx, sequence, dbTx)}
}

func (_c *storageSyncQueries_AddSequencedBatches_Call) Run(run func(ctx context.Context, sequence *entities.SequencedBatches, dbTx entities.Tx)) *storageSyncQueries_AddSequencedBatches_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.SequencedBatches), args[2].(entities.Tx))
	})
	return _c
}

func (_c *storageSyncQueries_AddSequencedBatches_Call) Return(_a0 error) *storageSyncQueries_AddSequencedBatches_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *storageSyncQueries_AddSequencedBatches_Call) RunAndReturn(run func(context.Context, *entities.SequencedBatches, entities.Tx) error) *storageSyncQueries_AddSequencedBatches_Call {
	_c.Call.Return(run)
	return _c
}

// GetBlockByNumber provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *storageSyncQueries) GetBlockByNumber(ctx context.Context, blockNumber uint64, dbTx entities.Tx) (*entities.L1Block, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockByNumber")
	}

	var r0 *entities.L1Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)); ok {
		return rf(ctx, blockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) *entities.L1Block); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.L1Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, entities.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageSyncQueries_GetBlockByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockByNumber'
type storageSyncQueries_GetBlockByNumber_Call struct {
	*mock.Call
}

// GetBlockByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - blockNumber uint64
//   - dbTx entities.Tx
func (_e *storageSyncQueries_Expecter) GetBlockByNumber(ctx interface{}, blockNumber interface{}, dbTx interface{}) *storageSyncQueries_GetBlockByNumber_Call {
	return &storageSyncQueries_GetBlockByNumber_Call{Call: _e.mock.On("GetBlockByNumber", ctx, blockNumber, dbTx)}
}

func (_c *storageSyncQueries_GetBlockByNumber_Call) Run(run func(ctx context.Context, blockNumber uint64, dbTx entities.Tx)) *storageSyncQueries_GetBlockByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *storageSyncQueries_GetBlockByNumber_Call) Return(_a0 *entities.L1Block, _a1 error) *storageSyncQueries_GetBlockByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageSyncQueries_GetBlockByNumber_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)) *storageSyncQueries_GetBlockByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetFirstUncheckedBlock provides a mock function with given fields: ctx, fromBlockNumber, dbTx
func (_m *storageSyncQueries) GetFirstUncheckedBlock(ctx context.Context, fromBlockNumber uint64, dbTx entities.Tx) (*entities.L1Block, error) {
	ret := _m.Called(ctx, fromBlockNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetFirstUncheckedBlock")
	}

	var r0 *entities.L1Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)); ok {
		return rf(ctx, fromBlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) *entities.L1Block); ok {
		r0 = rf(ctx, fromBlockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.L1Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, entities.Tx) error); ok {
		r1 = rf(ctx, fromBlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageSyncQueries_GetFirstUncheckedBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFirstUncheckedBlock'
type storageSyncQueries_GetFirstUncheckedBlock_Call struct {
	*mock.Call
}

// GetFirstUncheckedBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - fromBlockNumber uint64
//   - dbTx entities.Tx
func (_e *storageSyncQueries_Expecter) GetFirstUncheckedBlock(ctx interface{}, fromBlockNumber interface{}, dbTx interface{}) *storageSyncQueries_GetFirstUncheckedBlock_Call {
	return &storageSyncQueries_GetFirstUncheckedBlock_Call{Call: _e.mock.On("GetFirstUncheckedBlock", ctx, fromBlockNumber, dbTx)}
}

func (_c *storageSyncQueries_GetFirstUncheckedBlock_Call) Run(run func(ctx context.Context, fromBlockNumber uint64, dbTx entities.Tx)) *storageSyncQueries_GetFirstUncheckedBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *storageSyncQueries_GetFirstUncheckedBlock_Call) Return(_a0 *entities.L1Block, _a1 error) *storageSyncQueries_GetFirstUncheckedBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageSyncQueries_GetFirstUncheckedBlock_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)) *storageSyncQueries_GetFirstUncheckedBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastBlock provides a mock function with given fields: ctx, dbTx
func (_m *storageSyncQueries) GetLastBlock(ctx context.Context, dbTx entities.Tx) (*entities.L1Block, error) {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetLastBlock")
	}

	var r0 *entities.L1Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Tx) (*entities.L1Block, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.Tx) *entities.L1Block); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.L1Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageSyncQueries_GetLastBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastBlock'
type storageSyncQueries_GetLastBlock_Call struct {
	*mock.Call
}

// GetLastBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx entities.Tx
func (_e *storageSyncQueries_Expecter) GetLastBlock(ctx interface{}, dbTx interface{}) *storageSyncQueries_GetLastBlock_Call {
	return &storageSyncQueries_GetLastBlock_Call{Call: _e.mock.On("GetLastBlock", ctx, dbTx)}
}

func (_c *storageSyncQueries_GetLastBlock_Call) Run(run func(ctx context.Context, dbTx entities.Tx)) *storageSyncQueries_GetLastBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.Tx))
	})
	return _c
}

func (_c *storageSyncQueries_GetLastBlock_Call) Return(_a0 *entities.L1Block, _a1 error) *storageSyncQueries_GetLastBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageSyncQueries_GetLastBlock_Call) RunAndReturn(run func(context.Context, entities.Tx) (*entities.L1Block, error)) *storageSyncQueries_GetLastBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastestVirtualBatchNumber provides a mock function with given fields: ctx, constrains, dbTx
func (_m *storageSyncQueries) GetLastestVirtualBatchNumber(ctx context.Context, constrains *pgstorage.VirtualBatchConstraints, dbTx entities.Tx) (uint64, error) {
	ret := _m.Called(ctx, constrains, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetLastestVirtualBatchNumber")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgstorage.VirtualBatchConstraints, entities.Tx) (uint64, error)); ok {
		return rf(ctx, constrains, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgstorage.VirtualBatchConstraints, entities.Tx) uint64); ok {
		r0 = rf(ctx, constrains, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgstorage.VirtualBatchConstraints, entities.Tx) error); ok {
		r1 = rf(ctx, constrains, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageSyncQueries_GetLastestVirtualBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastestVirtualBatchNumber'
type storageSyncQueries_GetLastestVirtualBatchNumber_Call struct {
	*mock.Call
}

// GetLastestVirtualBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - constrains *pgstorage.VirtualBatchConstraints
//   - dbTx entities.Tx
func (_e *storageSyncQueries_Expecter) GetLastestVirtualBatchNumber(ctx interface{}, constrains interface{}, dbTx interface{}) *storageSyncQueries_GetLastestVirtualBatchNumber_Call {
	return &storageSyncQueries_GetLastestVirtualBatchNumber_Call{Call: _e.mock.On("GetLastestVirtualBatchNumber", ctx, constrains, dbTx)}
}

func (_c *storageSyncQueries_GetLastestVirtualBatchNumber_Call) Run(run func(ctx context.Context, constrains *pgstorage.VirtualBatchConstraints, dbTx entities.Tx)) *storageSyncQueries_GetLastestVirtualBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*pgstorage.VirtualBatchConstraints), args[2].(entities.Tx))
	})
	return _c
}

func (_c *storageSyncQueries_GetLastestVirtualBatchNumber_Call) Return(_a0 uint64, _a1 error) *storageSyncQueries_GetLastestVirtualBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageSyncQueries_GetLastestVirtualBatchNumber_Call) RunAndReturn(run func(context.Context, *pgstorage.VirtualBatchConstraints, entities.Tx) (uint64, error)) *storageSyncQueries_GetLastestVirtualBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetPreviousBlock provides a mock function with given fields: ctx, offset, dbTx
func (_m *storageSyncQueries) GetPreviousBlock(ctx context.Context, offset uint64, dbTx entities.Tx) (*entities.L1Block, error) {
	ret := _m.Called(ctx, offset, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetPreviousBlock")
	}

	var r0 *entities.L1Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)); ok {
		return rf(ctx, offset, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) *entities.L1Block); ok {
		r0 = rf(ctx, offset, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.L1Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, entities.Tx) error); ok {
		r1 = rf(ctx, offset, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageSyncQueries_GetPreviousBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPreviousBlock'
type storageSyncQueries_GetPreviousBlock_Call struct {
	*mock.Call
}

// GetPreviousBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - offset uint64
//   - dbTx entities.Tx
func (_e *storageSyncQueries_Expecter) GetPreviousBlock(ctx interface{}, offset interface{}, dbTx interface{}) *storageSyncQueries_GetPreviousBlock_Call {
	return &storageSyncQueries_GetPreviousBlock_Call{Call: _e.mock.On("GetPreviousBlock", ctx, offset, dbTx)}
}

func (_c *storageSyncQueries_GetPreviousBlock_Call) Run(run func(ctx context.Context, offset uint64, dbTx entities.Tx)) *storageSyncQueries_GetPreviousBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *storageSyncQueries_GetPreviousBlock_Call) Return(_a0 *entities.L1Block, _a1 error) *storageSyncQueries_GetPreviousBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageSyncQueries_GetPreviousBlock_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)) *storageSyncQueries_GetPreviousBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetSequenceByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *storageSyncQueries) GetSequenceByBatchNumber(ctx context.Context, batchNumber uint64, dbTx entities.Tx) (*entities.SequencedBatches, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetSequenceByBatchNumber")
	}

	var r0 *entities.SequencedBatches
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) (*entities.SequencedBatches, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) *entities.SequencedBatches); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.SequencedBatches)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, entities.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageSyncQueries_GetSequenceByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSequenceByBatchNumber'
type storageSyncQueries_GetSequenceByBatchNumber_Call struct {
	*mock.Call
}

// GetSequenceByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx entities.Tx
func (_e *storageSyncQueries_Expecter) GetSequenceByBatchNumber(ctx interface{}, batchNumber interface{}, dbTx interface{}) *storageSyncQueries_GetSequenceByBatchNumber_Call {
	return &storageSyncQueries_GetSequenceByBatchNumber_Call{Call: _e.mock.On("GetSequenceByBatchNumber", ctx, batchNumber, dbTx)}
}

func (_c *storageSyncQueries_GetSequenceByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx entities.Tx)) *storageSyncQueries_GetSequenceByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *storageSyncQueries_GetSequenceByBatchNumber_Call) Return(_a0 *entities.SequencedBatches, _a1 error) *storageSyncQueries_GetSequenceByBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageSyncQueries_GetSequenceByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.SequencedBatches, error)) *storageSyncQueries_GetSequenceByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetVirtualBatchByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *storageSyncQueries) GetVirtualBatchByBatchNumber(ctx context.Context, batchNumber uint64, dbTx entities.Tx) (*entities.VirtualBatch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetVirtualBatchByBatchNumber")
	}

	var r0 *entities.VirtualBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) (*entities.VirtualBatch, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) *entities.VirtualBatch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.VirtualBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, entities.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageSyncQueries_GetVirtualBatchByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVirtualBatchByBatchNumber'
type storageSyncQueries_GetVirtualBatchByBatchNumber_Call struct {
	*mock.Call
}

// GetVirtualBatchByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx entities.Tx
func (_e *storageSyncQueries_Expecter) GetVirtualBatchByBatchNumber(ctx interface{}, batchNumber interface{}, dbTx interface{}) *storageSyncQueries_GetVirtualBatchByBatchNumber_Call {
	return &storageSyncQueries_GetVirtualBatchByBatchNumber_Call{Call: _e.mock.On("GetVirtualBatchByBatchNumber", ctx, batchNumber, dbTx)}
}

func (_c *storageSyncQueries_GetVirtualBatchByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx entities.Tx)) *storageSyncQueries_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *storageSyncQueries_GetVirtualBatchByBatchNumber_Call) Return(_a0 *entities.VirtualBatch, _a1 error) *storageSyncQueries_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageSyncQueries_GetVirtualBatchByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.VirtualBatch, error)) *storageSyncQueries_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// newStorageSyncQueries creates a new instance of storageSyncQueries. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newStorageSyncQueries(t interface {
	mock.TestingT
	Cleanup(func())
}) *storageSyncQueries {
	mock := &storageSyncQueries{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
