// Code generated by mockery. DO NOT EDIT.

package mock_synchronizer

import (
	context "context"

	synchronizer "github.com/0xPolygonHermez/zkevm-synchronizer-l1/synchronizer"
	mock "github.com/stretchr/testify/mock"
)

// SynchronizerSequencedBatchesQuerier is an autogenerated mock type for the SynchronizerSequencedBatchesQuerier type
type SynchronizerSequencedBatchesQuerier struct {
	mock.Mock
}

type SynchronizerSequencedBatchesQuerier_Expecter struct {
	mock *mock.Mock
}

func (_m *SynchronizerSequencedBatchesQuerier) EXPECT() *SynchronizerSequencedBatchesQuerier_Expecter {
	return &SynchronizerSequencedBatchesQuerier_Expecter{mock: &_m.Mock}
}

// GetSequenceByBatchNumber provides a mock function with given fields: ctx, batchNumber
func (_m *SynchronizerSequencedBatchesQuerier) GetSequenceByBatchNumber(ctx context.Context, batchNumber uint64) (*synchronizer.SequencedBatches, error) {
	ret := _m.Called(ctx, batchNumber)

	if len(ret) == 0 {
		panic("no return value specified for GetSequenceByBatchNumber")
	}

	var r0 *synchronizer.SequencedBatches
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*synchronizer.SequencedBatches, error)); ok {
		return rf(ctx, batchNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *synchronizer.SequencedBatches); ok {
		r0 = rf(ctx, batchNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synchronizer.SequencedBatches)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, batchNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SynchronizerSequencedBatchesQuerier_GetSequenceByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSequenceByBatchNumber'
type SynchronizerSequencedBatchesQuerier_GetSequenceByBatchNumber_Call struct {
	*mock.Call
}

// GetSequenceByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
func (_e *SynchronizerSequencedBatchesQuerier_Expecter) GetSequenceByBatchNumber(ctx interface{}, batchNumber interface{}) *SynchronizerSequencedBatchesQuerier_GetSequenceByBatchNumber_Call {
	return &SynchronizerSequencedBatchesQuerier_GetSequenceByBatchNumber_Call{Call: _e.mock.On("GetSequenceByBatchNumber", ctx, batchNumber)}
}

func (_c *SynchronizerSequencedBatchesQuerier_GetSequenceByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64)) *SynchronizerSequencedBatchesQuerier_GetSequenceByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *SynchronizerSequencedBatchesQuerier_GetSequenceByBatchNumber_Call) Return(_a0 *synchronizer.SequencedBatches, _a1 error) *SynchronizerSequencedBatchesQuerier_GetSequenceByBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SynchronizerSequencedBatchesQuerier_GetSequenceByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64) (*synchronizer.SequencedBatches, error)) *SynchronizerSequencedBatchesQuerier_GetSequenceByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// NewSynchronizerSequencedBatchesQuerier creates a new instance of SynchronizerSequencedBatchesQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSynchronizerSequencedBatchesQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *SynchronizerSequencedBatchesQuerier {
	mock := &SynchronizerSequencedBatchesQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
