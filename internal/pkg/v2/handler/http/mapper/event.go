//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package mapper

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/correlation/models/core/data"
	dtos "github.com/edgexfoundry/edgex-go/internal/pkg/v2/dtos/coredata"
	dataModel "github.com/edgexfoundry/edgex-go/internal/pkg/v2/models/coredata"
)

func ToEventContract(eventDTO dtos.AddEventRequest) data.Event {
	readings := make([]dataModel.ReadingInterface, len(eventDTO.Readings))
	for i, r := range eventDTO.Readings {
		readings[i] = ToReadingContract(r, eventDTO.Device)
	}
	s := data.Event{
		Event: dataModel.Event{
			Device:   eventDTO.Device,
			Readings: readings,
		},
	}
	return s
}

func ToReadingContract(r dtos.Reading, device string) dataModel.ReadingInterface {
	var readingModel dataModel.ReadingInterface
	br := dataModel.BaseReading{
		Device: device,//r.Device,
		Name:   r.Name,
		Labels: r.Labels,
	}
	if r.ValueType != "" {
		readingModel = dataModel.SimpleReading{
			BaseReading: br,
			Value:       r.Value,
			ValueType:   r.ValueType,
		}
	} else {
		readingModel = dataModel.BinaryReading{
			BaseReading: br,
			BinaryValue: r.BinaryValue,
			MediaType:   r.MediaType,
		}
	}
	return readingModel
}

func ToEventDTO(event dataModel.Event) dtos.Event {
	readings := make([]dtos.Reading, len(event.Readings))
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

func ToReadingDTO(r dataModel.ReadingInterface) dtos.Reading {
	if _, ok := r.(dataModel.SimpleReading);ok {
		r = r.(dataModel.SimpleReading)
	} else {
		r = r.(dataModel.BinaryReading)
	}
	readingDTO := dtos.Reading{
		Id:            r.(dataModel.SimpleReading).Id,
		Pushed:        r.(dataModel.SimpleReading).Pushed,
		Created:       r.(dataModel.SimpleReading).Created,
		Origin:        r.(dataModel.SimpleReading).Origin,
		Modified:      r.(dataModel.SimpleReading).Modified,
		Device:        r.(dataModel.SimpleReading).Device,
		Name:          r.(dataModel.SimpleReading).Name,
		Labels:        r.(dataModel.SimpleReading).Labels,
		BinaryReading: dtos.BinaryReading{},
		SimpleReading: dtos.SimpleReading{
			Value: r.(dataModel.SimpleReading).Value,
			ValueType:r.(dataModel.SimpleReading).ValueType,
		},
	}
	return readingDTO
}
