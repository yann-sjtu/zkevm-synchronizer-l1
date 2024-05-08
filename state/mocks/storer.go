// Code generated by mockery. DO NOT EDIT.

package mock_state

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"

	mock "github.com/stretchr/testify/mock"
)

// Storer is an autogenerated mock type for the Storer type
type Storer struct {
	mock.Mock
}

type Storer_Expecter struct {
	mock *mock.Mock
}

func (_m *Storer) EXPECT() *Storer_Expecter {
	return &Storer_Expecter{mock: &_m.Mock}
}

// AddBlock provides a mock function with given fields: ctx, block, dbTx
func (_m *Storer) AddBlock(ctx context.Context, block *entities.L1Block, dbTx entities.Tx) error {
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

// Storer_AddBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBlock'
type Storer_AddBlock_Call struct {
	*mock.Call
}

// AddBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - block *entities.L1Block
//   - dbTx entities.Tx
func (_e *Storer_Expecter) AddBlock(ctx interface{}, block interface{}, dbTx interface{}) *Storer_AddBlock_Call {
	return &Storer_AddBlock_Call{Call: _e.mock.On("AddBlock", ctx, block, dbTx)}
}

func (_c *Storer_AddBlock_Call) Run(run func(ctx context.Context, block *entities.L1Block, dbTx entities.Tx)) *Storer_AddBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.L1Block), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_AddBlock_Call) Return(_a0 error) *Storer_AddBlock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storer_AddBlock_Call) RunAndReturn(run func(context.Context, *entities.L1Block, entities.Tx) error) *Storer_AddBlock_Call {
	_c.Call.Return(run)
	return _c
}

// AddForkID provides a mock function with given fields: ctx, forkID, dbTx
func (_m *Storer) AddForkID(ctx context.Context, forkID entities.ForkIDInterval, dbTx entities.Tx) error {
	ret := _m.Called(ctx, forkID, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for AddForkID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.ForkIDInterval, entities.Tx) error); ok {
		r0 = rf(ctx, forkID, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storer_AddForkID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddForkID'
type Storer_AddForkID_Call struct {
	*mock.Call
}

// AddForkID is a helper method to define mock.On call
//   - ctx context.Context
//   - forkID entities.ForkIDInterval
//   - dbTx entities.Tx
func (_e *Storer_Expecter) AddForkID(ctx interface{}, forkID interface{}, dbTx interface{}) *Storer_AddForkID_Call {
	return &Storer_AddForkID_Call{Call: _e.mock.On("AddForkID", ctx, forkID, dbTx)}
}

func (_c *Storer_AddForkID_Call) Run(run func(ctx context.Context, forkID entities.ForkIDInterval, dbTx entities.Tx)) *Storer_AddForkID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.ForkIDInterval), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_AddForkID_Call) Return(_a0 error) *Storer_AddForkID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storer_AddForkID_Call) RunAndReturn(run func(context.Context, entities.ForkIDInterval, entities.Tx) error) *Storer_AddForkID_Call {
	_c.Call.Return(run)
	return _c
}

// AddL1InfoTreeLeaf provides a mock function with given fields: ctx, exitRoot, dbTx
func (_m *Storer) AddL1InfoTreeLeaf(ctx context.Context, exitRoot *entities.L1InfoTreeLeaf, dbTx entities.Tx) error {
	ret := _m.Called(ctx, exitRoot, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for AddL1InfoTreeLeaf")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.L1InfoTreeLeaf, entities.Tx) error); ok {
		r0 = rf(ctx, exitRoot, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storer_AddL1InfoTreeLeaf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddL1InfoTreeLeaf'
type Storer_AddL1InfoTreeLeaf_Call struct {
	*mock.Call
}

// AddL1InfoTreeLeaf is a helper method to define mock.On call
//   - ctx context.Context
//   - exitRoot *entities.L1InfoTreeLeaf
//   - dbTx entities.Tx
func (_e *Storer_Expecter) AddL1InfoTreeLeaf(ctx interface{}, exitRoot interface{}, dbTx interface{}) *Storer_AddL1InfoTreeLeaf_Call {
	return &Storer_AddL1InfoTreeLeaf_Call{Call: _e.mock.On("AddL1InfoTreeLeaf", ctx, exitRoot, dbTx)}
}

func (_c *Storer_AddL1InfoTreeLeaf_Call) Run(run func(ctx context.Context, exitRoot *entities.L1InfoTreeLeaf, dbTx entities.Tx)) *Storer_AddL1InfoTreeLeaf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.L1InfoTreeLeaf), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_AddL1InfoTreeLeaf_Call) Return(_a0 error) *Storer_AddL1InfoTreeLeaf_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storer_AddL1InfoTreeLeaf_Call) RunAndReturn(run func(context.Context, *entities.L1InfoTreeLeaf, entities.Tx) error) *Storer_AddL1InfoTreeLeaf_Call {
	_c.Call.Return(run)
	return _c
}

