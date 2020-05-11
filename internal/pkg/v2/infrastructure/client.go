//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/google/uuid"
	"sync"
	"time"

	model "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"

	"github.com/gomodule/redigo/redis"
)

var currClient *CoreDataClient // a singleton so Readings can be de-referenced
var once sync.Once

var deleteReadingsChannel = make(chan string, 50)
var deleteEventsChannel = make(chan string, 50)

type CoreDataClient struct {
	Pool      *redis.Pool // A thread-safe pool of connections to Redis
	BatchSize int
	logger    logger.LoggingClient
}

func NewCoreDataClient(config Configuration, logger logger.LoggingClient) (*CoreDataClient, error) {
	var err error
	dc, err := NewClient(config, logger)
	if err != nil {
		return nil, err
	}
	dc.logger = logger
	// Background process for deleting device readings and events.
	// This only needs to be running for core-data since this is the service responsible for handling the deletion
	// of events
	go dc.AsyncDeleteEvents()
	go dc.AsyncDeleteReadings()

	return dc, err
}

// Return a pointer to the Redis client
func NewClient(config Configuration, lc logger.LoggingClient) (*CoreDataClient, error) {

	once.Do(func() {
		connectionString := fmt.Sprintf("%s:%d", config.Host, config.Port)
		opts := []redis.DialOption{
			redis.DialPassword(config.Password),
			redis.DialConnectTimeout(time.Duration(config.Timeout) * time.Millisecond),
		}

		dialFunc := func() (redis.Conn, error) {
			conn, err := redis.Dial(
				"tcp", connectionString, opts...,
			)
			if err != nil {
				return nil, fmt.Errorf("Could not dial Redis: %s", err)
			}
			return conn, nil
		}
		// Default the batch size to 1,000 if not set
		batchSize := 1000
		if config.BatchSize != 0 {
			batchSize = config.BatchSize
		}
		currClient = &CoreDataClient{
			Pool: &redis.Pool{
				IdleTimeout: 0,
				/* The current implementation processes nested structs using concurrent connections.
				 * With the deepest nesting level being 3, three shall be the number of maximum open
				 * idle connections in the pool, to allow reuse.
				 * TODO: Once we have a concurrent benchmark, this should be revisited.
				 * TODO: Longer term, once the objects are clean of external dependencies, the use
				 * of another serializer should make this moot.
				 */
				MaxIdle: 10,
				Dial:    dialFunc,
			},
			BatchSize: batchSize,
			logger:    lc,
		}
	})
	return currClient, nil
}

// CloseSession closes the connections to Redis
func (c *CoreDataClient) CloseSession() {
	c.Pool.Close()
	close(deleteEventsChannel)
	close(deleteReadingsChannel)
	currClient = nil
	once = sync.Once{}
}

// Add a new event
// UnexpectedError - failed to add to database
// NoValueDescriptor - no existing value descriptor for a reading in the event
func (c *CoreDataClient) AddEvent(e model.Event) (id string, err error) {
	conn := c.Pool.Get()
	defer conn.Close()

	if e.Id != "" {
		_, err = uuid.Parse(e.Id)
		if err != nil {
			return "", ErrInvalidObjectId
		}
	}
	return addEvent(conn, e)
}

// DeleteEventsByDevice Delete events and readings associated with the specified deviceID
func (c *CoreDataClient) DeleteEventsByDevice(deviceId string) (int, error) {
	err := c.DeleteReadingsByDevice(deviceId)
	if err != nil {
		return 0, err
	}

	conn := c.Pool.Get()
	defer conn.Close()

	ids, err := redis.Strings(conn.Do("ZRANGE", db.EventsCollection+":device:"+deviceId, 0, -1))
	if err != nil {
		return 0, err
	}

	err = conn.Send("MULTI")
	if err != nil {
		return 0, err
	}

	for _, id := range ids {
		err = conn.Send("RENAME", id, DeletedEventsCollection+":"+id)
		if err != nil {
			return 0, err
		}
	}

	err = conn.Send("EXEC")
	deleteEventsChannel <- deviceId

	return len(ids), nil
}

// DeleteReadingsByDevice deletes readings associated with the specified device
func (c *CoreDataClient) DeleteReadingsByDevice(device string) error {
	conn := c.Pool.Get()
	defer conn.Close()

	ids, err := redis.Strings(conn.Do("ZRANGE", ReadingsCollection+":device:"+device, 0, -1))
	if err != nil {
		return err
	}

	err = conn.Send("MULTI")
	if err != nil {
		return err
	}

	for _, id := range ids {
		err = conn.Send("RENAME", id, DeletedReadingsCollection+":"+id)
		if err != nil {
			return err
		}
	}

	err = conn.Send("EXEC")
	deleteReadingsChannel <- device

	return nil
}

// AsyncDeleteEvents Handles the deletion of device events asynchronously. This function is expected to be running in
// a go-routine and works with the "DeleteEventsByDevice" function for better performance.
func (c *CoreDataClient) AsyncDeleteEvents() {
	conn := c.Pool.Get()
	defer conn.Close()

	c.logger.Debug("Starting background event deletion process")
	for {
		select {
		case device, ok := <-deleteEventsChannel:
			if ok {
				c.logger.Debug("Deleting event data for device: " + device)
				startTime := time.Now()
				c.deleteRenamedEvents(device)
				c.logger.Debug(fmt.Sprintf("Deleted events for device: '%s', elapsed time: %s", device, time.Since(startTime)))
			}
		}
	}
}

