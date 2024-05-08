// Code generated by mockery. DO NOT EDIT.

package mock_model

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"

	mock "github.com/stretchr/testify/mock"
)

// StorageL1InfoTreeInterface is an autogenerated mock type for the StorageL1InfoTreeInterface type
type StorageL1InfoTreeInterface struct {
	mock.Mock
}

type StorageL1InfoTreeInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *StorageL1InfoTreeInterface) EXPECT() *StorageL1InfoTreeInterface_Expecter {
	return &StorageL1InfoTreeInterface_Expecter{mock: &_m.Mock}
}

// AddL1InfoTreeLeaf provides a mock function with given fields: ctx, exitRoot, dbTx
func (_m *StorageL1InfoTreeInterface) AddL1InfoTreeLeaf(ctx context.Context, exitRoot *entities.L1InfoTreeLeaf, dbTx entities.Tx) error {
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

// StorageL1InfoTreeInterface_AddL1InfoTreeLeaf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddL1InfoTreeLeaf'
type StorageL1InfoTreeInterface_AddL1InfoTreeLeaf_Call struct {
	*mock.Call
}

// AddL1InfoTreeLeaf is a helper method to define mock.On call
//   - ctx context.Context
//   - exitRoot *entities.L1InfoTreeLeaf
//   - dbTx entities.Tx
func (_e *StorageL1InfoTreeInterface_Expecter) AddL1InfoTreeLeaf(ctx interface{}, exitRoot interface{}, dbTx interface{}) *StorageL1InfoTreeInterface_AddL1InfoTreeLeaf_Call {
	return &StorageL1InfoTreeInterface_AddL1InfoTreeLeaf_Call{Call: _e.mock.On("AddL1InfoTreeLeaf", ctx, exitRoot, dbTx)}
}

func (_c *StorageL1InfoTreeInterface_AddL1InfoTreeLeaf_Call) Run(run func(ctx context.Context, exitRoot *entities.L1InfoTreeLeaf, dbTx entities.Tx)) *StorageL1InfoTreeInterface_AddL1InfoTreeLeaf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.L1InfoTreeLeaf), args[2].(entities.Tx))
	})
	return _c
}

func (_c *StorageL1InfoTreeInterface_AddL1InfoTreeLeaf_Call) Return(_a0 error) *StorageL1InfoTreeInterface_AddL1InfoTreeLeaf_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageL1InfoTreeInterface_AddL1InfoTreeLeaf_Call) RunAndReturn(run func(context.Context, *entities.L1InfoTreeLeaf, entities.Tx) error) *StorageL1InfoTreeInterface_AddL1InfoTreeLeaf_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllL1InfoTreeLeaves provides a mock function with given fields: ctx, dbTx
func (_m *StorageL1InfoTreeInterface) GetAllL1InfoTreeLeaves(ctx context.Context, dbTx entities.Tx) ([]entities.L1InfoTreeLeaf, error) {
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

// StorageL1InfoTreeInterface_GetAllL1InfoTreeLeaves_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllL1InfoTreeLeaves'
type StorageL1InfoTreeInterface_GetAllL1InfoTreeLeaves_Call struct {
	*mock.Call
}

// GetAllL1InfoTreeLeaves is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx entities.Tx
func (_e *StorageL1InfoTreeInterface_Expecter) GetAllL1InfoTreeLeaves(ctx interface{}, dbTx interface{}) *StorageL1InfoTreeInterface_GetAllL1InfoTreeLeaves_Call {
	return &StorageL1InfoTreeInterface_GetAllL1InfoTreeLeaves_Call{Call: _e.mock.On("GetAllL1InfoTreeLeaves", ctx, dbTx)}
}

func (_c *StorageL1InfoTreeInterface_GetAllL1InfoTreeLeaves_Call) Run(run func(ctx context.Context, dbTx entities.Tx)) *StorageL1InfoTreeInterface_GetAllL1InfoTreeLeaves_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.Tx))
	})
	return _c
}

func (_c *StorageL1InfoTreeInterface_GetAllL1InfoTreeLeaves_Call) Return(_a0 []entities.L1InfoTreeLeaf, _a1 error) *StorageL1InfoTreeInterface_GetAllL1InfoTreeLeaves_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageL1InfoTreeInterface_GetAllL1InfoTreeLeaves_Call) RunAndReturn(run func(context.Context, entities.Tx) ([]entities.L1InfoTreeLeaf, error)) *StorageL1InfoTreeInterface_GetAllL1InfoTreeLeaves_Call {
	_c.Call.Return(run)
	return _c
}

// GetL1InfoLeafPerIndex provides a mock function with given fields: ctx, L1InfoTreeIndex, dbTx
func (_m *StorageL1InfoTreeInterface) GetL1InfoLeafPerIndex(ctx context.Context, L1InfoTreeIndex uint32, dbTx entities.Tx) (*entities.L1InfoTreeLeaf, error) {
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

// StorageL1InfoTreeInterface_GetL1InfoLeafPerIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetL1InfoLeafPerIndex'
type StorageL1InfoTreeInterface_GetL1InfoLeafPerIndex_Call struct {
	*mock.Call
}

// GetL1InfoLeafPerIndex is a helper method to define mock.On call
//   - ctx context.Context
//   - L1InfoTreeIndex uint32
//   - dbTx entities.Tx
func (_e *StorageL1InfoTreeInterface_Expecter) GetL1InfoLeafPerIndex(ctx interface{}, L1InfoTreeIndex interface{}, dbTx interface{}) *StorageL1InfoTreeInterface_GetL1InfoLeafPerIndex_Call {
	return &StorageL1InfoTreeInterface_GetL1InfoLeafPerIndex_Call{Call: _e.mock.On("GetL1InfoLeafPerIndex", ctx, L1InfoTreeIndex, dbTx)}
}

func (_c *StorageL1InfoTreeInterface_GetL1InfoLeafPerIndex_Call) Run(run func(ctx context.Context, L1InfoTreeIndex uint32, dbTx entities.Tx)) *StorageL1InfoTreeInterface_GetL1InfoLeafPerIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(entities.Tx))
	})
	return _c
}