// AddSequencedBatches provides a mock function with given fields: ctx, sequence, dbTx
func (_m *Storer) AddSequencedBatches(ctx context.Context, sequence *entities.SequencedBatches, dbTx entities.Tx) error {
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

// Storer_AddSequencedBatches_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSequencedBatches'
type Storer_AddSequencedBatches_Call struct {
	*mock.Call
}

// AddSequencedBatches is a helper method to define mock.On call
//   - ctx context.Context
//   - sequence *entities.SequencedBatches
//   - dbTx entities.Tx
func (_e *Storer_Expecter) AddSequencedBatches(ctx interface{}, sequence interface{}, dbTx interface{}) *Storer_AddSequencedBatches_Call {
	return &Storer_AddSequencedBatches_Call{Call: _e.mock.On("AddSequencedBatches", ctx, sequence, dbTx)}
}

func (_c *Storer_AddSequencedBatches_Call) Run(run func(ctx context.Context, sequence *entities.SequencedBatches, dbTx entities.Tx)) *Storer_AddSequencedBatches_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.SequencedBatches), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_AddSequencedBatches_Call) Return(_a0 error) *Storer_AddSequencedBatches_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storer_AddSequencedBatches_Call) RunAndReturn(run func(context.Context, *entities.SequencedBatches, entities.Tx) error) *Storer_AddSequencedBatches_Call {
	_c.Call.Return(run)
	return _c
}

// AddVirtualBatch provides a mock function with given fields: ctx, virtualBatch, dbTx
func (_m *Storer) AddVirtualBatch(ctx context.Context, virtualBatch *entities.VirtualBatch, dbTx entities.Tx) error {
	ret := _m.Called(ctx, virtualBatch, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for AddVirtualBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.VirtualBatch, entities.Tx) error); ok {
		r0 = rf(ctx, virtualBatch, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storer_AddVirtualBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddVirtualBatch'
type Storer_AddVirtualBatch_Call struct {
	*mock.Call
}

// AddVirtualBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - virtualBatch *entities.VirtualBatch
//   - dbTx entities.Tx
func (_e *Storer_Expecter) AddVirtualBatch(ctx interface{}, virtualBatch interface{}, dbTx interface{}) *Storer_AddVirtualBatch_Call {
	return &Storer_AddVirtualBatch_Call{Call: _e.mock.On("AddVirtualBatch", ctx, virtualBatch, dbTx)}
}

func (_c *Storer_AddVirtualBatch_Call) Run(run func(ctx context.Context, virtualBatch *entities.VirtualBatch, dbTx entities.Tx)) *Storer_AddVirtualBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.VirtualBatch), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_AddVirtualBatch_Call) Return(_a0 error) *Storer_AddVirtualBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storer_AddVirtualBatch_Call) RunAndReturn(run func(context.Context, *entities.VirtualBatch, entities.Tx) error) *Storer_AddVirtualBatch_Call {
	_c.Call.Return(run)
	return _c
}

// BeginTransaction provides a mock function with given fields: ctx
func (_m *Storer) BeginTransaction(ctx context.Context) (entities.Tx, error) {
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

// Storer_BeginTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeginTransaction'
type Storer_BeginTransaction_Call struct {
	*mock.Call
}

// BeginTransaction is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Storer_Expecter) BeginTransaction(ctx interface{}) *Storer_BeginTransaction_Call {
	return &Storer_BeginTransaction_Call{Call: _e.mock.On("BeginTransaction", ctx)}
}

