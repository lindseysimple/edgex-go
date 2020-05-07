//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package db

import (
	"fmt"


	//"github.com/edgexfoundry/edgex-go/internal/pkg/db"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

var currClient *Client // a singleton so Readings can be de-referenced
var once sync.Once

type CoreDataClient struct {
	*Client
	logger logger.LoggingClient
}

func NewCoreDataClient(config Configuration, logger logger.LoggingClient) (*CoreDataClient, error) {
	var err error
	dc := &CoreDataClient{}
	dc.Client, err = NewClient(config, logger)
	if err != nil {
		return nil, err
	}
	dc.logger = logger
	// Background process for deleting device readings and events.
	// This only needs to be running for core-data since this is the service responsible for handling the deletion
	// of events
	// go dc.AsyncDeleteEvents()
	// go dc.AsyncDeleteReadings()

	return dc, err
}

// Return a pointer to the Redis client
func NewClient(config Configuration, lc logger.LoggingClient) (*Client, error) {
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
		currClient = &Client{
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
			BatchSize:     batchSize,
			loggingClient: lc,
		}
	})
	return currClient, nil
}

// CloseSession closes the connections to Redis
func (c *Client) CloseSession() {
	c.Pool.Close()
	//close(deleteEventsChannel)
	//close(deleteReadingsChannel)
	currClient = nil
	once = sync.Once{}
}
