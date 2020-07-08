//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package infrastructure

import (
	"errors"
	"time"
)

const (
	// Data
	EventsCollection   = "v2:event"
	ReadingsCollection = "v2:reading"
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