func (_c *Storer_BeginTransaction_Call) Run(run func(ctx context.Context)) *Storer_BeginTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Storer_BeginTransaction_Call) Return(_a0 entities.Tx, _a1 error) *Storer_BeginTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_BeginTransaction_Call) RunAndReturn(run func(context.Context) (entities.Tx, error)) *Storer_BeginTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllL1InfoTreeLeaves provides a mock function with given fields: ctx, dbTx
func (_m *Storer) GetAllL1InfoTreeLeaves(ctx context.Context, dbTx entities.Tx) ([]entities.L1InfoTreeLeaf, error) {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllL1InfoTreeLeaves")
	}

	var r0 []entities.L1InfoTreeLeaf
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Tx) ([]entities.L1InfoTreeLeaf, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.Tx) []entities.L1InfoTreeLeaf); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.L1InfoTreeLeaf)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storer_GetAllL1InfoTreeLeaves_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllL1InfoTreeLeaves'
type Storer_GetAllL1InfoTreeLeaves_Call struct {
	*mock.Call
}

// GetAllL1InfoTreeLeaves is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetAllL1InfoTreeLeaves(ctx interface{}, dbTx interface{}) *Storer_GetAllL1InfoTreeLeaves_Call {
	return &Storer_GetAllL1InfoTreeLeaves_Call{Call: _e.mock.On("GetAllL1InfoTreeLeaves", ctx, dbTx)}
}

func (_c *Storer_GetAllL1InfoTreeLeaves_Call) Run(run func(ctx context.Context, dbTx entities.Tx)) *Storer_GetAllL1InfoTreeLeaves_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetAllL1InfoTreeLeaves_Call) Return(_a0 []entities.L1InfoTreeLeaf, _a1 error) *Storer_GetAllL1InfoTreeLeaves_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetAllL1InfoTreeLeaves_Call) RunAndReturn(run func(context.Context, entities.Tx) ([]entities.L1InfoTreeLeaf, error)) *Storer_GetAllL1InfoTreeLeaves_Call {
	_c.Call.Return(run)
	return _c
}

// GetBlockByNumber provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *Storer) GetBlockByNumber(ctx context.Context, blockNumber uint64, dbTx entities.Tx) (*entities.L1Block, error) {
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

// Storer_GetBlockByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockByNumber'
type Storer_GetBlockByNumber_Call struct {
	*mock.Call
}

// GetBlockByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - blockNumber uint64
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetBlockByNumber(ctx interface{}, blockNumber interface{}, dbTx interface{}) *Storer_GetBlockByNumber_Call {
	return &Storer_GetBlockByNumber_Call{Call: _e.mock.On("GetBlockByNumber", ctx, blockNumber, dbTx)}
}

