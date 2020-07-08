//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	contract "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/google/uuid"
)

var testEvent contract.Event

const (
	testDeviceName string = "Test Device"
	testOrigin     int64  = 123456789
	testUUIDString string = "ca93c8fa-9919-4ec5-85d3-f81b2b6a7bc1"
)

// Supporting methods
// Reset() re-initializes dependencies for each test
func reset() {
	testEvent.Id = testUUIDString
	testEvent.Device = testDeviceName
	testEvent.Origin = testOrigin
	testEvent.Readings = buildReadings()
}

func buildReadings() []contract.Reading {
	ticks := db.MakeTimestamp()

	r1 := contract.SimpleReading{
		BaseReading: contract.BaseReading{
			Id:       uuid.New().String(),
			Pushed:   ticks,
			Created:  ticks,
			Origin:   testOrigin,
			Modified: ticks,
			Device:   testDeviceName,
			Name:     "Temperature",
			Labels:   []string{"Fahrenheit"},
		},
		Value:         "45",
		ValueType:     "Float32",
		FloatEncoding: "Base64",
	}

	r2 := contract.BinaryReading{
		BaseReading: contract.BaseReading{
			Id:       uuid.New().String(),
			Pushed:   ticks,
			Created:  ticks,
			Origin:   testOrigin,
			Modified: ticks,
			Device:   testDeviceName,
			Name:     "FileData",
			Labels:   []string{"text"},
		},
		BinaryValue: []byte("1010"),
		MediaType:   "file",
	}

	var readings []contract.Reading
	readings = append(readings, r1, r2)
	return readings
}
