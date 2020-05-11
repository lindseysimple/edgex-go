//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"github.com/edgexfoundry/edgex-go/internal/core/data/config"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/controller/errorconcept"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/infrastructure/interfaces"

	"github.com/edgexfoundry/go-mod-bootstrap/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/di"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/metadata"

	"github.com/edgexfoundry/go-mod-messaging/messaging"
)

type Container struct {
	ConfigClient       *config.ConfigurationStruct
	DBClient           interfaces.DBClient
	ErrorHandlerClient *errorconcept.Handler
	EventChan          chan<- interface{}
	LoggingClient      logger.LoggingClient
	MessageClient      messaging.MessageClient
	MetadataClient     metadata.DeviceClient
}

var BootstrapContainer *di.Container
var RegistryClients Container

func Init() {
	dic := BootstrapContainer
	chEvents := make(chan interface{}, 100)
	dic.Update(di.ServiceConstructorMap{
		EventsChannelName: func(get di.Get) interface{} {
			return chEvents
		},
		ErrorHandlerName: func(get di.Get) interface{} {
			return errorconcept.NewErrorHandler(nil)
		},
	})

	RegistryClients = Container{
		ConfigClient:       ConfigurationFrom(dic.Get),
		DBClient:           DBClientFrom(dic.Get),
		ErrorHandlerClient: ErrorHandlerFrom(dic.Get),
		EventChan:          PublisherEventsChannelFrom(dic.Get),
		LoggingClient:      container.LoggingClientFrom(dic.Get),
		MessageClient:      MessagingClientFrom(dic.Get),
		MetadataClient:     MetadataDeviceClientFrom(dic.Get),
	}
}