func (_c *Storer_GetBlockByNumber_Call) Run(run func(ctx context.Context, blockNumber uint64, dbTx entities.Tx)) *Storer_GetBlockByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetBlockByNumber_Call) Return(_a0 *entities.L1Block, _a1 error) *Storer_GetBlockByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetBlockByNumber_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)) *Storer_GetBlockByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetFirstUncheckedBlock provides a mock function with given fields: ctx, fromBlockNumber, dbTx
func (_m *Storer) GetFirstUncheckedBlock(ctx context.Context, fromBlockNumber uint64, dbTx entities.Tx) (*entities.L1Block, error) {
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

// Storer_GetFirstUncheckedBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFirstUncheckedBlock'
type Storer_GetFirstUncheckedBlock_Call struct {
	*mock.Call
}

// GetFirstUncheckedBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - fromBlockNumber uint64
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetFirstUncheckedBlock(ctx interface{}, fromBlockNumber interface{}, dbTx interface{}) *Storer_GetFirstUncheckedBlock_Call {
	return &Storer_GetFirstUncheckedBlock_Call{Call: _e.mock.On("GetFirstUncheckedBlock", ctx, fromBlockNumber, dbTx)}
}

func (_c *Storer_GetFirstUncheckedBlock_Call) Run(run func(ctx context.Context, fromBlockNumber uint64, dbTx entities.Tx)) *Storer_GetFirstUncheckedBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetFirstUncheckedBlock_Call) Return(_a0 *entities.L1Block, _a1 error) *Storer_GetFirstUncheckedBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetFirstUncheckedBlock_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)) *Storer_GetFirstUncheckedBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetForkIDByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *Storer) GetForkIDByBatchNumber(ctx context.Context, batchNumber uint64, dbTx entities.Tx) uint64 {
	ret := _m.Called(ctx, batchNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetForkIDByBatchNumber")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) uint64); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// Storer_GetForkIDByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetForkIDByBatchNumber'
type Storer_GetForkIDByBatchNumber_Call struct {
	*mock.Call
}

// GetForkIDByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetForkIDByBatchNumber(ctx interface{}, batchNumber interface{}, dbTx interface{}) *Storer_GetForkIDByBatchNumber_Call {
	return &Storer_GetForkIDByBatchNumber_Call{Call: _e.mock.On("GetForkIDByBatchNumber", ctx, batchNumber, dbTx)}
}

func (_c *Storer_GetForkIDByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx entities.Tx)) *Storer_GetForkIDByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetForkIDByBatchNumber_Call) Return(_a0 uint64) *Storer_GetForkIDByBatchNumber_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storer_GetForkIDByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) uint64) *Storer_GetForkIDByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetForkIDs provides a mock function with given fields: ctx, dbTx
func (_m *Storer) GetForkIDs(ctx context.Context, dbTx entities.Tx) ([]entities.ForkIDInterval, error) {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetForkIDs")
	}

	var r0 []entities.ForkIDInterval
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Tx) ([]entities.ForkIDInterval, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.Tx) []entities.ForkIDInterval); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.ForkIDInterval)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storer_GetForkIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetForkIDs'
type Storer_GetForkIDs_Call struct {
	*mock.Call
}

// GetForkIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetForkIDs(ctx interface{}, dbTx interface{}) *Storer_GetForkIDs_Call {
	return &Storer_GetForkIDs_Call{Call: _e.mock.On("GetForkIDs", ctx, dbTx)}
}

func (_c *Storer_GetForkIDs_Call) Run(run func(ctx context.Context, dbTx entities.Tx)) *Storer_GetForkIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetForkIDs_Call) Return(_a0 []entities.ForkIDInterval, _a1 error) *Storer_GetForkIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetForkIDs_Call) RunAndReturn(run func(context.Context, entities.Tx) ([]entities.ForkIDInterval, error)) *Storer_GetForkIDs_Call {
	_c.Call.Return(run)
	return _c
}

// GetL1InfoLeafPerIndex provides a mock function with given fields: ctx, L1InfoTreeIndex, dbTx
func (_m *Storer) GetL1InfoLeafPerIndex(ctx context.Context, L1InfoTreeIndex uint32, dbTx entities.Tx) (*entities.L1InfoTreeLeaf, error) {
	ret := _m.Called(ctx, L1InfoTreeIndex, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetL1InfoLeafPerIndex")
	}

	var r0 *entities.L1InfoTreeLeaf
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, entities.Tx) (*entities.L1InfoTreeLeaf, error)); ok {
		return rf(ctx, L1InfoTreeIndex, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, entities.Tx) *entities.L1InfoTreeLeaf); ok {
		r0 = rf(ctx, L1InfoTreeIndex, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.L1InfoTreeLeaf)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, entities.Tx) error); ok {
		r1 = rf(ctx, L1InfoTreeIndex, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storer_GetL1InfoLeafPerIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetL1InfoLeafPerIndex'
type Storer_GetL1InfoLeafPerIndex_Call struct {
	*mock.Call
}

// GetL1InfoLeafPerIndex is a helper method to define mock.On call
//   - ctx context.Context
//   - L1InfoTreeIndex uint32
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetL1InfoLeafPerIndex(ctx interface{}, L1InfoTreeIndex interface{}, dbTx interface{}) *Storer_GetL1InfoLeafPerIndex_Call {
	return &Storer_GetL1InfoLeafPerIndex_Call{Call: _e.mock.On("GetL1InfoLeafPerIndex", ctx, L1InfoTreeIndex, dbTx)}
}

func (_c *Storer_GetL1InfoLeafPerIndex_Call) Run(run func(ctx context.Context, L1InfoTreeIndex uint32, dbTx entities.Tx)) *Storer_GetL1InfoLeafPerIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetL1InfoLeafPerIndex_Call) Return(_a0 *entities.L1InfoTreeLeaf, _a1 error) *Storer_GetL1InfoLeafPerIndex_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetL1InfoLeafPerIndex_Call) RunAndReturn(run func(context.Context, uint32, entities.Tx) (*entities.L1InfoTreeLeaf, error)) *Storer_GetL1InfoLeafPerIndex_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastBlock provides a mock function with given fields: ctx, dbTx
func (_m *Storer) GetLastBlock(ctx context.Context, dbTx entities.Tx) (*entities.L1Block, error) {
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

// Storer_GetLastBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastBlock'
type Storer_GetLastBlock_Call struct {
	*mock.Call
}

// GetLastBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetLastBlock(ctx interface{}, dbTx interface{}) *Storer_GetLastBlock_Call {
	return &Storer_GetLastBlock_Call{Call: _e.mock.On("GetLastBlock", ctx, dbTx)}
}

