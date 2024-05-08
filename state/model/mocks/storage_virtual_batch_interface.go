// Code generated by mockery. DO NOT EDIT.

package mock_model

import (
	context "context"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	mock "github.com/stretchr/testify/mock"
)

// StorageVirtualBatchInterface is an autogenerated mock type for the StorageVirtualBatchInterface type
type StorageVirtualBatchInterface struct {
	mock.Mock
}

type StorageVirtualBatchInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *StorageVirtualBatchInterface) EXPECT() *StorageVirtualBatchInterface_Expecter {
	return &StorageVirtualBatchInterface_Expecter{mock: &_m.Mock}
}

// AddSequencedBatches provides a mock function with given fields: ctx, sequence, dbTx
func (_m *StorageVirtualBatchInterface) AddSequencedBatches(ctx context.Context, sequence *entities.SequencedBatches, dbTx entities.Tx) error {
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

// StorageVirtualBatchInterface_AddSequencedBatches_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSequencedBatches'
type StorageVirtualBatchInterface_AddSequencedBatches_Call struct {
	*mock.Call
}

// AddSequencedBatches is a helper method to define mock.On call
//   - ctx context.Context
//   - sequence *entities.SequencedBatches
//   - dbTx entities.Tx
func (_e *StorageVirtualBatchInterface_Expecter) AddSequencedBatches(ctx interface{}, sequence interface{}, dbTx interface{}) *StorageVirtualBatchInterface_AddSequencedBatches_Call {
	return &StorageVirtualBatchInterface_AddSequencedBatches_Call{Call: _e.mock.On("AddSequencedBatches", ctx, sequence, dbTx)}
}

func (_c *StorageVirtualBatchInterface_AddSequencedBatches_Call) Run(run func(ctx context.Context, sequence *entities.SequencedBatches, dbTx entities.Tx)) *StorageVirtualBatchInterface_AddSequencedBatches_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.SequencedBatches), args[2].(entities.Tx))
	})
	return _c
}

func (_c *StorageVirtualBatchInterface_AddSequencedBatches_Call) Return(_a0 error) *StorageVirtualBatchInterface_AddSequencedBatches_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageVirtualBatchInterface_AddSequencedBatches_Call) RunAndReturn(run func(context.Context, *entities.SequencedBatches, entities.Tx) error) *StorageVirtualBatchInterface_AddSequencedBatches_Call {
	_c.Call.Return(run)
	return _c
}

// AddVirtualBatch provides a mock function with given fields: ctx, virtualBatch, dbTx
func (_m *StorageVirtualBatchInterface) AddVirtualBatch(ctx context.Context, virtualBatch *entities.VirtualBatch, dbTx entities.Tx) error {
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

// StorageVirtualBatchInterface_AddVirtualBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddVirtualBatch'
type StorageVirtualBatchInterface_AddVirtualBatch_Call struct {
	*mock.Call
}

// AddVirtualBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - virtualBatch *entities.VirtualBatch
//   - dbTx entities.Tx
func (_e *StorageVirtualBatchInterface_Expecter) AddVirtualBatch(ctx interface{}, virtualBatch interface{}, dbTx interface{}) *StorageVirtualBatchInterface_AddVirtualBatch_Call {
	return &StorageVirtualBatchInterface_AddVirtualBatch_Call{Call: _e.mock.On("AddVirtualBatch", ctx, virtualBatch, dbTx)}
}

func (_c *StorageVirtualBatchInterface_AddVirtualBatch_Call) Run(run func(ctx context.Context, virtualBatch *entities.VirtualBatch, dbTx entities.Tx)) *StorageVirtualBatchInterface_AddVirtualBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.VirtualBatch), args[2].(entities.Tx))
	})
	return _c
}

func (_c *StorageVirtualBatchInterface_AddVirtualBatch_Call) Return(_a0 error) *StorageVirtualBatchInterface_AddVirtualBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageVirtualBatchInterface_AddVirtualBatch_Call) RunAndReturn(run func(context.Context, *entities.VirtualBatch, entities.Tx) error) *StorageVirtualBatchInterface_AddVirtualBatch_Call {
	_c.Call.Return(run)
	return _c
}

// GetSequenceByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StorageVirtualBatchInterface) GetSequenceByBatchNumber(ctx context.Context, batchNumber uint64, dbTx entities.Tx) (*entities.SequencedBatches, error) {
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

// StorageVirtualBatchInterface_GetSequenceByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSequenceByBatchNumber'
type StorageVirtualBatchInterface_GetSequenceByBatchNumber_Call struct {
	*mock.Call
}

// GetSequenceByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx entities.Tx
func (_e *StorageVirtualBatchInterface_Expecter) GetSequenceByBatchNumber(ctx interface{}, batchNumber interface{}, dbTx interface{}) *StorageVirtualBatchInterface_GetSequenceByBatchNumber_Call {
	return &StorageVirtualBatchInterface_GetSequenceByBatchNumber_Call{Call: _e.mock.On("GetSequenceByBatchNumber", ctx, batchNumber, dbTx)}
}

func (_c *StorageVirtualBatchInterface_GetSequenceByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx entities.Tx)) *StorageVirtualBatchInterface_GetSequenceByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *StorageVirtualBatchInterface_GetSequenceByBatchNumber_Call) Return(_a0 *entities.SequencedBatches, _a1 error) *StorageVirtualBatchInterface_GetSequenceByBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageVirtualBatchInterface_GetSequenceByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.SequencedBatches, error)) *StorageVirtualBatchInterface_GetSequenceByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetVirtualBatchByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StorageVirtualBatchInterface) GetVirtualBatchByBatchNumber(ctx context.Context, batchNumber uint64, dbTx entities.Tx) (*entities.VirtualBatch, error) {
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

// StorageVirtualBatchInterface_GetVirtualBatchByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVirtualBatchByBatchNumber'
type StorageVirtualBatchInterface_GetVirtualBatchByBatchNumber_Call struct {
	*mock.Call
}

// GetVirtualBatchByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx entities.Tx
func (_e *StorageVirtualBatchInterface_Expecter) GetVirtualBatchByBatchNumber(ctx interface{}, batchNumber interface{}, dbTx interface{}) *StorageVirtualBatchInterface_GetVirtualBatchByBatchNumber_Call {
	return &StorageVirtualBatchInterface_GetVirtualBatchByBatchNumber_Call{Call: _e.mock.On("GetVirtualBatchByBatchNumber", ctx, batchNumber, dbTx)}
}

func (_c *StorageVirtualBatchInterface_GetVirtualBatchByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx entities.Tx)) *StorageVirtualBatchInterface_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *StorageVirtualBatchInterface_GetVirtualBatchByBatchNumber_Call) Return(_a0 *entities.VirtualBatch, _a1 error) *StorageVirtualBatchInterface_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageVirtualBatchInterface_GetVirtualBatchByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.VirtualBatch, error)) *StorageVirtualBatchInterface_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// NewStorageVirtualBatchInterface creates a new instance of StorageVirtualBatchInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorageVirtualBatchInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *StorageVirtualBatchInterface {
	mock := &StorageVirtualBatchInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
