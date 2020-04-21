//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package db

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/correlation/models/core/data"
	v2model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/models/coredata"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type Client struct {
	Pool          *redis.Pool // A thread-safe pool of connections to Redis
	BatchSize     int
	loggingClient logger.LoggingClient
}

// Add a new event
// UnexpectedError - failed to add to database
// NoValueDescriptor - no existing value descriptor for a reading in the event
//func (c *Client) AddEvent(e data.Event) (id string, err error) {
func (c *Client) AddEvent(e data.Event) (id string, err error) {
	conn := c.Pool.Get()
	defer conn.Close()

	if e.ID != "" {
		_, err = uuid.Parse(e.ID)
		if err != nil {
			return "", ErrInvalidObjectId
		}
	}
	return addEvent(conn, e)
}

// ************************** HELPER FUNCTIONS ***************************
func addEvent(conn redis.Conn, e data.Event) (id string, err error) {
	if e.Created == 0 {
		e.Created = MakeTimestamp()
	}

	if e.ID == "" {
		e.ID = uuid.New().String()
	}

	eventHashes := data.Event{
		Bytes:         e.Bytes,
		CorrelationId: e.CorrelationId,
		Checksum:      e.Checksum,
		Event: v2model.Event{
			ID:       e.ID,
			Pushed:   e.Pushed,
			Device:   e.Device,
			Created:  e.Created,
			Modified: e.Modified,
			Origin:   e.Origin,
		},
	}

	_ = conn.Send("MULTI")
	_ = conn.Send("HMSET", redis.Args{}.Add(EventsCollection+":id:"+e.ID).AddFlat(&eventHashes)...)
	_ = conn.Send("ZADD", EventsCollection+":by_created", e.Created, e.ID)
	_ = conn.Send("ZADD", EventsCollection+":by_pushed", e.Pushed, e.ID)
	_ = conn.Send("ZADD", EventsCollection+":by_device:"+e.Device, e.Created, e.ID)
	if e.Checksum != "" {
		_ = conn.Send("ZADD", EventsCollection+":by_checksum:"+e.Checksum, 0, e.ID)
	}

	rids := make([]interface{}, len(e.Readings)*2+1)
	rids[0] = EventsCollection + ":readings:" + e.ID
	for i, r := range e.Readings {
		newReading := r.(v2model.SimpleReading)
		newReading.Created = e.Created
		newReading.Device = e.Device
		if newReading.Id != "" {
			_, err = uuid.Parse(newReading.Id)
			if err != nil {
				return "", ErrInvalidObjectId
			}
		}
		id, err = addReading(conn, false, newReading)
		if err != nil {
			return id, err
		}
		rids[i*2+1] = 0
		rids[i*2+2] = id
	}
	if len(rids) > 1 {
		_ = conn.Send("ZADD", rids...)
	}

	_, err = conn.Do("EXEC")
	return e.ID, err
}

func (c *Client) GetEventById(id string) (event v2model.Event, err error) {
	conn := c.Pool.Get()
	defer conn.Close()

	event, err = eventByID(conn, id)
	if err != nil {
		return event, err
	}

	return event, nil
}

func eventByID(conn redis.Conn, id string) (event v2model.Event, err error) {
	obj, err := redis.Values(conn.Do("HGETALL", EventsCollection+":id:"+id))
	if err == redis.ErrNil {
		return event, ErrNotFound
	}
	if err != nil {
		return event, err
	}

	redis.ScanStruct(obj, &event)
	values, err := redis.Values(conn.Do("zrange", "v2:event:readings:"+id, 0, -1))
	var readingIDs []string
	redis.ScanSlice(values, &readingIDs)
	var readings = make([]v2model.ReadingInterface, len(readingIDs))
	for i, rid := range readingIDs {
		var s v2model.SimpleReading
		if r, err := redis.Values(conn.Do("HGETALL", ReadingsCollection+":id:"+rid)); err == nil {
			redis.ScanStruct(r, &s)
			readings[i] = s
		} else {
			return event, err
		}
	}

	event.Readings = readings
	return event, err
}
