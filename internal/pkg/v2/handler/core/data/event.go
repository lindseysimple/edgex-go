//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/edgex-go/internal/pkg/correlation"

	v2container "github.com/edgexfoundry/edgex-go/internal/pkg/v2/bootstrap/container"
	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	model "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	msgTypes "github.com/edgexfoundry/go-mod-messaging/pkg/types"
)

// The AddEvent function accepts the new event model from the controller functions
// and invokes addEvent function in the infrustructure layer
func AddEvent(e model.Event, ctx context.Context) (string, error) {
	configuration := v2container.RegistryClients.ConfigClient
	dbClient := v2container.RegistryClients.DBClient

	err := checkDevice(e.Device, ctx)
	if err != nil {
		return "", err
	}

	// Add the event and readings to the database
	if configuration.Writable.PersistData {
		id, err := dbClient.AddEvent(e)
		if err != nil {
			return "", err
		}
		e.Id = id
		//savedEvent, err := dbClient.EventById(id)
		//if err == nil {
		//	e = savedEvent
		//}
	}

	putEventOnQueue(e, ctx) // Push event to message bus for App Services to consume

	return e.Id, nil
}

// Put event on the message queue to be processed by the rules engine
func putEventOnQueue(evt model.Event, ctx context.Context) {
	lc := v2container.RegistryClients.LoggingClient
	msgClient := v2container.RegistryClients.MessageClient
	configuration := v2container.RegistryClients.ConfigClient

	lc.Info("Putting event on message queue")

	evt.CorrelationId = correlation.FromContext(ctx)
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
