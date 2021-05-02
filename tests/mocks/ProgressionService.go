// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	leveling "tasquest.com/server/application/gamification/leveling"
)

// ProgressionService is an autogenerated mock type for the ProgressionService type
type ProgressionService struct {
	mock.Mock
}

// AwardExperience provides a mock function with given fields: command
func (_m *ProgressionService) AwardExperience(command leveling.AwardExperience) error {
	ret := _m.Called(command)

	var r0 error
	if rf, ok := ret.Get(0).(func(leveling.AwardExperience) error); ok {
		r0 = rf(command)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateLevel provides a mock function with given fields: command
func (_m *ProgressionService) CreateLevel(command leveling.CreateLevel) (leveling.ExpLevel, error) {
	ret := _m.Called(command)

	var r0 leveling.ExpLevel
	if rf, ok := ret.Get(0).(func(leveling.CreateLevel) leveling.ExpLevel); ok {
		r0 = rf(command)
	} else {
		r0 = ret.Get(0).(leveling.ExpLevel)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(leveling.CreateLevel) error); ok {
		r1 = rf(command)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLevel provides a mock function with given fields: command
func (_m *ProgressionService) DeleteLevel(command leveling.DeleteLevel) (leveling.ExpLevel, error) {
	ret := _m.Called(command)

	var r0 leveling.ExpLevel
	if rf, ok := ret.Get(0).(func(leveling.DeleteLevel) leveling.ExpLevel); ok {
		r0 = rf(command)
	} else {
		r0 = ret.Get(0).(leveling.ExpLevel)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(leveling.DeleteLevel) error); ok {
		r1 = rf(command)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}