//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"context"
	"github.com/edgexfoundry/edgex-go/internal/core/data/config"
	"github.com/edgexfoundry/edgex-go/internal/core/data/errors"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/errorconcept"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/dtos/common/base"
	dto "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/dtos/coredata"
	model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/models/coredata"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/db/interfaces"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/http/common"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/http/mapper"
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
		events, err := reader.Read(r.Body, &ctx)
		if err != nil {
			httpErrorHandler.Handle(w, err, errorconcept.Default.InternalServerError)
			return
		}
		var addResponses []dto.AddEventResponse
		for _, e := range events {
			newId, err := addNewEvent(e, ctx, lc, dbClient, chEvents, msgClient, mdc, configuration)
			var addEventResponse dto.AddEventResponse
			if err == nil {
				addEventResponse = dto.AddEventResponse{
					Response: base.Response{
						CorrelationID: e.CorrelationId,
						RequestID: e.RequestId,
						StatusCode: http.StatusAccepted,
					},
					ID:       newId,
				}
			} else {
				addEventResponse = dto.AddEventResponse{
					Response: base.Response{
						CorrelationID: e.CorrelationId,
						RequestID: e.RequestId,
						StatusCode: http.StatusBadRequest,
					},
				}
			}
			addResponses = append(addResponses, addEventResponse)
		}
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

		w.WriteHeader(http.StatusMultiStatus)
		common.Encode(addResponses, w, lc)
		break
	}
}

// GET
// Return all the events
// /api/v2/event/all
func V2GetAllEventHandler(
	w http.ResponseWriter,
	r *http.Request,
	lc logger.LoggingClient,
	dbClient interfaces.DBClient,
	httpErrorHandler errorconcept.ErrorHandler) {

	if r.Body != nil {
		defer func() { _ = r.Body.Close() }()
	}

	// Get the event
	events, err := getAllEvents(dbClient)
	if err != nil {
		httpErrorHandler.HandleOneVariant(
			w,
			err,
			errorconcept.Events.NotFound,
			errorconcept.Default.InternalServerError)
		return
	}
	var eventDTOs []dto.Event
	for _, e := range events {
		eventDTO := mapper.ToEventDTO(e)
		eventDTOs = append(eventDTOs, eventDTO)
	}
	common.Encode(eventDTOs, w, lc)
}

// GET
// Return the event specified by the event ID
// /api/v2/event/id/{id}
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

// GET
// Return the event count
// /api/v2/event/count
func V2GetEventCountHandler(
	w http.ResponseWriter,
	r *http.Request,
	lc logger.LoggingClient,
	dbClient interfaces.DBClient,
	httpErrorHandler errorconcept.ErrorHandler) {

	if r.Body != nil {
		defer func() { _ = r.Body.Close() }()
	}

	// Get the event
	count, err := getEventCount(dbClient)
	if err != nil {
		httpErrorHandler.HandleOneVariant(
			w,
			err,
			errorconcept.Events.NotFound,
			errorconcept.Default.InternalServerError)
		return
	}
	eventCountResp := dto.EventCountResponse{
		Response: base.Response{},
		Count:    count,
		DeviceID: "",
	}
	common.Encode(eventCountResp, w, lc)
}

// GET
// Return the event count specified by the device ID
// /api/v2/event/count/device/{deviceId}
func V2GetEventCountByDeviceIdHandler(
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
	deviceId := vars["deviceId"]
	// Get the event
	count, err := getEventCountByDeviceId(deviceId, dbClient)
	if err != nil {
		httpErrorHandler.HandleOneVariant(
			w,
			err,
			errorconcept.Events.NotFound,
			errorconcept.Default.InternalServerError)
		return
	}
	eventCountResp := dto.EventCountResponse{
		Response: base.Response{},
		Count:    count,
		DeviceID: "",
	}
	common.Encode(eventCountResp, w, lc)
}

// GET
// Return all the events specified by the device ID
// /api/v2/event/device/{deviceId}/all
func V2GetAllEventByDeviceIdHandler(
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
	deviceId := vars["deviceId"]
	// Get the event
	events, err := getAllEventsByDeviceId(deviceId, dbClient)
	if err != nil {
		httpErrorHandler.HandleOneVariant(
			w,
			err,
			errorconcept.Events.NotFound,
			errorconcept.Default.InternalServerError)
		return
	}
	var eventDTOs []dto.Event
	for _, e := range events {
		eventDTO := mapper.ToEventDTO(e)
		eventDTOs = append(eventDTOs, eventDTO)
	}
	common.Encode(eventDTOs, w, lc)
}

func addNewEvent(
	e model.Event, ctx context.Context,
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

func getAllEvents(dbClient interfaces.DBClient) ([]model.Event, error) {
	events, err := dbClient.GetEvents()
	if err != nil {
		return []model.Event{}, err
	}
	return events, nil
}

func getEventById(id string, dbClient interfaces.DBClient) (event model.Event, err error) {
	event, err = dbClient.GetEventById(id)
	if err != nil {
		if err == db.ErrNotFound {
			err = errors.NewErrEventNotFound(id)
		}
		return event, err
	}

	return event, nil
}

func getEventCount(dbClient interfaces.DBClient) (int, error) {
	count, err := dbClient.GetEventCount()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func getEventCountByDeviceId(deviceId string, dbClient interfaces.DBClient) (int, error) {
	count, err := dbClient.GetEventCountByDeviceId(deviceId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func getAllEventsByDeviceId(deviceId string, dbClient interfaces.DBClient) ([]model.Event, error) {
	events, err := dbClient.GetEventsByDeviceId(deviceId)
	if err != nil {
		return []model.Event{}, err
	}
	return events, nil
}