// AsyncDeleteReadings Handles the deletion of device readings asynchronously. This function is expected to be running
// in a go-routine and works with the "DeleteReadingsByDevice" function for better performance.
func (c *CoreDataClient) AsyncDeleteReadings() {
	c.logger.Debug("Starting background event deletion process")
	for {
		select {
		case device, ok := <-deleteReadingsChannel:
			if ok {
				c.logger.Debug("Deleting reading data for device: " + device)
				startTime := time.Now()
				c.deleteRenamedReadings(device)
				c.logger.Debug(fmt.Sprintf("Deleted readings for device: '%s', elapsed time: %s", device, time.Since(startTime)))
			}
		}
	}
}

// deleteRenamedEvents deletes all events associated with the specified device which have been marked for deletion.
// See the "DeleteEventsByDevice" function for details on the how events are marked for deletion(renamed)
func (c *CoreDataClient) deleteRenamedEvents(device string) {
	conn := c.Pool.Get()
	defer conn.Close()

	ids, err := redis.Strings(conn.Do("ZRANGE", EventsCollection+":device:"+device, 0, -1))
	if err != nil {
		c.logger.Error("Unable to delete event:" + err.Error())
		return
	}

	_, err = conn.Do("MULTI")
	if err != nil {
		c.logger.Error("Unable to start transaction for deletion:" + err.Error())
	}

	for _, id := range ids {
		_, err = conn.Do("GET", DeletedEventsCollection+":"+id)
		if err != nil {
			c.logger.Error("Unable to obtain events marked for deletion:" + err.Error())
		}
	}
	events, err := redis.Strings(conn.Do("EXEC"))

	queriesInQueue := 0
	var e model.Event
	_, err = conn.Do("MULTI")
	if err != nil {
		c.logger.Error("Unable to start batch processing for event deletion:" + err.Error())
	}

	for _, event := range events {
		err = json.Unmarshal([]byte(event), &e)
		if err != nil {
			c.logger.Error("Unable to marshal event: " + err.Error())
		}
		_ = conn.Send("UNLINK", DeletedEventsCollection+":"+e.Id)
		_ = conn.Send("ZREM", EventsCollection, e.Id)
		_ = conn.Send("ZREM", EventsCollection+":created", e.Id)
		_ = conn.Send("ZREM", EventsCollection+":device:"+e.Device, e.Id)
		_ = conn.Send("ZREM", EventsCollection+":pushed", e.Id)
		if e.Checksum != "" {
			_ = conn.Send("ZREM", EventsCollection+":checksum:"+e.Checksum, 0)
		}

		queriesInQueue++
		if queriesInQueue >= c.BatchSize {
			_, err = conn.Do("EXEC")
			queriesInQueue = 0
			if err != nil {
				c.logger.Error("Unable to execute batch deletion: " + err.Error())
				return
			}
		}
	}

	if queriesInQueue > 0 {
		_, err = conn.Do("EXEC")

		if err != nil {
			c.logger.Error("Unable to execute batch deletion: " + err.Error())
		}
	}
}

// deleteRenamedReadings deletes all readings associated with the specified device which have been marked for deletion.
// See the "DeleteReadingsByDevice" function for details on the how readings are marked for deletion(renamed)
func (c *CoreDataClient) deleteRenamedReadings(device string) {
	conn := c.Pool.Get()
	defer conn.Close()

	ids, err := redis.Strings(conn.Do("ZRANGE", db.ReadingsCollection+":device:"+device, 0, -1))
	if err != nil {
		c.logger.Error("Unable to delete reading:" + err.Error())
		return
	}

	_, err = conn.Do("MULTI")
	if err != nil {
		c.logger.Error("Unable to start transaction for deletion:" + err.Error())
	}

	for _, id := range ids {
		_, err = conn.Do("GET", DeletedReadingsCollection+":"+id)
		if err != nil {
			c.logger.Error("Unable to obtain readings marked for deletion:" + err.Error())
		}
	}
	readings, err := redis.Strings(conn.Do("EXEC"))

	queriesInQueue := 0

	_, err = conn.Do("MULTI")
	if err != nil {
		c.logger.Error("Unable to start batch processing for reading deletion:" + err.Error())
	}

	for _, reading := range readings {
		var r model.Reading
		var readingId string
		var readingName string
		var device string
		err = json.Unmarshal([]byte(reading), &r)
		if err != nil {
			c.logger.Error("Unable to marshal reading: " + err.Error())
		}
		switch r.(type) {
		case model.BinaryReading:
			readingStruct, ok := r.(model.BinaryReading)
			if ok {
				readingId = readingStruct.Id
				readingName = readingStruct.Name
				device = readingStruct.Device
			}
		case model.SimpleReading:
			readingStruct, ok := r.(model.SimpleReading)
			if ok {
				readingId = readingStruct.Id
				readingName = readingStruct.Name
				device = readingStruct.Device
			}
		}
		_ = conn.Send("UNLINK", DeletedReadingsCollection+":"+readingId)
		_ = conn.Send("ZREM", ReadingsCollection, readingId)
		_ = conn.Send("ZREM", ReadingsCollection+":created", readingId)
		_ = conn.Send("ZREM", ReadingsCollection+":device:"+device, readingId)
		_ = conn.Send("ZREM", ReadingsCollection+":name:"+readingName, readingId)
		queriesInQueue++

		if queriesInQueue >= c.BatchSize {
			_, err = conn.Do("EXEC")
			queriesInQueue = 0
			if err != nil {
				c.logger.Error("Unable to execute batch deletion: " + err.Error())
				return
			}
		}
	}

	if queriesInQueue > 0 {
		_, err = conn.Do("EXEC")

		if err != nil {
			c.logger.Error("Unable to execute batch deletion: " + err.Error())
		}
	}
}
