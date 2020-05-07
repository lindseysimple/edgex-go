//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package mapper

import (
	dtos "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/dtos/coredata"
	model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/models/coredata"
)

func ToEventDTO(event model.Event) dtos.Event {
	readings := make([]dtos.BaseReading, len(event.Readings))
	for i, r := range event.Readings {
		readings[i] = ToReadingDTO(r)
	}
	s := dtos.Event{
		ID:       event.ID,
		Pushed:   event.Pushed,
		Device:   event.Device,
		Created:  event.Created,
		Modified: event.Modified,
		Origin:   event.Origin,
		Readings: readings,
	}
	return s
}

func ToReadingDTO(r model.Reading) dtos.BaseReading {
	if _, ok := r.(model.SimpleReading); ok {
		r = r.(model.SimpleReading)
	} else {
		r = r.(model.BinaryReading)
	}
	readingDTO := dtos.BaseReading{
		Id:            r.(model.SimpleReading).Id,
		Pushed:        r.(model.SimpleReading).Pushed,
		Created:       r.(model.SimpleReading).Created,
		Origin:        r.(model.SimpleReading).Origin,
		Modified:      r.(model.SimpleReading).Modified,
		Device:        r.(model.SimpleReading).Device,
		Name:          r.(model.SimpleReading).Name,
		Labels:        r.(model.SimpleReading).Labels,
		BinaryReading: dtos.BinaryReading{},
		SimpleReading: dtos.SimpleReading{
			Value:     r.(model.SimpleReading).Value,
			ValueType: r.(model.SimpleReading).ValueType,
		},
	}
	return readingDTO
}
