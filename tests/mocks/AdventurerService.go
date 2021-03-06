// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	adventurers "tasquest.com/server/application/gamification/adventurers"

	uuid "github.com/google/uuid"
)

// AdventurerService is an autogenerated mock type for the AdventurerService type
type AdventurerService struct {
	mock.Mock
}

// CreateAdventurer provides a mock function with given fields: command
func (_m *AdventurerService) CreateAdventurer(command adventurers.CreateAdventurer) (adventurers.Adventurer, error) {
	ret := _m.Called(command)

	var r0 adventurers.Adventurer
	if rf, ok := ret.Get(0).(func(adventurers.CreateAdventurer) adventurers.Adventurer); ok {
		r0 = rf(command)
	} else {
		r0 = ret.Get(0).(adventurers.Adventurer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(adventurers.CreateAdventurer) error); ok {
		r1 = rf(command)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAdventurer provides a mock function with given fields: adventurerID, command
func (_m *AdventurerService) UpdateAdventurer(adventurerID uuid.UUID, command adventurers.UpdateAdventurer) (adventurers.Adventurer, error) {
	ret := _m.Called(adventurerID, command)

	var r0 adventurers.Adventurer
	if rf, ok := ret.Get(0).(func(uuid.UUID, adventurers.UpdateAdventurer) adventurers.Adventurer); ok {
		r0 = rf(adventurerID, command)
	} else {
		r0 = ret.Get(0).(adventurers.Adventurer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, adventurers.UpdateAdventurer) error); ok {
		r1 = rf(adventurerID, command)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExperience provides a mock function with given fields: command
func (_m *AdventurerService) UpdateExperience(command adventurers.UpdateExperience) (adventurers.Adventurer, error) {
	ret := _m.Called(command)

	var r0 adventurers.Adventurer
	if rf, ok := ret.Get(0).(func(adventurers.UpdateExperience) adventurers.Adventurer); ok {
		r0 = rf(command)
	} else {
		r0 = ret.Get(0).(adventurers.Adventurer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(adventurers.UpdateExperience) error); ok {
		r1 = rf(command)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
