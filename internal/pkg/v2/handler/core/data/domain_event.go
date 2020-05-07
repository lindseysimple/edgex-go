//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package data

// An event indicating that a given device has just reported some data
type DeviceLastReported struct {
	DeviceName string
}

// An event indicating that the service associated with the device that just reported data is alive.
type DeviceServiceLastReported struct {
	DeviceName string
}
