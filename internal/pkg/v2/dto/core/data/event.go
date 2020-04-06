package data

import (
	dtoBase  "github.com/edgexfoundry/edgex-go/internal/pkg/v2/dto/common/base"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
)

// AddEventRequest defines the Request Content for POST event DTO. This object and its properties correspond to the
// AddEventRequest object in the APIv2 specification.
type AddEventRequest struct {
	dtoBase.Request `json:",inline"`
	Event           contract.Event
}

// AddEventResponse defines the Response Content for POST event DTOs.  This object and its properties correspond to the
// AddEventResponse object in the APIv2 specification.
type AddEventResponse struct {
	dtoBase.Response `json:",inline"`
	ID string        `json:"id"`       // ID uniquely identifies an event, for example a UUID
}

// EventCountResponse defines the Response Content for GET event count DTO.  This object and its properties correspond to the
// EventCountResponse object in the APIv2 specification.
type EventCountResponse struct {
	dtoBase.Response `json:",inline"`
	Count            int
}

// EventResponse defines the Response Content for GET event DTOs.  This object and its properties correspond to the
// EventResponse object in the APIv2 specification.
type EventResponse struct {
	dtoBase.Response `json:",inline"`
	Event            contract.Event
}

// UpdateEventPushedRequest defines the Request Content for PUT event as pushed DTO. This object and its properties correspond to the
// UpdateEventPushedByChecksumRequest object in the APIv2 specification.
type UpdateEventPushedRequest struct {
	dtoBase.Request `json:",inline"`
	Checksum string `json:"checksum"`
}