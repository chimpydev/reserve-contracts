// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/core/chains/evm/headtracker/types"
)

// HeadListener is an autogenerated mock type for the HeadListener type
type HeadListener struct {
	mock.Mock
}

// Connected provides a mock function with given fields:
func (_m *HeadListener) Connected() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ListenForNewHeads provides a mock function with given fields: handleNewHead, done
func (_m *HeadListener) ListenForNewHeads(handleNewHead types.NewHeadHandler, done func()) {
	_m.Called(handleNewHead, done)
}

// ReceivingHeads provides a mock function with given fields:
func (_m *HeadListener) ReceivingHeads() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewHeadListener creates a new instance of HeadListener. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewHeadListener(t testing.TB) *HeadListener {
	mock := &HeadListener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
