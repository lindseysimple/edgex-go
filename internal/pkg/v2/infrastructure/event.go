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

	rids := make([]interface{}, len(e.Readings)*2+1)
	rids[0] = EventsCollection + ":readings:" + e.Id
	for i, r := range e.Readings {
		newReading := r.(model.SimpleReading)
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
	return e.Id, err
}
