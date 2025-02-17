// Code generated by mockery. DO NOT EDIT.

package mock_storage

import (
	context "context"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	mock "github.com/stretchr/testify/mock"
)

// sequencedBatchStorer is an autogenerated mock type for the sequencedBatchStorer type
type sequencedBatchStorer struct {
	mock.Mock
}

type sequencedBatchStorer_Expecter struct {
	mock *mock.Mock
}

func (_m *sequencedBatchStorer) EXPECT() *sequencedBatchStorer_Expecter {
	return &sequencedBatchStorer_Expecter{mock: &_m.Mock}
}

// AddSequencedBatches provides a mock function with given fields: ctx, sequence, dbTx
func (_m *sequencedBatchStorer) AddSequencedBatches(ctx context.Context, sequence *entities.SequencedBatches, dbTx entities.Tx) error {
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

// sequencedBatchStorer_AddSequencedBatches_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSequencedBatches'
type sequencedBatchStorer_AddSequencedBatches_Call struct {
	*mock.Call
}

// AddSequencedBatches is a helper method to define mock.On call
//   - ctx context.Context
//   - sequence *entities.SequencedBatches
//   - dbTx entities.Tx
func (_e *sequencedBatchStorer_Expecter) AddSequencedBatches(ctx interface{}, sequence interface{}, dbTx interface{}) *sequencedBatchStorer_AddSequencedBatches_Call {
	return &sequencedBatchStorer_AddSequencedBatches_Call{Call: _e.mock.On("AddSequencedBatches", ctx, sequence, dbTx)}
}

func (_c *sequencedBatchStorer_AddSequencedBatches_Call) Run(run func(ctx context.Context, sequence *entities.SequencedBatches, dbTx entities.Tx)) *sequencedBatchStorer_AddSequencedBatches_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.SequencedBatches), args[2].(entities.Tx))
	})
	return _c
}

func (_c *sequencedBatchStorer_AddSequencedBatches_Call) Return(_a0 error) *sequencedBatchStorer_AddSequencedBatches_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *sequencedBatchStorer_AddSequencedBatches_Call) RunAndReturn(run func(context.Context, *entities.SequencedBatches, entities.Tx) error) *sequencedBatchStorer_AddSequencedBatches_Call {
	_c.Call.Return(run)
	return _c
}

// GetSequenceByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *sequencedBatchStorer) GetSequenceByBatchNumber(ctx context.Context, batchNumber uint64, dbTx entities.Tx) (*entities.SequencedBatches, error) {
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

// sequencedBatchStorer_GetSequenceByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSequenceByBatchNumber'
type sequencedBatchStorer_GetSequenceByBatchNumber_Call struct {
	*mock.Call
}

// GetSequenceByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
//   - dbTx entities.Tx
func (_e *sequencedBatchStorer_Expecter) GetSequenceByBatchNumber(ctx interface{}, batchNumber interface{}, dbTx interface{}) *sequencedBatchStorer_GetSequenceByBatchNumber_Call {
	return &sequencedBatchStorer_GetSequenceByBatchNumber_Call{Call: _e.mock.On("GetSequenceByBatchNumber", ctx, batchNumber, dbTx)}
}

func (_c *sequencedBatchStorer_GetSequenceByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64, dbTx entities.Tx)) *sequencedBatchStorer_GetSequenceByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *sequencedBatchStorer_GetSequenceByBatchNumber_Call) Return(_a0 *entities.SequencedBatches, _a1 error) *sequencedBatchStorer_GetSequenceByBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *sequencedBatchStorer_GetSequenceByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.SequencedBatches, error)) *sequencedBatchStorer_GetSequenceByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// newSequencedBatchStorer creates a new instance of sequencedBatchStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newSequencedBatchStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *sequencedBatchStorer {
	mock := &sequencedBatchStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
