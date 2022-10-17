// Code generated by mockery v2.12.2. DO NOT EDIT.

package commandmocks

import (
	command "github.com/challenge-mercadolibre-cl/api-shortener-url/app/shared/application/command"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Command is an autogenerated mock type for the Command type
type Command struct {
	mock.Mock
}

// Type provides a mock function with given fields:
func (_m *Command) Type() command.Type {
	ret := _m.Called()

	var r0 command.Type
	if rf, ok := ret.Get(0).(func() command.Type); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(command.Type)
	}

	return r0
}

// NewCommand creates a new instance of Command. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommand(t testing.TB) *Command {
	mock := &Command{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
