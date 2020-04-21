//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"context"
	"fmt"
	"github.com/edgexfoundry/edgex-go/internal/core/data/config"
	"github.com/edgexfoundry/edgex-go/internal/core/data/errors"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/errorconcept"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/correlation/models/core/data"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/db/interfaces"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/http/common"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/http/mapper"
	v2model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/models/coredata"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/metadata"
	"github.com/edgexfoundry/go-mod-messaging/messaging"
	"github.com/gorilla/mux"
	"net/http"
)

func V2EventHandler(
	w http.ResponseWriter,
	r *http.Request,
	lc logger.LoggingClient,
	dbClient interfaces.DBClient,
	chEvents chan<- interface{},
	msgClient messaging.MessageClient,
	mdc metadata.DeviceClient,
	httpErrorHandler errorconcept.ErrorHandler,
	configuration *config.ConfigurationStruct) {

	if r.Body != nil {
		defer func() { _ = r.Body.Close() }()
	}

	ctx := r.Context()

	switch r.Method {
	case http.MethodPost:
		reader := common.NewRequestReader(r, configuration)

		//evt := data.Event{}
		evt, err := reader.Read(r.Body, &ctx)

		fmt.Printf("post event %+v\n", evt)
		//err := json.NewDecoder(reader).Decode(&evt)

		if err != nil {
			httpErrorHandler.Handle(w, err, errorconcept.Default.InternalServerError)
			return
		}
		newId, err := addNewEvent(evt, ctx, lc, dbClient, chEvents, msgClient, mdc, configuration)

		if err != nil {
			httpErrorHandler.HandleManyVariants(
				w,
				err,
				[]errorconcept.ErrorConceptType{
					errorconcept.ValueDescriptors.NotFound,
					errorconcept.ValueDescriptors.Invalid,
					errorconcept.NewServiceClientHttpError(err),
				},
				errorconcept.Default.InternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(newId))
		break
	}
}

// GET
// Return the event specified by the event ID
// /api/v1/event/{id}
// id - ID of the event to return
func V2GetEventByIdHandler(
	w http.ResponseWriter,
	r *http.Request,
	lc logger.LoggingClient,
	dbClient interfaces.DBClient,
	httpErrorHandler errorconcept.ErrorHandler) {

	if r.Body != nil {
		defer func() { _ = r.Body.Close() }()
	}

	// URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the event
	event, err := getEventById(id, dbClient)
	if err != nil {
		httpErrorHandler.HandleOneVariant(
			w,
			err,
			errorconcept.Events.NotFound,
			errorconcept.Default.InternalServerError)
		return
	}
	eventDTO := mapper.ToEventDTO(event)
	common.Encode(eventDTO, w, lc)
}

func addNewEvent(
	e data.Event, ctx context.Context,
	lc logger.LoggingClient,
	dbClient interfaces.DBClient,
	chEvents chan<- interface{},
	msgClient messaging.MessageClient,
	mdc metadata.DeviceClient,
	configuration *config.ConfigurationStruct) (string, error) {

	//err := checkDevice(e.Device, ctx, mdc, configuration)
	//if err != nil {
	//	return "", err
	//}

	//if configuration.Writable.ValidateCheck {
	//	lc.Debug("Validation enabled, parsing events")
	//	for reading := range e.Readings {
	//		// Check value descriptor
	//		name := e.Readings[reading].Name
	//		vd, err := dbClient.ValueDescriptorByName(name)
	//		if err != nil {
	//			if err == db.ErrNotFound {
	//				return "", errors.NewErrValueDescriptorNotFound(name)
	//			} else {
	//				return "", err
	//			}
	//		}
	//		err = isValidValueDescriptor(vd, e.Readings[reading])
	//		if err != nil {
	//			return "", err
	//		}
	//	}
	//}

	// Add the event and readings to the database
	if configuration.Writable.PersistData {
		id, err := dbClient.AddEvent(e)
		if err != nil {
			return "", err
		}
		e.ID = id
	}

	//putEventOnQueue(e, ctx, lc, msgClient, configuration) // Push event to message bus for App Services to consume
	//chEvents <- DeviceLastReported{e.Device}              // update last reported connected (device)
	//chEvents <- DeviceServiceLastReported{e.Device}       // update last reported connected (device service)

	return e.ID, nil
}

func getEventById(id string, dbClient interfaces.DBClient) (event v2model.Event, err error) {
	event, err = dbClient.GetEventById(id)
	if err != nil {
		if err == db.ErrNotFound {
			err = errors.NewErrEventNotFound(id)
		}
		return event, err
	}

	return event, nil
}