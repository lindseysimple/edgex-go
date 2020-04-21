//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package db

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	v2model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/models/coredata"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

// Add a reading to the database
func addReading(conn redis.Conn, tx bool, r v2model.ReadingInterface) (id string, err error) {
	newReading := r.(v2model.SimpleReading)
	if newReading.Created == 0 {
		newReading.Created = db.MakeTimestamp()
	}

	if newReading.Id == "" {
		newReading.Id = uuid.New().String()
	}

	//m, err := marshalObject(r)
	if err != nil {
		return newReading.Id, err
	}

	if tx {
		_ = conn.Send("MULTI")
	}
	//_ = conn.Send("SET", newReading.Id, m)
	_ = conn.Send("HMSET", redis.Args{}.Add(ReadingsCollection+":id:"+newReading.Id).AddFlat(&newReading)...)
	//_ = conn.Send("ZADD", db.ReadingsCollection, 0, newReading.Id)
	_ = conn.Send("ZADD", ReadingsCollection+":created", newReading.Created, newReading.Id)
	_ = conn.Send("ZADD", ReadingsCollection+":device:"+newReading.Device, newReading.Created, newReading.Id)
	_ = conn.Send("ZADD", ReadingsCollection+":name:"+newReading.Name, newReading.Created, newReading.Id)
	if tx {
		_, err = conn.Do("EXEC")
	}

	return newReading.Id, err
}