func (_c *Storer_GetLastBlock_Call) Run(run func(ctx context.Context, dbTx entities.Tx)) *Storer_GetLastBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetLastBlock_Call) Return(_a0 *entities.L1Block, _a1 error) *Storer_GetLastBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetLastBlock_Call) RunAndReturn(run func(context.Context, entities.Tx) (*entities.L1Block, error)) *Storer_GetLastBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestL1InfoTreeLeaf provides a mock function with given fields: ctx, dbTx
func (_m *Storer) GetLatestL1InfoTreeLeaf(ctx context.Context, dbTx entities.Tx) (*entities.L1InfoTreeLeaf, error) {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestL1InfoTreeLeaf")
	}

	var r0 *entities.L1InfoTreeLeaf
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Tx) (*entities.L1InfoTreeLeaf, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.Tx) *entities.L1InfoTreeLeaf); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.L1InfoTreeLeaf)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storer_GetLatestL1InfoTreeLeaf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestL1InfoTreeLeaf'
type Storer_GetLatestL1InfoTreeLeaf_Call struct {
	*mock.Call
}

// GetLatestL1InfoTreeLeaf is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetLatestL1InfoTreeLeaf(ctx interface{}, dbTx interface{}) *Storer_GetLatestL1InfoTreeLeaf_Call {
	return &Storer_GetLatestL1InfoTreeLeaf_Call{Call: _e.mock.On("GetLatestL1InfoTreeLeaf", ctx, dbTx)}
}

func (_c *Storer_GetLatestL1InfoTreeLeaf_Call) Run(run func(ctx context.Context, dbTx entities.Tx)) *Storer_GetLatestL1InfoTreeLeaf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetLatestL1InfoTreeLeaf_Call) Return(_a0 *entities.L1InfoTreeLeaf, _a1 error) *Storer_GetLatestL1InfoTreeLeaf_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetLatestL1InfoTreeLeaf_Call) RunAndReturn(run func(context.Context, entities.Tx) (*entities.L1InfoTreeLeaf, error)) *Storer_GetLatestL1InfoTreeLeaf_Call {
	_c.Call.Return(run)
	return _c
}

