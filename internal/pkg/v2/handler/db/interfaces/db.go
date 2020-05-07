//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/models/coredata"
)

type DBClient interface {
	CloseSession()

	AddEvent(e model.Event) (string, error)
	GetEventById(id string) (model.Event, error)
	GetEvents() ([]model.Event, error)
	GetEventCount() (int, error)
	GetEventCountByDeviceId(id string) (int, error)
	GetEventsByDeviceId(id string) ([]model.Event, error)
}
