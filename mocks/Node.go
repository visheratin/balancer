// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import balancer "github.com/visheratin/balancer"
import mock "github.com/stretchr/testify/mock"

// Node is an autogenerated mock type for the Node type
type Node struct {
	mock.Mock
}

// Capacity provides a mock function with given fields:
func (_m *Node) Capacity() balancer.Capacity {
	ret := _m.Called()

	var r0 balancer.Capacity
	if rf, ok := ret.Get(0).(func() balancer.Capacity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(balancer.Capacity)
		}
	}

	return r0
}

// Hash provides a mock function with given fields:
func (_m *Node) Hash() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Node) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Power provides a mock function with given fields:
func (_m *Node) Power() balancer.Power {
	ret := _m.Called()

	var r0 balancer.Power
	if rf, ok := ret.Get(0).(func() balancer.Power); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(balancer.Power)
		}
	}

	return r0
}