// GetLeafsByL1InfoRoot provides a mock function with given fields: ctx, l1InfoRoot, dbTx
func (_m *Storer) GetLeafsByL1InfoRoot(ctx context.Context, l1InfoRoot common.Hash, dbTx entities.Tx) ([]entities.L1InfoTreeLeaf, error) {
	ret := _m.Called(ctx, l1InfoRoot, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetLeafsByL1InfoRoot")
	}

	var r0 []entities.L1InfoTreeLeaf
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, entities.Tx) ([]entities.L1InfoTreeLeaf, error)); ok {
		return rf(ctx, l1InfoRoot, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, entities.Tx) []entities.L1InfoTreeLeaf); ok {
		r0 = rf(ctx, l1InfoRoot, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.L1InfoTreeLeaf)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, entities.Tx) error); ok {
		r1 = rf(ctx, l1InfoRoot, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storer_GetLeafsByL1InfoRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLeafsByL1InfoRoot'
type Storer_GetLeafsByL1InfoRoot_Call struct {
	*mock.Call
}

// GetLeafsByL1InfoRoot is a helper method to define mock.On call
//   - ctx context.Context
//   - l1InfoRoot common.Hash
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetLeafsByL1InfoRoot(ctx interface{}, l1InfoRoot interface{}, dbTx interface{}) *Storer_GetLeafsByL1InfoRoot_Call {
	return &Storer_GetLeafsByL1InfoRoot_Call{Call: _e.mock.On("GetLeafsByL1InfoRoot", ctx, l1InfoRoot, dbTx)}
}

func (_c *Storer_GetLeafsByL1InfoRoot_Call) Run(run func(ctx context.Context, l1InfoRoot common.Hash, dbTx entities.Tx)) *Storer_GetLeafsByL1InfoRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetLeafsByL1InfoRoot_Call) Return(_a0 []entities.L1InfoTreeLeaf, _a1 error) *Storer_GetLeafsByL1InfoRoot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetLeafsByL1InfoRoot_Call) RunAndReturn(run func(context.Context, common.Hash, entities.Tx) ([]entities.L1InfoTreeLeaf, error)) *Storer_GetLeafsByL1InfoRoot_Call {
	_c.Call.Return(run)
	return _c
}

// GetPreviousBlock provides a mock function with given fields: ctx, offset, dbTx
func (_m *Storer) GetPreviousBlock(ctx context.Context, offset uint64, dbTx entities.Tx) (*entities.L1Block, error) {
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

// Storer_GetPreviousBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPreviousBlock'
type Storer_GetPreviousBlock_Call struct {
	*mock.Call
}

// GetPreviousBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - offset uint64
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetPreviousBlock(ctx interface{}, offset interface{}, dbTx interface{}) *Storer_GetPreviousBlock_Call {
	return &Storer_GetPreviousBlock_Call{Call: _e.mock.On("GetPreviousBlock", ctx, offset, dbTx)}
}

func (_c *Storer_GetPreviousBlock_Call) Run(run func(ctx context.Context, offset uint64, dbTx entities.Tx)) *Storer_GetPreviousBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetPreviousBlock_Call) Return(_a0 *entities.L1Block, _a1 error) *Storer_GetPreviousBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetPreviousBlock_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)) *Storer_GetPreviousBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetSequenceByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *Storer) GetSequenceByBatchNumber(ctx context.Context, batchNumber uint64, dbTx entities.Tx) (*entities.SequencedBatches, error) {
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

// Storer_GetSequenceByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSequenceByBatchNumber'
type Storer_GetSequenceByBatchNumber_Call struct {
	*mock.Call
}

// GetSequenceByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetSequenceByBatchNumber(ctx interface{}, batchNumber interface{}, dbTx interface{}) *Storer_GetSequenceByBatchNumber_Call {
	return &Storer_GetSequenceByBatchNumber_Call{Call: _e.mock.On("GetSequenceByBatchNumber", ctx, batchNumber, dbTx)}
}

