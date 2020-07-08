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
func addReading(conn redis.Conn, r model.Reading) (id string, err error) {
	switch r.(type) {
	case model.BinaryReading:
		newReading, ok := r.(model.BinaryReading)
		if ok {
			if newReading.Created == 0 {
				newReading.Created = MakeTimestamp()
			}
			// check if id is a valid uuid
			if newReading.Id == "" {
				newReading.Id = uuid.New().String()
			} else {
				_, err = uuid.Parse(newReading.Id)
				if err != nil {
					return "", ErrInvalidObjectId
				}
			}
			_ = conn.Send("HMSET", redis.Args{}.Add(ReadingsCollection+":id:"+newReading.Id).AddFlat(&newReading)...)
			_ = conn.Send("ZADD", ReadingsCollection+":created", newReading.Created, newReading.Id)
			_ = conn.Send("ZADD", ReadingsCollection+":device:"+newReading.Device, newReading.Created, newReading.Id)
			return newReading.Id, nil
		} else {
			return "", ErrUnsupportedDatabase
		}

	case model.SimpleReading:
		newReading, ok := r.(model.SimpleReading)
		if ok {
			if newReading.Created == 0 {
				newReading.Created = MakeTimestamp()
			}
			// check if id is a valid uuid
			if newReading.Id == "" {
				newReading.Id = uuid.New().String()
			} else {
				_, err = uuid.Parse(newReading.Id)
				if err != nil {
					return "", ErrInvalidObjectId
				}
			}
			_ = conn.Send("HMSET", redis.Args{}.Add(ReadingsCollection+":id:"+newReading.Id).AddFlat(&newReading)...)
			_ = conn.Send("ZADD", ReadingsCollection+":created", newReading.Created, newReading.Id)
			_ = conn.Send("ZADD", ReadingsCollection+":device:"+newReading.Device, newReading.Created, newReading.Id)
			return newReading.Id, nil
		} else {
			return "", ErrUnsupportedDatabase
		}
	default:
		return "", ErrUnsupportedDatabase
	}
}
