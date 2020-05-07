//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package coredata

import (
	dtoBase "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/dtos/common/base"
	model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/models/coredata"
)

// Reading type implements SimpleReading or BinaryReading.
// BaseReading object in the APIv2 specification.
type BaseReading struct {
	Id            string   `json:"id,omitempty" codec:"id,omitempty"`
	Pushed        int64    `json:"pushed,omitempty" codec:"pushed,omitempty"`   // When the data was pushed out of EdgeX (0 - not pushed yet)
	Created       int64    `json:"created,omitempty" codec:"created,omitempty"` // When the reading was created
	Origin        int64    `json:"origin,omitempty" codec:"origin,omitempty"`
	Modified      int64    `json:"modified,omitempty" codec:"modified,omitempty"`
	Device        string   `json:"device,omitempty" codec:"device,omitempty"`
	Name          string   `json:"name,omitempty" codec:"name,omitempty"`
	Labels        []string `json:"labels,omitempty" codec:"labels,omitempty"` // Custom labels assigned to a reading, added in the APIv2 specification.
	BinaryReading `json:",inline"`
	SimpleReading `json:",inline"`
}

// An event reading for a binary data type
// BinaryReading object in the APIv2 specification.
type BinaryReading struct {
	BinaryValue []byte `json:"binaryValue,omitempty" codec:"binaryValue,omitempty"` // Binary data payload
	MediaType   string `json:"mediaType,omitempty" codec:"mediaType,omitempty"`     // indicates what the content type of the binaryValue property is
}

// An event reading for a simple data type
// SimpleReading object in the APIv2 specification.
type SimpleReading struct {
	Value         string `json:"value,omitempty" codec:"value,omitempty"`                 // Device sensor data value
	ValueType     string `json:"valueType,omitempty" codec:"valueType,omitempty"`         // Indicates the datatype of the value property
	FloatEncoding string `json:"floatEncoding,omitempty" codec:"floatEncoding,omitempty"` // Indicates how a float value is encoded
}

// ReadingCountResponse defines the Response Content for GET reading count DTO.  This object and its properties correspond to the
// ReadingCountResponse object in the APIv2 specification.
type ReadingCountResponse struct {
	dtoBase.Response `json:",inline"`
	Count            int
}

// ReadingResponse defines the Response Content for GET reading DTO.  This object and its properties correspond to the
// ReadingResponse object in the APIv2 specification.
type ReadingResponse struct {
	dtoBase.Response `json:",inline"`
	Reading          BaseReading
}

// convert Reading DTO to ReadingInterface struct model
func toReadingModel(r BaseReading, device string) model.Reading {
	var readingModel model.Reading
	br := model.BaseReading{
		Device: device,
		Name:   r.Name,
		Labels: r.Labels,
	}
	if r.ValueType != "" {
		readingModel = model.SimpleReading{
			BaseReading: br,
			Value:       r.Value,
			ValueType:   r.ValueType,
		}
	} else {
		readingModel = model.BinaryReading{
			BaseReading: br,
			BinaryValue: r.BinaryValue,
			MediaType:   r.MediaType,
		}
	}
	return readingModel
}
