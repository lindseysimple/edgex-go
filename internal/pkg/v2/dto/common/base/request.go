package base

// Request defines the Request Content for request DTOs. This object and its properties correspond to the BaseRequest
// object in the APIv2 specification.
type Request struct {
	CorrelationID  string `json:"correlationId"`
	RequestID      string `json:"requestId"`
}

// NewRequest is a factory function that returns a Request struct.
func NewRequest(correlationID string, requestID string) *Request {
	return &Request{
		CorrelationID: correlationID,
		RequestID:     requestID,
	}
}