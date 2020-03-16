// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// Workflow is an autogenerated mock type for the Workflow type
type Workflow struct {
	mock.Mock
}

type Workflow_HandleAbortedWorkflow struct {
	*mock.Call
}

func (_m Workflow_HandleAbortedWorkflow) Return(_a0 error) *Workflow_HandleAbortedWorkflow {
	return &Workflow_HandleAbortedWorkflow{Call: _m.Call.Return(_a0)}
}

func (_m *Workflow) OnHandleAbortedWorkflow(ctx context.Context, w *v1alpha1.FlyteWorkflow, maxRetries uint32) *Workflow_HandleAbortedWorkflow {
	c := _m.On("HandleAbortedWorkflow", ctx, w, maxRetries)
	return &Workflow_HandleAbortedWorkflow{Call: c}
}

func (_m *Workflow) OnHandleAbortedWorkflowMatch(matchers ...interface{}) *Workflow_HandleAbortedWorkflow {
	c := _m.On("HandleAbortedWorkflow", matchers...)
	return &Workflow_HandleAbortedWorkflow{Call: c}
}

// HandleAbortedWorkflow provides a mock function with given fields: ctx, w, maxRetries
func (_m *Workflow) HandleAbortedWorkflow(ctx context.Context, w *v1alpha1.FlyteWorkflow, maxRetries uint32) error {
	ret := _m.Called(ctx, w, maxRetries)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.FlyteWorkflow, uint32) error); ok {
		r0 = rf(ctx, w, maxRetries)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Workflow_HandleFlyteWorkflow struct {
	*mock.Call
}

func (_m Workflow_HandleFlyteWorkflow) Return(_a0 error) *Workflow_HandleFlyteWorkflow {
	return &Workflow_HandleFlyteWorkflow{Call: _m.Call.Return(_a0)}
}

func (_m *Workflow) OnHandleFlyteWorkflow(ctx context.Context, w *v1alpha1.FlyteWorkflow) *Workflow_HandleFlyteWorkflow {
	c := _m.On("HandleFlyteWorkflow", ctx, w)
	return &Workflow_HandleFlyteWorkflow{Call: c}
}

func (_m *Workflow) OnHandleFlyteWorkflowMatch(matchers ...interface{}) *Workflow_HandleFlyteWorkflow {
	c := _m.On("HandleFlyteWorkflow", matchers...)
	return &Workflow_HandleFlyteWorkflow{Call: c}
}

// HandleFlyteWorkflow provides a mock function with given fields: ctx, w
func (_m *Workflow) HandleFlyteWorkflow(ctx context.Context, w *v1alpha1.FlyteWorkflow) error {
	ret := _m.Called(ctx, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.FlyteWorkflow) error); ok {
		r0 = rf(ctx, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Workflow_Initialize struct {
	*mock.Call
}

func (_m Workflow_Initialize) Return(_a0 error) *Workflow_Initialize {
	return &Workflow_Initialize{Call: _m.Call.Return(_a0)}
}

func (_m *Workflow) OnInitialize(ctx context.Context) *Workflow_Initialize {
	c := _m.On("Initialize", ctx)
	return &Workflow_Initialize{Call: c}
}

func (_m *Workflow) OnInitializeMatch(matchers ...interface{}) *Workflow_Initialize {
	c := _m.On("Initialize", matchers...)
	return &Workflow_Initialize{Call: c}
}

// Initialize provides a mock function with given fields: ctx
func (_m *Workflow) Initialize(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
