//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package data

import (
	dtoBase "github.com/edgexfoundry/edgex-go/internal/pkg/v2/dto/common/base"
	model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/model/core/data"
)

// AddEventRequest defines the Request Content for POST event DTO. This object and its properties correspond to the
// AddEventRequest object in the APIv2 specification.
type AddEventRequest struct {
	dtoBase.Request `json:",inline"`
	Device          string          `json:"device"`
	Origin          string          `json:"origin,omitempty" codec:"origin,omitempty"`
	Readings        []model.Reading `json:"readings,omitempty" codec:"readings,omitempty"`
}

// AddEventResponse defines the Response Content for POST event DTOs.  This object and its properties correspond to the
// AddEventResponse object in the APIv2 specification.
type AddEventResponse struct {
	dtoBase.Response `json:",inline"`
	ID               string `json:"id"` // ID uniquely identifies an event, for example a UUID
}

// EventCountResponse defines the Response Content for GET event count DTO.  This object and its properties correspond to the
// EventCountResponse object in the APIv2 specification.
type EventCountResponse struct {
	dtoBase.Response `json:",inline"`
	Count            int
	DeviceID         string `json:"deviceId"` // ID uniquely identifies a device
}

// EventResponse defines the Response Content for GET event DTOs.  This object and its properties correspond to the
// EventResponse object in the APIv2 specification.
type EventResponse struct {
	dtoBase.Response `json:",inline"`
	Event            model.Event
}

// UpdateEventPushedRequest defines the Request Content for PUT event as pushed DTO. This object and its properties correspond to the
// UpdateEventPushedByChecksumRequest object in the APIv2 specification.
type UpdateEventPushedRequest struct {
	dtoBase.Request `json:",inline"`
	Checksum        string `json:"checksum"`
}
