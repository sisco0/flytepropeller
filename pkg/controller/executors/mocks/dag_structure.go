// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// DAGStructure is an autogenerated mock type for the DAGStructure type
type DAGStructure struct {
	mock.Mock
}

type DAGStructure_FromNode struct {
	*mock.Call
}

func (_m DAGStructure_FromNode) Return(_a0 []string, _a1 error) *DAGStructure_FromNode {
	return &DAGStructure_FromNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DAGStructure) OnFromNode(id string) *DAGStructure_FromNode {
	c := _m.On("FromNode", id)
	return &DAGStructure_FromNode{Call: c}
}

func (_m *DAGStructure) OnFromNodeMatch(matchers ...interface{}) *DAGStructure_FromNode {
	c := _m.On("FromNode", matchers...)
	return &DAGStructure_FromNode{Call: c}
}

// FromNode provides a mock function with given fields: id
func (_m *DAGStructure) FromNode(id string) ([]string, error) {
	ret := _m.Called(id)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DAGStructure_StartNode struct {
	*mock.Call
}

func (_m DAGStructure_StartNode) Return(_a0 v1alpha1.ExecutableNode) *DAGStructure_StartNode {
	return &DAGStructure_StartNode{Call: _m.Call.Return(_a0)}
}

func (_m *DAGStructure) OnStartNode() *DAGStructure_StartNode {
	c := _m.On("StartNode")
	return &DAGStructure_StartNode{Call: c}
}

func (_m *DAGStructure) OnStartNodeMatch(matchers ...interface{}) *DAGStructure_StartNode {
	c := _m.On("StartNode", matchers...)
	return &DAGStructure_StartNode{Call: c}
}

// StartNode provides a mock function with given fields:
func (_m *DAGStructure) StartNode() v1alpha1.ExecutableNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	return r0
}

type DAGStructure_ToNode struct {
	*mock.Call
}

func (_m DAGStructure_ToNode) Return(_a0 []string, _a1 error) *DAGStructure_ToNode {
	return &DAGStructure_ToNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DAGStructure) OnToNode(id string) *DAGStructure_ToNode {
	c := _m.On("ToNode", id)
	return &DAGStructure_ToNode{Call: c}
}

func (_m *DAGStructure) OnToNodeMatch(matchers ...interface{}) *DAGStructure_ToNode {
	c := _m.On("ToNode", matchers...)
	return &DAGStructure_ToNode{Call: c}
}

// ToNode provides a mock function with given fields: id
func (_m *DAGStructure) ToNode(id string) ([]string, error) {
	ret := _m.Called(id)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
