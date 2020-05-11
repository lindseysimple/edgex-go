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

// Add a reading to the database
func addReading(conn redis.Conn, tx bool, r model.Reading) (id string, err error) {
	newReading := r.(model.SimpleReading)
	if newReading.Created == 0 {
		newReading.Created = MakeTimestamp()
	}
	if newReading.Id == "" {
		newReading.Id = uuid.New().String()
	}
	if err != nil {
		return newReading.Id, err
	}
	if tx {
		_ = conn.Send("MULTI")
	}
	_ = conn.Send("HMSET", redis.Args{}.Add(ReadingsCollection+":id:"+newReading.Id).AddFlat(&newReading)...)
	_ = conn.Send("ZADD", ReadingsCollection+":created", newReading.Created, newReading.Id)
	_ = conn.Send("ZADD", ReadingsCollection+":device:"+newReading.Device, newReading.Created, newReading.Id)
	if tx {
		_, err = conn.Do("EXEC")
	}

	return newReading.Id, err
}
