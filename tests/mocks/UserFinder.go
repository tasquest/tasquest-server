// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	commons "tasquest.com/server/commons"

	security "tasquest.com/server/application/security"

	uuid "github.com/google/uuid"
)

// UserFinder is an autogenerated mock type for the UserFinder type
type UserFinder struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: email
func (_m *UserFinder) FindByEmail(email string) (security.User, error) {
	ret := _m.Called(email)

	var r0 security.User
	if rf, ok := ret.Get(0).(func(string) security.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(security.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByFilter provides a mock function with given fields: filter
func (_m *UserFinder) FindByFilter(filter commons.Map) (security.User, error) {
	ret := _m.Called(filter)

	var r0 security.User
	if rf, ok := ret.Get(0).(func(commons.Map) security.User); ok {
		r0 = rf(filter)
	} else {
		r0 = ret.Get(0).(security.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(commons.Map) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *UserFinder) FindByID(id uuid.UUID) (security.User, error) {
	ret := _m.Called(id)

	var r0 security.User
	if rf, ok := ret.Get(0).(func(uuid.UUID) security.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(security.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}