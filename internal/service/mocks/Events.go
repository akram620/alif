// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	errors "github.com/akram620/alif/pkg/errors"
	mock "github.com/stretchr/testify/mock"

	models "github.com/akram620/alif/internal/models"
)

// Events is an autogenerated mock type for the Events type
type Events struct {
	mock.Mock
}

// CreateEvent provides a mock function with given fields: e
func (_m *Events) CreateEvent(e *models.Event) *errors.ExportableError {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for CreateEvent")
	}

	var r0 *errors.ExportableError
	if rf, ok := ret.Get(0).(func(*models.Event) *errors.ExportableError); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.ExportableError)
		}
	}

	return r0
}

// NewEvents creates a new instance of Events. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEvents(t interface {
	mock.TestingT
	Cleanup(func())
}) *Events {
	mock := &Events{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}