// Code generated by mockery v2.49.1. DO NOT EDIT.

package mocks

import (
	errors "github.com/edgexfoundry/go-mod-core-contracts/v4/errors"

	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddEvent provides a mock function with given fields: e
func (_m *DBClient) AddEvent(e models.Event) (models.Event, errors.EdgeX) {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for AddEvent")
	}

	var r0 models.Event
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(models.Event) (models.Event, errors.EdgeX)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(models.Event) models.Event); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	if rf, ok := ret.Get(1).(func(models.Event) errors.EdgeX); ok {
		r1 = rf(e)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllEvents provides a mock function with given fields: offset, limit
func (_m *DBClient) AllEvents(offset int, limit int) ([]models.Event, errors.EdgeX) {
	ret := _m.Called(offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for AllEvents")
	}

	var r0 []models.Event
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int) ([]models.Event, errors.EdgeX)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []models.Event); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) errors.EdgeX); ok {
		r1 = rf(offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllReadings provides a mock function with given fields: offset, limit
func (_m *DBClient) AllReadings(offset int, limit int) ([]models.Reading, errors.EdgeX) {
	ret := _m.Called(offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for AllReadings")
	}

	var r0 []models.Reading
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int) ([]models.Reading, errors.EdgeX)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []models.Reading); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) errors.EdgeX); ok {
		r1 = rf(offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}

// DeleteEventById provides a mock function with given fields: id
func (_m *DBClient) DeleteEventById(id string) errors.EdgeX {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEventById")
	}

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteEventsByAge provides a mock function with given fields: age
func (_m *DBClient) DeleteEventsByAge(age int64) errors.EdgeX {
	ret := _m.Called(age)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEventsByAge")
	}

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int64) errors.EdgeX); ok {
		r0 = rf(age)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteEventsByAgeAndDeviceNameAndSourceName provides a mock function with given fields: age, deviceName, sourceName
func (_m *DBClient) DeleteEventsByAgeAndDeviceNameAndSourceName(age int64, deviceName string, sourceName string) errors.EdgeX {
	ret := _m.Called(age, deviceName, sourceName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEventsByAgeAndDeviceNameAndSourceName")
	}

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int64, string, string) errors.EdgeX); ok {
		r0 = rf(age, deviceName, sourceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteEventsByDeviceName provides a mock function with given fields: deviceName
func (_m *DBClient) DeleteEventsByDeviceName(deviceName string) errors.EdgeX {
	ret := _m.Called(deviceName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEventsByDeviceName")
	}

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(deviceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteEventsByDeviceNameAndSourceName provides a mock function with given fields: deviceName, sourceName
func (_m *DBClient) DeleteEventsByDeviceNameAndSourceName(deviceName string, sourceName string) errors.EdgeX {
	ret := _m.Called(deviceName, sourceName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEventsByDeviceNameAndSourceName")
	}

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, string) errors.EdgeX); ok {
		r0 = rf(deviceName, sourceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// EventById provides a mock function with given fields: id
func (_m *DBClient) EventById(id string) (models.Event, errors.EdgeX) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for EventById")
	}

	var r0 models.Event
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (models.Event, errors.EdgeX)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Event); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// EventCountByDeviceName provides a mock function with given fields: deviceName
func (_m *DBClient) EventCountByDeviceName(deviceName string) (uint32, errors.EdgeX) {
	ret := _m.Called(deviceName)

	if len(ret) == 0 {
		panic("no return value specified for EventCountByDeviceName")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(deviceName)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(deviceName)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(deviceName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// EventCountByDeviceNameAndSourceNameAndLimit provides a mock function with given fields: deviceName, sourceName, limit
func (_m *DBClient) EventCountByDeviceNameAndSourceNameAndLimit(deviceName string, sourceName string, limit int) (uint32, errors.EdgeX) {
	ret := _m.Called(deviceName, sourceName, limit)

	if len(ret) == 0 {
		panic("no return value specified for EventCountByDeviceNameAndSourceNameAndLimit")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, string, int) (uint32, errors.EdgeX)); ok {
		return rf(deviceName, sourceName, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int) uint32); ok {
		r0 = rf(deviceName, sourceName, limit)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string, string, int) errors.EdgeX); ok {
		r1 = rf(deviceName, sourceName, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// EventCountByTimeRange provides a mock function with given fields: start, end
func (_m *DBClient) EventCountByTimeRange(start int64, end int64) (uint32, errors.EdgeX) {
	ret := _m.Called(start, end)

	if len(ret) == 0 {
		panic("no return value specified for EventCountByTimeRange")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int64, int64) (uint32, errors.EdgeX)); ok {
		return rf(start, end)
	}
	if rf, ok := ret.Get(0).(func(int64, int64) uint32); ok {
		r0 = rf(start, end)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(int64, int64) errors.EdgeX); ok {
		r1 = rf(start, end)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// EventTotalCount provides a mock function with given fields:
func (_m *DBClient) EventTotalCount() (uint32, errors.EdgeX) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EventTotalCount")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func() (uint32, errors.EdgeX)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func() errors.EdgeX); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// EventsByDeviceName provides a mock function with given fields: offset, limit, name
func (_m *DBClient) EventsByDeviceName(offset int, limit int, name string) ([]models.Event, errors.EdgeX) {
	ret := _m.Called(offset, limit, name)

	if len(ret) == 0 {
		panic("no return value specified for EventsByDeviceName")
	}

	var r0 []models.Event
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Event, errors.EdgeX)); ok {
		return rf(offset, limit, name)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Event); ok {
		r0 = rf(offset, limit, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// EventsByTimeRange provides a mock function with given fields: start, end, offset, limit
func (_m *DBClient) EventsByTimeRange(start int64, end int64, offset int, limit int) ([]models.Event, errors.EdgeX) {
	ret := _m.Called(start, end, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for EventsByTimeRange")
	}

	var r0 []models.Event
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int64, int64, int, int) ([]models.Event, errors.EdgeX)); ok {
		return rf(start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, int, int) []models.Event); ok {
		r0 = rf(start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, int64, int, int) errors.EdgeX); ok {
		r1 = rf(start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// LatestEventByDeviceNameAndSourceNameAndAgeAndOffset provides a mock function with given fields: deviceName, sourceName, age, offset
func (_m *DBClient) LatestEventByDeviceNameAndSourceNameAndAgeAndOffset(deviceName string, sourceName string, age int64, offset uint32) (models.Event, errors.EdgeX) {
	ret := _m.Called(deviceName, sourceName, age, offset)

	if len(ret) == 0 {
		panic("no return value specified for LatestEventByDeviceNameAndSourceNameAndAgeAndOffset")
	}

	var r0 models.Event
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, string, int64, uint32) (models.Event, errors.EdgeX)); ok {
		return rf(deviceName, sourceName, age, offset)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64, uint32) models.Event); ok {
		r0 = rf(deviceName, sourceName, age, offset)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	if rf, ok := ret.Get(1).(func(string, string, int64, uint32) errors.EdgeX); ok {
		r1 = rf(deviceName, sourceName, age, offset)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// LatestEventByDeviceNameAndSourceNameAndOffset provides a mock function with given fields: deviceName, sourceName, offset
func (_m *DBClient) LatestEventByDeviceNameAndSourceNameAndOffset(deviceName string, sourceName string, offset uint32) (models.Event, errors.EdgeX) {
	ret := _m.Called(deviceName, sourceName, offset)

	if len(ret) == 0 {
		panic("no return value specified for LatestEventByDeviceNameAndSourceNameAndOffset")
	}

	var r0 models.Event
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, string, uint32) (models.Event, errors.EdgeX)); ok {
		return rf(deviceName, sourceName, offset)
	}
	if rf, ok := ret.Get(0).(func(string, string, uint32) models.Event); ok {
		r0 = rf(deviceName, sourceName, offset)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	if rf, ok := ret.Get(1).(func(string, string, uint32) errors.EdgeX); ok {
		r1 = rf(deviceName, sourceName, offset)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// LatestReadingByOffset provides a mock function with given fields: offset
func (_m *DBClient) LatestReadingByOffset(offset uint32) (models.Reading, errors.EdgeX) {
	ret := _m.Called(offset)

	if len(ret) == 0 {
		panic("no return value specified for LatestReadingByOffset")
	}

	var r0 models.Reading
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(uint32) (models.Reading, errors.EdgeX)); ok {
		return rf(offset)
	}
	if rf, ok := ret.Get(0).(func(uint32) models.Reading); ok {
		r0 = rf(offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.Reading)
		}
	}

	if rf, ok := ret.Get(1).(func(uint32) errors.EdgeX); ok {
		r1 = rf(offset)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingCountByDeviceName provides a mock function with given fields: deviceName
func (_m *DBClient) ReadingCountByDeviceName(deviceName string) (uint32, errors.EdgeX) {
	ret := _m.Called(deviceName)

	if len(ret) == 0 {
		panic("no return value specified for ReadingCountByDeviceName")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(deviceName)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(deviceName)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(deviceName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingCountByDeviceNameAndResourceName provides a mock function with given fields: deviceName, resourceName
func (_m *DBClient) ReadingCountByDeviceNameAndResourceName(deviceName string, resourceName string) (uint32, errors.EdgeX) {
	ret := _m.Called(deviceName, resourceName)

	if len(ret) == 0 {
		panic("no return value specified for ReadingCountByDeviceNameAndResourceName")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, string) (uint32, errors.EdgeX)); ok {
		return rf(deviceName, resourceName)
	}
	if rf, ok := ret.Get(0).(func(string, string) uint32); ok {
		r0 = rf(deviceName, resourceName)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string, string) errors.EdgeX); ok {
		r1 = rf(deviceName, resourceName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingCountByDeviceNameAndResourceNameAndTimeRange provides a mock function with given fields: deviceName, resourceName, start, end
func (_m *DBClient) ReadingCountByDeviceNameAndResourceNameAndTimeRange(deviceName string, resourceName string, start int64, end int64) (uint32, errors.EdgeX) {
	ret := _m.Called(deviceName, resourceName, start, end)

	if len(ret) == 0 {
		panic("no return value specified for ReadingCountByDeviceNameAndResourceNameAndTimeRange")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, string, int64, int64) (uint32, errors.EdgeX)); ok {
		return rf(deviceName, resourceName, start, end)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64, int64) uint32); ok {
		r0 = rf(deviceName, resourceName, start, end)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string, string, int64, int64) errors.EdgeX); ok {
		r1 = rf(deviceName, resourceName, start, end)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingCountByDeviceNameAndResourceNamesAndTimeRange provides a mock function with given fields: deviceName, resourceName, start, end
func (_m *DBClient) ReadingCountByDeviceNameAndResourceNamesAndTimeRange(deviceName string, resourceName []string, start int64, end int64) (uint32, errors.EdgeX) {
	ret := _m.Called(deviceName, resourceName, start, end)

	if len(ret) == 0 {
		panic("no return value specified for ReadingCountByDeviceNameAndResourceNamesAndTimeRange")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, []string, int64, int64) (uint32, errors.EdgeX)); ok {
		return rf(deviceName, resourceName, start, end)
	}
	if rf, ok := ret.Get(0).(func(string, []string, int64, int64) uint32); ok {
		r0 = rf(deviceName, resourceName, start, end)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string, []string, int64, int64) errors.EdgeX); ok {
		r1 = rf(deviceName, resourceName, start, end)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingCountByDeviceNameAndTimeRange provides a mock function with given fields: deviceName, start, end
func (_m *DBClient) ReadingCountByDeviceNameAndTimeRange(deviceName string, start int64, end int64) (uint32, errors.EdgeX) {
	ret := _m.Called(deviceName, start, end)

	if len(ret) == 0 {
		panic("no return value specified for ReadingCountByDeviceNameAndTimeRange")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, int64, int64) (uint32, errors.EdgeX)); ok {
		return rf(deviceName, start, end)
	}
	if rf, ok := ret.Get(0).(func(string, int64, int64) uint32); ok {
		r0 = rf(deviceName, start, end)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string, int64, int64) errors.EdgeX); ok {
		r1 = rf(deviceName, start, end)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingCountByResourceName provides a mock function with given fields: resourceName
func (_m *DBClient) ReadingCountByResourceName(resourceName string) (uint32, errors.EdgeX) {
	ret := _m.Called(resourceName)

	if len(ret) == 0 {
		panic("no return value specified for ReadingCountByResourceName")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(resourceName)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(resourceName)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(resourceName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingCountByResourceNameAndTimeRange provides a mock function with given fields: resourceName, start, end
func (_m *DBClient) ReadingCountByResourceNameAndTimeRange(resourceName string, start int64, end int64) (uint32, errors.EdgeX) {
	ret := _m.Called(resourceName, start, end)

	if len(ret) == 0 {
		panic("no return value specified for ReadingCountByResourceNameAndTimeRange")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, int64, int64) (uint32, errors.EdgeX)); ok {
		return rf(resourceName, start, end)
	}
	if rf, ok := ret.Get(0).(func(string, int64, int64) uint32); ok {
		r0 = rf(resourceName, start, end)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string, int64, int64) errors.EdgeX); ok {
		r1 = rf(resourceName, start, end)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingCountByTimeRange provides a mock function with given fields: start, end
func (_m *DBClient) ReadingCountByTimeRange(start int64, end int64) (uint32, errors.EdgeX) {
	ret := _m.Called(start, end)

	if len(ret) == 0 {
		panic("no return value specified for ReadingCountByTimeRange")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int64, int64) (uint32, errors.EdgeX)); ok {
		return rf(start, end)
	}
	if rf, ok := ret.Get(0).(func(int64, int64) uint32); ok {
		r0 = rf(start, end)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(int64, int64) errors.EdgeX); ok {
		r1 = rf(start, end)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingTotalCount provides a mock function with given fields:
func (_m *DBClient) ReadingTotalCount() (uint32, errors.EdgeX) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReadingTotalCount")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func() (uint32, errors.EdgeX)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func() errors.EdgeX); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByDeviceName provides a mock function with given fields: offset, limit, name
func (_m *DBClient) ReadingsByDeviceName(offset int, limit int, name string) ([]models.Reading, errors.EdgeX) {
	ret := _m.Called(offset, limit, name)

	if len(ret) == 0 {
		panic("no return value specified for ReadingsByDeviceName")
	}

	var r0 []models.Reading
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Reading, errors.EdgeX)); ok {
		return rf(offset, limit, name)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Reading); ok {
		r0 = rf(offset, limit, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByDeviceNameAndResourceName provides a mock function with given fields: deviceName, resourceName, offset, limit
func (_m *DBClient) ReadingsByDeviceNameAndResourceName(deviceName string, resourceName string, offset int, limit int) ([]models.Reading, errors.EdgeX) {
	ret := _m.Called(deviceName, resourceName, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for ReadingsByDeviceNameAndResourceName")
	}

	var r0 []models.Reading
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, string, int, int) ([]models.Reading, errors.EdgeX)); ok {
		return rf(deviceName, resourceName, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) []models.Reading); ok {
		r0 = rf(deviceName, resourceName, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) errors.EdgeX); ok {
		r1 = rf(deviceName, resourceName, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByDeviceNameAndResourceNameAndTimeRange provides a mock function with given fields: deviceName, resourceName, start, end, offset, limit
func (_m *DBClient) ReadingsByDeviceNameAndResourceNameAndTimeRange(deviceName string, resourceName string, start int64, end int64, offset int, limit int) ([]models.Reading, errors.EdgeX) {
	ret := _m.Called(deviceName, resourceName, start, end, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for ReadingsByDeviceNameAndResourceNameAndTimeRange")
	}

	var r0 []models.Reading
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, string, int64, int64, int, int) ([]models.Reading, errors.EdgeX)); ok {
		return rf(deviceName, resourceName, start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64, int64, int, int) []models.Reading); ok {
		r0 = rf(deviceName, resourceName, start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int64, int64, int, int) errors.EdgeX); ok {
		r1 = rf(deviceName, resourceName, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByDeviceNameAndResourceNamesAndTimeRange provides a mock function with given fields: deviceName, resourceNames, start, end, offset, limit
func (_m *DBClient) ReadingsByDeviceNameAndResourceNamesAndTimeRange(deviceName string, resourceNames []string, start int64, end int64, offset int, limit int) ([]models.Reading, errors.EdgeX) {
	ret := _m.Called(deviceName, resourceNames, start, end, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for ReadingsByDeviceNameAndResourceNamesAndTimeRange")
	}

	var r0 []models.Reading
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, []string, int64, int64, int, int) ([]models.Reading, errors.EdgeX)); ok {
		return rf(deviceName, resourceNames, start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, []string, int64, int64, int, int) []models.Reading); ok {
		r0 = rf(deviceName, resourceNames, start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string, int64, int64, int, int) errors.EdgeX); ok {
		r1 = rf(deviceName, resourceNames, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByDeviceNameAndTimeRange provides a mock function with given fields: deviceName, start, end, offset, limit
func (_m *DBClient) ReadingsByDeviceNameAndTimeRange(deviceName string, start int64, end int64, offset int, limit int) ([]models.Reading, errors.EdgeX) {
	ret := _m.Called(deviceName, start, end, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for ReadingsByDeviceNameAndTimeRange")
	}

	var r0 []models.Reading
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, int64, int64, int, int) ([]models.Reading, errors.EdgeX)); ok {
		return rf(deviceName, start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int64, int64, int, int) []models.Reading); ok {
		r0 = rf(deviceName, start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int64, int64, int, int) errors.EdgeX); ok {
		r1 = rf(deviceName, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByResourceName provides a mock function with given fields: offset, limit, resourceName
func (_m *DBClient) ReadingsByResourceName(offset int, limit int, resourceName string) ([]models.Reading, errors.EdgeX) {
	ret := _m.Called(offset, limit, resourceName)

	if len(ret) == 0 {
		panic("no return value specified for ReadingsByResourceName")
	}

	var r0 []models.Reading
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Reading, errors.EdgeX)); ok {
		return rf(offset, limit, resourceName)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Reading); ok {
		r0 = rf(offset, limit, resourceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, resourceName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByResourceNameAndTimeRange provides a mock function with given fields: resourceName, start, end, offset, limit
func (_m *DBClient) ReadingsByResourceNameAndTimeRange(resourceName string, start int64, end int64, offset int, limit int) ([]models.Reading, errors.EdgeX) {
	ret := _m.Called(resourceName, start, end, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for ReadingsByResourceNameAndTimeRange")
	}

	var r0 []models.Reading
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string, int64, int64, int, int) ([]models.Reading, errors.EdgeX)); ok {
		return rf(resourceName, start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int64, int64, int, int) []models.Reading); ok {
		r0 = rf(resourceName, start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int64, int64, int, int) errors.EdgeX); ok {
		r1 = rf(resourceName, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByTimeRange provides a mock function with given fields: start, end, offset, limit
func (_m *DBClient) ReadingsByTimeRange(start int64, end int64, offset int, limit int) ([]models.Reading, errors.EdgeX) {
	ret := _m.Called(start, end, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for ReadingsByTimeRange")
	}

	var r0 []models.Reading
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int64, int64, int, int) ([]models.Reading, errors.EdgeX)); ok {
		return rf(start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, int, int) []models.Reading); ok {
		r0 = rf(start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, int64, int, int) errors.EdgeX); ok {
		r1 = rf(start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NewDBClient creates a new instance of DBClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDBClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *DBClient {
	mock := &DBClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
