//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package container

import "github.com/edgexfoundry/go-mod-bootstrap/di"

// EventsChannelName contains the name of the Events channel instance in the DIC.
var EventsChannelName = "CoreDataEventsChannel"

// PublisherEventsChannelFrom helper function queries the DIC and returns the Events channel instance used for
// publishing over the channel.
//
// NOTE If there is a need to obtain a consuming version of the channel create a new helper function which will get the
// channel from the container and cast it to a consuming channel. The type casting will aid in avoiding errors by
// restricting functionality.
func PublisherEventsChannelFrom(get di.Get) chan<- interface{} {
	return get(EventsChannelName).(chan interface{})
}
