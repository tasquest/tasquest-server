// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	adventurers "tasquest.com/server/application/gamification/adventurers"
)

// AdventurerPersistence is an autogenerated mock type for the AdventurerPersistence type
type AdventurerPersistence struct {
	mock.Mock
}

// Save provides a mock function with given fields: adventurer
func (_m *AdventurerPersistence) Save(adventurer adventurers.Adventurer) (adventurers.Adventurer, error) {
	ret := _m.Called(adventurer)

	var r0 adventurers.Adventurer
	if rf, ok := ret.Get(0).(func(adventurers.Adventurer) adventurers.Adventurer); ok {
		r0 = rf(adventurer)
	} else {
		r0 = ret.Get(0).(adventurers.Adventurer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(adventurers.Adventurer) error); ok {
		r1 = rf(adventurer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: adventurer
func (_m *AdventurerPersistence) Update(adventurer adventurers.Adventurer) (adventurers.Adventurer, error) {
	ret := _m.Called(adventurer)

	var r0 adventurers.Adventurer
	if rf, ok := ret.Get(0).(func(adventurers.Adventurer) adventurers.Adventurer); ok {
		r0 = rf(adventurer)
	} else {
		r0 = ret.Get(0).(adventurers.Adventurer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(adventurers.Adventurer) error); ok {
		r1 = rf(adventurer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
