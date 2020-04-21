//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package db

import (
	"errors"
	"time"
)

const (
	// Databases
	MongoDB = "mongodb"
	RedisDB = "redisdb"

	// Data
	EventsCollection          = "v2:event"
	ReadingsCollection        = "v2:reading"
	ValueDescriptorCollection = "valueDescriptor"

	//Export
	ExportCollection = "exportConfiguration"

	//Logging
	LogsCollection = "logEntry"

	// Metadata
	Device           = "device"
	DeviceProfile    = "deviceProfile"
	DeviceService    = "deviceService"
	Addressable      = "addressable"
	Command          = "command"
	DeviceReport     = "deviceReport"
	ProvisionWatcher = "provisionWatcher"
	Interval         = "interval"
	IntervalAction   = "intervalAction"

	// Notification
	Notification = "notification"
	Subscription = "subscription"
	Transmission = "transmission"
)

var (
	ErrNotFound            = errors.New("Item not found")
	ErrUnsupportedDatabase = errors.New("Unsupported database type")
	ErrInvalidObjectId     = errors.New("Invalid object ID")
	ErrNotUnique           = errors.New("Resource already exists")
	ErrCommandStillInUse   = errors.New("Command is still in use by device profiles")
	ErrSlugEmpty           = errors.New("Slug is nil or empty")
	ErrNameEmpty           = errors.New("Name is required")
)

type Configuration struct {
	DbType       string
	Host         string
	Port         int
	Timeout      int
	DatabaseName string
	Username     string
	Password     string
	BatchSize    int
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
