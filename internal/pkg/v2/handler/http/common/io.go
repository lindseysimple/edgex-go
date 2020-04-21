//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/edgex-go/internal/core/data/config"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/correlation/models/core/data"
	dtos "github.com/edgexfoundry/edgex-go/internal/pkg/v2/dtos/coredata"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/http/mapper"
	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"gopkg.in/dealancer/validate.v2"
	"io"
	"net/http"
	"strings"
)

// jsonReader handles unmarshaling of a JSON request body payload
type jsonReader struct{}

// EventReader unmarshals a request body into an Event type
type EventReader interface {
	//Read(reader io.Reader, ctx *context.Context) (data.Event, error)
	Read(reader io.Reader, ctx *context.Context) (data.Event, error)
}

// Read reads and converts the request's JSON event data into an Event struct
//func (jsonReader) Read(reader io.Reader, ctx *context.Context) (data.Event, error) {
func (jsonReader) Read(reader io.Reader, ctx *context.Context) (event data.Event, err error) {
	c := context.WithValue(*ctx, clients.ContentType, clients.ContentTypeJSON)
	*ctx = c
	addEventDTO := dtos.AddEventRequest{}
	err = json.NewDecoder(reader).Decode(&addEventDTO)
	if err != nil {
		return event, err
	}
	if err := validate.Validate(addEventDTO); err != nil {
		fmt.Println(err) // Prints "Field must not be empty"
		return event, err
	}
	event = mapper.ToEventContract(addEventDTO)
	if err != nil {
		return event, err
	}
	return event, nil
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
