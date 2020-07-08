//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/edgexfoundry/edgex-go/internal/pkg/bootstrap/interfaces"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/bootstrap/container"
	db "github.com/edgexfoundry/edgex-go/internal/pkg/v2/infrastructure"
	dbInterfaces "github.com/edgexfoundry/edgex-go/internal/pkg/v2/infrastructure/interfaces"

	bootstrapContainer "github.com/edgexfoundry/go-mod-bootstrap/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/bootstrap/startup"
	bootstrapConfig "github.com/edgexfoundry/go-mod-bootstrap/config"
	"github.com/edgexfoundry/go-mod-bootstrap/di"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
)

// httpServer defines the contract used to determine whether or not the http httpServer is running.
type httpServer interface {
	IsRunning() bool
}

// Database contains references to dependencies required by the database bootstrap implementation.
type Database struct {
	httpServer httpServer
	database   interfaces.Database
	isCoreData bool
}

// NewDatabaseForCoreData is a factory method that returns an initialized Database receiver struct.
func NewDatabaseForCoreData(httpServer httpServer, database interfaces.Database) Database {
	return Database{
		httpServer: httpServer,
		database:   database,
		isCoreData: true,
	}
}

// Return the dbClient interface
func (d Database) newDBClient(
	lc logger.LoggingClient,
	credentials bootstrapConfig.Credentials) (dbInterfaces.DBClient, error) {
	databaseInfo := d.database.GetDatabaseInfo()["Primary"]

	return db.NewCoreDataClient(
		db.Configuration{
			Host: databaseInfo.Host,
			Port: databaseInfo.Port,
		},
		lc)
}

// BootstrapHandler fulfills the BootstrapHandler contract and initializes the database.
func (d Database) BootstrapHandler(
	ctx context.Context,
	wg *sync.WaitGroup,
	startupTimer startup.Timer,
	dic *di.Container) bool {

	lc := bootstrapContainer.LoggingClientFrom(dic.Get)

	// get database credentials.
	var credentials bootstrapConfig.Credentials
	for startupTimer.HasNotElapsed() {
		var err error
		credentials, err = bootstrapContainer.CredentialsProviderFrom(dic.Get).GetDatabaseCredentials(d.database.GetDatabaseInfo()["Primary"])
		if err == nil {
			break
		}
		lc.Warn(fmt.Sprintf("couldn't retrieve database credentials: %v", err.Error()))
		startupTimer.SleepForInterval()
	}

	// initialize database.
	var dbClient dbInterfaces.DBClient

	for startupTimer.HasNotElapsed() {
		var err error
		dbClient, err = d.newDBClient(lc, credentials)
		if err == nil {
			break
		}
		dbClient = nil
		lc.Warn(fmt.Sprintf("couldn't create database client: %v", err.Error()))
		startupTimer.SleepForInterval()
	}

	if dbClient == nil {
		return false
	}

	dic.Update(di.ServiceConstructorMap{
		container.DBClientInterfaceName: func(get di.Get) interface{} {
			return dbClient
		},
	})

	lc.Info("Database connected")
	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		for {
			// wait for httpServer to stop running (e.g. handling requests) before closing the database connection.
			if d.httpServer.IsRunning() == false {
				dbClient.CloseSession()
				break
			}
			time.Sleep(time.Second)
		}
		lc.Info("Database disconnected")
	}()

	return true
}
