//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package coredata

import (
	"encoding/json"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod"
	dtoBase "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/dtos/common/base"
	model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/models/coredata"
)

// AddEventRequest defines the Request Content for POST event DTO. This object and its properties correspond to the
// AddEventRequest object in the APIv2 specification.
type AddEventRequest struct {
	dtoBase.Request `json:",inline"`
	Device          string        `json:"device" codec:"device,omitempty" validate:"required"`
	Origin          int64         `json:"origin,omitempty" codec:"origin,omitempty" validate:"gte=1"`
	Readings        []BaseReading `json:"readings,omitempty" codec:"readings,omitempty"`
}

// AddEventResponse defines the Response Content for POST event DTOs.  This object and its properties correspond to the
// AddEventResponse object in the APIv2 specification.
type AddEventResponse struct {
	dtoBase.Response `json:",inline"`
	ID               string `json:"id"` // ID uniquely identifies an event, for example a UUID
}

// Event represents a single measurable event read from a device
type Event struct {
	ID       string        `json:"id,omitempty" codec:"id,omitempty"`             // ID uniquely identifies an event, for example a UUID
	Pushed   int64         `json:"pushed,omitempty" codec:"pushed,omitempty"`     // Pushed is a timestamp indicating when the event was exported. If unexported, the value is zero.
	Device   string        `json:"device,omitempty" codec:"device,omitempty"`     // Device identifies the source of the event, can be a device name or id. Usually the device name.
	Created  int64         `json:"created,omitempty" codec:"created,omitempty"`   // Created is a timestamp indicating when the event was created.
	Modified int64         `json:"modified,omitempty" codec:"modified,omitempty"` // Modified is a timestamp indicating when the event was last modified.
	Origin   int64         `json:"origin,omitempty" codec:"origin,omitempty"`     // Origin is a timestamp that can communicate the time of the original reading, prior to event creation
	Readings []BaseReading `json:"readings,omitempty" codec:"readings,omitempty"` // Readings will contain zero to many entries for the associated readings of a given event.
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
	Event            Event
}

// UpdateEventPushedByChecksumRequest defines the Request Content for PUT event as pushed DTO. This object and its properties correspond to the
// UpdateEventPushedByChecksumRequest object in the APIv2 specification.
type UpdateEventPushedByChecksumRequest struct {
	dtoBase.Request `json:",inline"`
	Checksum        string `json:"checksum"`
}

// UpdateEventPushedByChecksumResponse defines the Response Content for PUT event as pushed DTO. This object and its properties correspond to the
// UpdateEventPushedByChecksumResponse object in the APIv2 specification.
type UpdateEventPushedByChecksumResponse struct {
	dtoBase.Response `json:",inline"`
	Checksum         string `json:"checksum"`
}

// Validate satisfies the Validator interface
func (a AddEventRequest) Validate() error {
	v2.NewValidator()
	err := v2.Validate(a)
	return err
}

func (a *AddEventRequest) UnmarshalJSON(b []byte) error {
	var addEvent struct {
		dtoBase.Request
		Device   string
		Origin   int64
		Readings []BaseReading
	}
	if err := json.Unmarshal(b, &addEvent); err != nil {
		return v2.NewErrContractInvalid("Request body type is not a JSON array.")
	}

	a.RequestID = addEvent.RequestID
	a.CorrelationID = addEvent.CorrelationID
	a.Device = addEvent.Device
	a.Origin = addEvent.Origin
	a.Readings = addEvent.Readings

	// validate AddEventRequest DTO
	if err := a.Validate(); err != nil {
		return err
	}
	return nil
}

func ToEventModels(addRequests []AddEventRequest) (events []model.Event) {
	for _, a := range addRequests {
		var e model.Event
		readings := make([]model.Reading, len(a.Readings))
		for i, r := range a.Readings {
			readings[i] = toReadingModel(r, a.Device)
		}
		e.CorrelationId = a.CorrelationID
		e.RequestId = a.RequestID
		e.Device = a.Device
		e.Origin = a.Origin
		e.Readings = readings
		events = append(events, e)
	}
	return events
}
