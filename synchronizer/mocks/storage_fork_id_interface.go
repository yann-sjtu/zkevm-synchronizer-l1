// Code generated by mockery. DO NOT EDIT.

package mock_synchronizer

import (
	context "context"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	mock "github.com/stretchr/testify/mock"
)

// StorageForkIDInterface is an autogenerated mock type for the StorageForkIDInterface type
type StorageForkIDInterface struct {
	mock.Mock
}

type StorageForkIDInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *StorageForkIDInterface) EXPECT() *StorageForkIDInterface_Expecter {
	return &StorageForkIDInterface_Expecter{mock: &_m.Mock}
}

// AddForkID provides a mock function with given fields: ctx, forkID, dbTx
func (_m *StorageForkIDInterface) AddForkID(ctx context.Context, forkID entities.ForkIDInterval, dbTx entities.Tx) error {
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

// StorageForkIDInterface_AddForkID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddForkID'
type StorageForkIDInterface_AddForkID_Call struct {
	*mock.Call
}

// AddForkID is a helper method to define mock.On call
//   - ctx context.Context
//   - forkID entities.ForkIDInterval
//   - dbTx entities.Tx
func (_e *StorageForkIDInterface_Expecter) AddForkID(ctx interface{}, forkID interface{}, dbTx interface{}) *StorageForkIDInterface_AddForkID_Call {
	return &StorageForkIDInterface_AddForkID_Call{Call: _e.mock.On("AddForkID", ctx, forkID, dbTx)}
}

func (_c *StorageForkIDInterface_AddForkID_Call) Run(run func(ctx context.Context, forkID entities.ForkIDInterval, dbTx entities.Tx)) *StorageForkIDInterface_AddForkID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.ForkIDInterval), args[2].(entities.Tx))
	})
	return _c
}

func (_c *StorageForkIDInterface_AddForkID_Call) Return(_a0 error) *StorageForkIDInterface_AddForkID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageForkIDInterface_AddForkID_Call) RunAndReturn(run func(context.Context, entities.ForkIDInterval, entities.Tx) error) *StorageForkIDInterface_AddForkID_Call {
	_c.Call.Return(run)
	return _c
}

// GetForkIDs provides a mock function with given fields: ctx, dbTx
func (_m *StorageForkIDInterface) GetForkIDs(ctx context.Context, dbTx entities.Tx) ([]entities.ForkIDInterval, error) {
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

// StorageForkIDInterface_GetForkIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetForkIDs'
type StorageForkIDInterface_GetForkIDs_Call struct {
	*mock.Call
}

// GetForkIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx entities.Tx
func (_e *StorageForkIDInterface_Expecter) GetForkIDs(ctx interface{}, dbTx interface{}) *StorageForkIDInterface_GetForkIDs_Call {
	return &StorageForkIDInterface_GetForkIDs_Call{Call: _e.mock.On("GetForkIDs", ctx, dbTx)}
}

func (_c *StorageForkIDInterface_GetForkIDs_Call) Run(run func(ctx context.Context, dbTx entities.Tx)) *StorageForkIDInterface_GetForkIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.Tx))
	})
	return _c
}

func (_c *StorageForkIDInterface_GetForkIDs_Call) Return(_a0 []entities.ForkIDInterval, _a1 error) *StorageForkIDInterface_GetForkIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageForkIDInterface_GetForkIDs_Call) RunAndReturn(run func(context.Context, entities.Tx) ([]entities.ForkIDInterval, error)) *StorageForkIDInterface_GetForkIDs_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateForkID provides a mock function with given fields: ctx, forkID, dbTx
func (_m *StorageForkIDInterface) UpdateForkID(ctx context.Context, forkID entities.ForkIDInterval, dbTx entities.Tx) error {
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

// StorageForkIDInterface_UpdateForkID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateForkID'
type StorageForkIDInterface_UpdateForkID_Call struct {
	*mock.Call
}

// UpdateForkID is a helper method to define mock.On call
//   - ctx context.Context
//   - forkID entities.ForkIDInterval
//   - dbTx entities.Tx
func (_e *StorageForkIDInterface_Expecter) UpdateForkID(ctx interface{}, forkID interface{}, dbTx interface{}) *StorageForkIDInterface_UpdateForkID_Call {
	return &StorageForkIDInterface_UpdateForkID_Call{Call: _e.mock.On("UpdateForkID", ctx, forkID, dbTx)}
}

func (_c *StorageForkIDInterface_UpdateForkID_Call) Run(run func(ctx context.Context, forkID entities.ForkIDInterval, dbTx entities.Tx)) *StorageForkIDInterface_UpdateForkID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.ForkIDInterval), args[2].(entities.Tx))
	})
	return _c
}

func (_c *StorageForkIDInterface_UpdateForkID_Call) Return(_a0 error) *StorageForkIDInterface_UpdateForkID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageForkIDInterface_UpdateForkID_Call) RunAndReturn(run func(context.Context, entities.ForkIDInterval, entities.Tx) error) *StorageForkIDInterface_UpdateForkID_Call {
	_c.Call.Return(run)
	return _c
}

// NewStorageForkIDInterface creates a new instance of StorageForkIDInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorageForkIDInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *StorageForkIDInterface {
	mock := &StorageForkIDInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