func (_c *Storer_GetSequenceByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx entities.Tx)) *Storer_GetSequenceByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetSequenceByBatchNumber_Call) Return(_a0 *entities.SequencedBatches, _a1 error) *Storer_GetSequenceByBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetSequenceByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.SequencedBatches, error)) *Storer_GetSequenceByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetUncheckedBlocks provides a mock function with given fields: ctx, fromBlockNumber, toBlockNumber, dbTx
func (_m *Storer) GetUncheckedBlocks(ctx context.Context, fromBlockNumber uint64, toBlockNumber uint64, dbTx entities.Tx) (*[]entities.L1Block, error) {
	ret := _m.Called(ctx, fromBlockNumber, toBlockNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetUncheckedBlocks")
	}

	var r0 *[]entities.L1Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, entities.Tx) (*[]entities.L1Block, error)); ok {
		return rf(ctx, fromBlockNumber, toBlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, entities.Tx) *[]entities.L1Block); ok {
		r0 = rf(ctx, fromBlockNumber, toBlockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.L1Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, entities.Tx) error); ok {
		r1 = rf(ctx, fromBlockNumber, toBlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storer_GetUncheckedBlocks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUncheckedBlocks'
type Storer_GetUncheckedBlocks_Call struct {
	*mock.Call
}

// GetUncheckedBlocks is a helper method to define mock.On call
//   - ctx context.Context
//   - fromBlockNumber uint64
//   - toBlockNumber uint64
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetUncheckedBlocks(ctx interface{}, fromBlockNumber interface{}, toBlockNumber interface{}, dbTx interface{}) *Storer_GetUncheckedBlocks_Call {
	return &Storer_GetUncheckedBlocks_Call{Call: _e.mock.On("GetUncheckedBlocks", ctx, fromBlockNumber, toBlockNumber, dbTx)}
}

func (_c *Storer_GetUncheckedBlocks_Call) Run(run func(ctx context.Context, fromBlockNumber uint64, toBlockNumber uint64, dbTx entities.Tx)) *Storer_GetUncheckedBlocks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(uint64), args[3].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetUncheckedBlocks_Call) Return(_a0 *[]entities.L1Block, _a1 error) *Storer_GetUncheckedBlocks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetUncheckedBlocks_Call) RunAndReturn(run func(context.Context, uint64, uint64, entities.Tx) (*[]entities.L1Block, error)) *Storer_GetUncheckedBlocks_Call {
	_c.Call.Return(run)
	return _c
}

// GetVirtualBatchByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *Storer) GetVirtualBatchByBatchNumber(ctx context.Context, batchNumber uint64, dbTx entities.Tx) (*entities.VirtualBatch, error) {
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

// Storer_GetVirtualBatchByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVirtualBatchByBatchNumber'
type Storer_GetVirtualBatchByBatchNumber_Call struct {
	*mock.Call
}

// GetVirtualBatchByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx entities.Tx
func (_e *Storer_Expecter) GetVirtualBatchByBatchNumber(ctx interface{}, batchNumber interface{}, dbTx interface{}) *Storer_GetVirtualBatchByBatchNumber_Call {
	return &Storer_GetVirtualBatchByBatchNumber_Call{Call: _e.mock.On("GetVirtualBatchByBatchNumber", ctx, batchNumber, dbTx)}
}

func (_c *Storer_GetVirtualBatchByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx entities.Tx)) *Storer_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_GetVirtualBatchByBatchNumber_Call) Return(_a0 *entities.VirtualBatch, _a1 error) *Storer_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storer_GetVirtualBatchByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.VirtualBatch, error)) *Storer_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateForkID provides a mock function with given fields: ctx, forkID, dbTx
func (_m *Storer) UpdateForkID(ctx context.Context, forkID entities.ForkIDInterval, dbTx entities.Tx) error {
	ret := _m.Called(ctx, forkID, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for UpdateForkID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.ForkIDInterval, entities.Tx) error); ok {
		r0 = rf(ctx, forkID, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storer_UpdateForkID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateForkID'
type Storer_UpdateForkID_Call struct {
	*mock.Call
}

// UpdateForkID is a helper method to define mock.On call
//   - ctx context.Context
//   - forkID entities.ForkIDInterval
//   - dbTx entities.Tx
func (_e *Storer_Expecter) UpdateForkID(ctx interface{}, forkID interface{}, dbTx interface{}) *Storer_UpdateForkID_Call {
	return &Storer_UpdateForkID_Call{Call: _e.mock.On("UpdateForkID", ctx, forkID, dbTx)}
}

func (_c *Storer_UpdateForkID_Call) Run(run func(ctx context.Context, forkID entities.ForkIDInterval, dbTx entities.Tx)) *Storer_UpdateForkID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.ForkIDInterval), args[2].(entities.Tx))
	})
	return _c
}

func (_c *Storer_UpdateForkID_Call) Return(_a0 error) *Storer_UpdateForkID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storer_UpdateForkID_Call) RunAndReturn(run func(context.Context, entities.ForkIDInterval, entities.Tx) error) *Storer_UpdateForkID_Call {
	_c.Call.Return(run)
	return _c
}

// NewStorer creates a new instance of Storer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storer {
	mock := &Storer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
