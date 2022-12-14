// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// Subscription is an autogenerated mock type for the Subscription type
type Subscription struct {
	mock.Mock
}

// Err provides a mock function with given fields:
func (_m *Subscription) Err() <-chan error {
	ret := _m.Called()

	var r0 <-chan error
	if rf, ok := ret.Get(0).(func() <-chan error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan error)
		}
	}

	return r0
}

// Unsubscribe provides a mock function with given fields:
func (_m *Subscription) Unsubscribe() {
	_m.Called()
}

// NewSubscription creates a new instance of Subscription. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewSubscription(t testing.TB) *Subscription {
	mock := &Subscription{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
