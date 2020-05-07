//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/edgexfoundry/edgex-go/internal/core/data/config"
	model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/models/coredata"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/infrastructure/interfaces"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/metadata"

	"github.com/edgexfoundry/go-mod-messaging/messaging"
	msgTypes "github.com/edgexfoundry/go-mod-messaging/pkg/types"
)

func AddNewEvent(
	e model.Event, ctx context.Context,
	lc logger.LoggingClient,
	dbClient interfaces.DBClient,
	chEvents chan<- interface{},
	msgClient messaging.MessageClient,
	mdc metadata.DeviceClient,
	configuration *config.ConfigurationStruct) (string, error) {

	err := checkDevice(e.Device, ctx, mdc, configuration)
	if err != nil {
		return "", err
	}

	if configuration.Writable.ValidateCheck {
		lc.Debug("Validation enabled, parsing events")
		//for reading := range e.Readings {
		//	// Check value descriptor
		//	name := e.Readings[reading].Name
		//	vd, err := dbClient.ValueDescriptorByName(name)
		//	if err != nil {
		//		if err == db.ErrNotFound {
		//			return "", errors.NewErrValueDescriptorNotFound(name)
		//		} else {
		//			return "", err
		//		}
		//	}
		//	err = isValidValueDescriptor(vd, e.Readings[reading])
		//	if err != nil {
		//		return "", err
		//	}
		//}
	}

	// Add the event and readings to the database
	if configuration.Writable.PersistData {
		id, err := dbClient.AddEvent(e)
		if err != nil {
			return "", err
		}
		e.ID = id
	}

	putEventOnQueue(e, ctx, lc, msgClient, configuration) // Push event to message bus for App Services to consume
	chEvents <- DeviceLastReported{e.Device}              // update last reported connected (device)
	chEvents <- DeviceServiceLastReported{e.Device}       // update last reported connected (device service)

	return e.ID, nil
}

// Put event on the message queue to be processed by the rules engine
func putEventOnQueue(
	evt model.Event,
	ctx context.Context,
	lc logger.LoggingClient,
	msgClient messaging.MessageClient,
	configuration *config.ConfigurationStruct) {

	lc.Info("Putting event on message queue")

	// evt.CorrelationId = correlation.FromContext(ctx)
	var data []byte
	var err error
	// Re-marshal JSON content into bytes.
	if clients.FromContext(ctx, clients.ContentType) == clients.ContentTypeJSON {
		data, err = json.Marshal(evt)
		if err != nil {
			lc.Error(fmt.Sprintf("error marshaling event: %+v", evt))
			return
		}
	}

	msgEnvelope := msgTypes.NewMessageEnvelope(data, ctx)
	err = msgClient.Publish(msgEnvelope, configuration.MessageQueue.Topic)
	if err != nil {
		lc.Error(fmt.Sprintf("Unable to send message for event: %+v %v", evt, err))
	} else {
		lc.Info(fmt.Sprintf(
			"Event Published on message queue. Topic: %s, Correlation-id: %s ",
			configuration.MessageQueue.Topic,
			msgEnvelope.CorrelationID,
		))
	}
}
