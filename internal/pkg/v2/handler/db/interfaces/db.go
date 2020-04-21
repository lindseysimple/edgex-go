//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/correlation/models/core/data"
	v2model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/models/coredata"
)

type DBClient interface {
	CloseSession()

	AddEvent(e data.Event) (string, error)
	GetEventById(id string) (v2model.Event, error)
}