func (_c *StorageL1InfoTreeInterface_GetL1InfoLeafPerIndex_Call) Return(_a0 *entities.L1InfoTreeLeaf, _a1 error) *StorageL1InfoTreeInterface_GetL1InfoLeafPerIndex_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageL1InfoTreeInterface_GetL1InfoLeafPerIndex_Call) RunAndReturn(run func(context.Context, uint32, entities.Tx) (*entities.L1InfoTreeLeaf, error)) *StorageL1InfoTreeInterface_GetL1InfoLeafPerIndex_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestL1InfoTreeLeaf provides a mock function with given fields: ctx, dbTx
func (_m *StorageL1InfoTreeInterface) GetLatestL1InfoTreeLeaf(ctx context.Context, dbTx entities.Tx) (*entities.L1InfoTreeLeaf, error) {
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

// StorageL1InfoTreeInterface_GetLatestL1InfoTreeLeaf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestL1InfoTreeLeaf'
type StorageL1InfoTreeInterface_GetLatestL1InfoTreeLeaf_Call struct {
	*mock.Call
}

// GetLatestL1InfoTreeLeaf is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx entities.Tx
func (_e *StorageL1InfoTreeInterface_Expecter) GetLatestL1InfoTreeLeaf(ctx interface{}, dbTx interface{}) *StorageL1InfoTreeInterface_GetLatestL1InfoTreeLeaf_Call {
	return &StorageL1InfoTreeInterface_GetLatestL1InfoTreeLeaf_Call{Call: _e.mock.On("GetLatestL1InfoTreeLeaf", ctx, dbTx)}
}

func (_c *StorageL1InfoTreeInterface_GetLatestL1InfoTreeLeaf_Call) Run(run func(ctx context.Context, dbTx entities.Tx)) *StorageL1InfoTreeInterface_GetLatestL1InfoTreeLeaf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.Tx))
	})
	return _c
}

func (_c *StorageL1InfoTreeInterface_GetLatestL1InfoTreeLeaf_Call) Return(_a0 *entities.L1InfoTreeLeaf, _a1 error) *StorageL1InfoTreeInterface_GetLatestL1InfoTreeLeaf_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageL1InfoTreeInterface_GetLatestL1InfoTreeLeaf_Call) RunAndReturn(run func(context.Context, entities.Tx) (*entities.L1InfoTreeLeaf, error)) *StorageL1InfoTreeInterface_GetLatestL1InfoTreeLeaf_Call {
	_c.Call.Return(run)
	return _c
}

// GetLeafsByL1InfoRoot provides a mock function with given fields: ctx, l1InfoRoot, dbTx
func (_m *StorageL1InfoTreeInterface) GetLeafsByL1InfoRoot(ctx context.Context, l1InfoRoot common.Hash, dbTx entities.Tx) ([]entities.L1InfoTreeLeaf, error) {
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

// StorageL1InfoTreeInterface_GetLeafsByL1InfoRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLeafsByL1InfoRoot'
type StorageL1InfoTreeInterface_GetLeafsByL1InfoRoot_Call struct {
	*mock.Call
}

// GetLeafsByL1InfoRoot is a helper method to define mock.On call
//   - ctx context.Context
//   - l1InfoRoot common.Hash
//   - dbTx entities.Tx
func (_e *StorageL1InfoTreeInterface_Expecter) GetLeafsByL1InfoRoot(ctx interface{}, l1InfoRoot interface{}, dbTx interface{}) *StorageL1InfoTreeInterface_GetLeafsByL1InfoRoot_Call {
	return &StorageL1InfoTreeInterface_GetLeafsByL1InfoRoot_Call{Call: _e.mock.On("GetLeafsByL1InfoRoot", ctx, l1InfoRoot, dbTx)}
}

func (_c *StorageL1InfoTreeInterface_GetLeafsByL1InfoRoot_Call) Run(run func(ctx context.Context, l1InfoRoot common.Hash, dbTx entities.Tx)) *StorageL1InfoTreeInterface_GetLeafsByL1InfoRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash), args[2].(entities.Tx))
	})
	return _c
}

func (_c *StorageL1InfoTreeInterface_GetLeafsByL1InfoRoot_Call) Return(_a0 []entities.L1InfoTreeLeaf, _a1 error) *StorageL1InfoTreeInterface_GetLeafsByL1InfoRoot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StorageL1InfoTreeInterface_GetLeafsByL1InfoRoot_Call) RunAndReturn(run func(context.Context, common.Hash, entities.Tx) ([]entities.L1InfoTreeLeaf, error)) *StorageL1InfoTreeInterface_GetLeafsByL1InfoRoot_Call {
	_c.Call.Return(run)
	return _c
}

// NewStorageL1InfoTreeInterface creates a new instance of StorageL1InfoTreeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorageL1InfoTreeInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *StorageL1InfoTreeInterface {
	mock := &StorageL1InfoTreeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
