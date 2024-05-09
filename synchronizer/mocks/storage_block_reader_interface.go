// Code generated by mockery. DO NOT EDIT.

package mock_synchronizer

import (
	context "context"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	mock "github.com/stretchr/testify/mock"
)

// StorageBlockReaderInterface is an autogenerated mock type for the StorageBlockReaderInterface type
type StorageBlockReaderInterface struct {
	mock.Mock
}

type StorageBlockReaderInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *StorageBlockReaderInterface) EXPECT() *StorageBlockReaderInterface_Expecter {
	return &StorageBlockReaderInterface_Expecter{mock: &_m.Mock}
}

// AddBlock provides a mock function with given fields: ctx, block, dbTx
func (_m *StorageBlockReaderInterface) AddBlock(ctx context.Context, block *entities.L1Block, dbTx entities.Tx) error {
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

// StorageBlockReaderInterface_AddBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBlock'
type StorageBlockReaderInterface_AddBlock_Call struct {
	*mock.Call
}

// AddBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - block *entities.L1Block
//   - dbTx entities.Tx
func (_e *StorageBlockReaderInterface_Expecter) AddBlock(ctx interface{}, block interface{}, dbTx interface{}) *StorageBlockReaderInterface_AddBlock_Call {
	return &StorageBlockReaderInterface_AddBlock_Call{Call: _e.mock.On("AddBlock", ctx, block, dbTx)}
}

func (_c *StorageBlockReaderInterface_AddBlock_Call) Run(run func(ctx context.Context, block *entities.L1Block, dbTx entities.Tx)) *StorageBlockReaderInterface_AddBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.L1Block), args[2].(entities.Tx))
	})
	return _c
}

func (_c *StorageBlockReaderInterface_AddBlock_Call) Return(_a0 error) *StorageBlockReaderInterface_AddBlock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageBlockReaderInterface_AddBlock_Call) RunAndReturn(run func(context.Context, *entities.L1Block, entities.Tx) error) *StorageBlockReaderInterface_AddBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetFirstUncheckedBlock provides a mock function with given fields: ctx, fromBlockNumber, dbTx
func (_m *StorageBlockReaderInterface) GetFirstUncheckedBlock(ctx context.Context, fromBlockNumber uint64, dbTx entities.Tx) (*entities.L1Block, error) {
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

// StorageBlockReaderInterface_GetFirstUncheckedBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFirstUncheckedBlock'
type StorageBlockReaderInterface_GetFirstUncheckedBlock_Call struct {
	*mock.Call
}

// GetFirstUncheckedBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - fromBlockNumber uint64
//   - dbTx entities.Tx
func (_e *StorageBlockReaderInterface_Expecter) GetFirstUncheckedBlock(ctx interface{}, fromBlockNumber interface{}, dbTx interface{}) *StorageBlockReaderInterface_GetFirstUncheckedBlock_Call {
	return &StorageBlockReaderInterface_GetFirstUncheckedBlock_Call{Call: _e.mock.On("GetFirstUncheckedBlock", ctx, fromBlockNumber, dbTx)}
}

func (_c *StorageBlockReaderInterface_GetFirstUncheckedBlock_Call) Run(run func(ctx context.Context, fromBlockNumber uint64, dbTx entities.Tx)) *StorageBlockReaderInterface_GetFirstUncheckedBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *StorageBlockReaderInterface_GetFirstUncheckedBlock_Call) Return(_a0 *entities.L1Block, _a1 error) *StorageBlockReaderInterface_GetFirstUncheckedBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageBlockReaderInterface_GetFirstUncheckedBlock_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) (*entities.L1Block, error)) *StorageBlockReaderInterface_GetFirstUncheckedBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastBlock provides a mock function with given fields: ctx, dbTx
func (_m *StorageBlockReaderInterface) GetLastBlock(ctx context.Context, dbTx entities.Tx) (*entities.L1Block, error) {
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

// StorageBlockReaderInterface_GetLastBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastBlock'
type StorageBlockReaderInterface_GetLastBlock_Call struct {
	*mock.Call
}

// GetLastBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx entities.Tx
func (_e *StorageBlockReaderInterface_Expecter) GetLastBlock(ctx interface{}, dbTx interface{}) *StorageBlockReaderInterface_GetLastBlock_Call {
	return &StorageBlockReaderInterface_GetLastBlock_Call{Call: _e.mock.On("GetLastBlock", ctx, dbTx)}
}

func (_c *StorageBlockReaderInterface_GetLastBlock_Call) Run(run func(ctx context.Context, dbTx entities.Tx)) *StorageBlockReaderInterface_GetLastBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.Tx))
	})
	return _c
}

func (_c *StorageBlockReaderInterface_GetLastBlock_Call) Return(_a0 *entities.L1Block, _a1 error) *StorageBlockReaderInterface_GetLastBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageBlockReaderInterface_GetLastBlock_Call) RunAndReturn(run func(context.Context, entities.Tx) (*entities.L1Block, error)) *StorageBlockReaderInterface_GetLastBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetPreviousBlock provides a mock function with given fields: ctx, offset, fromBlockNumber, dbTx
func (_m *StorageBlockReaderInterface) GetPreviousBlock(ctx context.Context, offset uint64, fromBlockNumber *uint64, dbTx entities.Tx) (*entities.L1Block, error) {
	ret := _m.Called(ctx, offset, fromBlockNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetPreviousBlock")
	}

	var r0 *entities.L1Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *uint64, entities.Tx) (*entities.L1Block, error)); ok {
		return rf(ctx, offset, fromBlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *uint64, entities.Tx) *entities.L1Block); ok {
		r0 = rf(ctx, offset, fromBlockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.L1Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, *uint64, entities.Tx) error); ok {
		r1 = rf(ctx, offset, fromBlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StorageBlockReaderInterface_GetPreviousBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPreviousBlock'
type StorageBlockReaderInterface_GetPreviousBlock_Call struct {
	*mock.Call
}

// GetPreviousBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - offset uint64
//   - fromBlockNumber *uint64
//   - dbTx entities.Tx
func (_e *StorageBlockReaderInterface_Expecter) GetPreviousBlock(ctx interface{}, offset interface{}, fromBlockNumber interface{}, dbTx interface{}) *StorageBlockReaderInterface_GetPreviousBlock_Call {
	return &StorageBlockReaderInterface_GetPreviousBlock_Call{Call: _e.mock.On("GetPreviousBlock", ctx, offset, fromBlockNumber, dbTx)}
}

func (_c *StorageBlockReaderInterface_GetPreviousBlock_Call) Run(run func(ctx context.Context, offset uint64, fromBlockNumber *uint64, dbTx entities.Tx)) *StorageBlockReaderInterface_GetPreviousBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(*uint64), args[3].(entities.Tx))
	})
	return _c
}

func (_c *StorageBlockReaderInterface_GetPreviousBlock_Call) Return(_a0 *entities.L1Block, _a1 error) *StorageBlockReaderInterface_GetPreviousBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageBlockReaderInterface_GetPreviousBlock_Call) RunAndReturn(run func(context.Context, uint64, *uint64, entities.Tx) (*entities.L1Block, error)) *StorageBlockReaderInterface_GetPreviousBlock_Call {
	_c.Call.Return(run)
	return _c
}

// NewStorageBlockReaderInterface creates a new instance of StorageBlockReaderInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorageBlockReaderInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *StorageBlockReaderInterface {
	mock := &StorageBlockReaderInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
