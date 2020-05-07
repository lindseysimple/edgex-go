//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import "github.com/edgexfoundry/edgex-go/internal/pkg/config"

// Database interface provides an abstraction for obtaining the database configuration information.
type Database interface {
	// GetDatabaseInfo returns a database information map.
	GetDatabaseInfo() config.DatabaseInfo
}
