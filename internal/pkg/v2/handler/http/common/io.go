//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"context"
	"encoding/json"
	"github.com/edgexfoundry/edgex-go/internal/core/data/config"
	dto "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/dtos/coredata"
	model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/models/coredata"
	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"io"
	"net/http"
	"strings"
)

// jsonReader handles unmarshaling of a JSON request body payload
type jsonReader struct{}

// EventReader unmarshals a request body into an Event type
type EventReader interface {
	Read(reader io.Reader, ctx *context.Context) ([]model.Event, error)
}

// Read reads and converts the request's JSON event data into an Event struct
func (jsonReader) Read(reader io.Reader, ctx *context.Context) (events []model.Event, err error) {
	c := context.WithValue(*ctx, clients.ContentType, clients.ContentTypeJSON)
	*ctx = c
	var addEvents []dto.AddEventRequest
	err = json.NewDecoder(reader).Decode(&addEvents)
	if err != nil {
		return events, err
	}
	events = dto.ToEventModels(addEvents)
	return events, nil
}

// NewJsonReader creates a new instance of cborReader.
func NewJsonReader() jsonReader {
	return jsonReader{}
}

// NewRequestReader returns a BodyReader capable of processing the request body
func NewRequestReader(request *http.Request, configuration *config.ConfigurationStruct) EventReader {
	contentType := request.Header.Get(clients.ContentType)
	switch strings.ToLower(contentType) {
	case clients.ContentTypeCBOR:
		return NewJsonReader()
	default:
		return NewJsonReader()
	}
}
