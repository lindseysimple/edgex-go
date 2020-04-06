package base

import (
	"encoding/json"
	"net/http"
)

// Response defines the Response Content for response DTOs.  This object and its properties correspond to the
// BaseResponse object in the APIv2 specification.
type Response struct {
	CorrelationID  string       `json:"correlationId"`
	RequestID      string       `json:"requestId"`
	Message        interface{}  `json:"message,omitempty"`
	StatusCode     int          `json:"statusCode"`
}

// NewResponse is a factory function that returns a Response struct.
func NewResponse(correlationID string, requestID string, message interface{}, statusCode int) *Response {
	if message != nil {
		if _, ok := message.(string); !ok {
			marshal := func(r interface{}) string {
				b, e := json.Marshal(r)
				if e != nil {
					return e.Error()
				}
				return string(b)
			}

			// if we were passed a non-nil reference to a structure, stringify it.
			message = marshal(message)
		}
	}
	return &Response{
		CorrelationID: correlationID,
		RequestID:     requestID,
		Message:       message,
		StatusCode:    statusCode,
	}
}

// NewResponseForSuccess is a factory function that returns a Response struct.
func NewResponseForSuccess(correlationID string, requestID string) *Response {
	return NewResponse(correlationID, requestID, nil, http.StatusMultiStatus)
}
