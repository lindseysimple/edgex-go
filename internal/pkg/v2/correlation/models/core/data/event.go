//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package data

import "github.com/edgexfoundry/edgex-go/internal/pkg/v2/models/coredata"

type Event struct {
	Bytes         []byte // This will NOT be marshaled via the JSON below. It is only populated and read as an instance member.
	CorrelationId string
	Checksum      string
	data.Event
}
