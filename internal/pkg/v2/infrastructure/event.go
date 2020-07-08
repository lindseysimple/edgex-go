//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package infrastructure

import (
	model "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

// ************************** DB HELPER FUNCTIONS ***************************
func addEvent(conn redis.Conn, e model.Event) (id string, err error) {
	if e.Created == 0 {
		e.Created = MakeTimestamp()
	}

	if e.Id == "" {
		e.Id = uuid.New().String()
	}

	eventHashes := model.Event{
		CorrelationId: e.CorrelationId,
		Checksum:      e.Checksum,
		Id:            e.Id,
		Pushed:        e.Pushed,
		Device:        e.Device,
		Created:       e.Created,
		Modified:      e.Modified,
		Origin:        e.Origin,
	}

	_ = conn.Send("MULTI")
	_ = conn.Send("HMSET", redis.Args{}.Add(EventsCollection+":id:"+e.Id).AddFlat(&eventHashes)...)
	_ = conn.Send("ZADD", EventsCollection+":created", e.Created, e.Id)
	_ = conn.Send("ZADD", EventsCollection+":pushed", e.Pushed, e.Id)
	_ = conn.Send("ZADD", EventsCollection+":device:"+e.Device, e.Created, e.Id)
	if e.Checksum != "" {
		_ = conn.Send("ZADD", EventsCollection+":checksum:"+e.Checksum, 0, e.Id)
	}
	// add reading ids as sorted set under each event id
	rids := make([]interface{}, len(e.Readings)*2+1)
	rids[0] = EventsCollection + ":readings:" + e.Id
	for i, r := range e.Readings {
		id, err = addReading(conn, r)
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
	return e.Id, err
}

func eventByID(conn redis.Conn, id string) (event model.Event, err error) {
	obj, err := redis.Values(conn.Do("HGETALL", EventsCollection+":id:"+id))
	if err == redis.ErrNil {
		return event, ErrNotFound
	}
	if err != nil {
		return event, err
	}

	redis.ScanStruct(obj, &event)
	values, err := redis.Values(conn.Do("zrange", EventsCollection+":readings:"+id, 0, -1))
	var readingIDs []string
	redis.ScanSlice(values, &readingIDs)
	var readings = make([]model.Reading, len(readingIDs))
	for i, rid := range readingIDs {
		var s model.SimpleReading
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
