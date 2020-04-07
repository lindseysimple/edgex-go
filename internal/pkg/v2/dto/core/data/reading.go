//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package data

import (
	dtoBase "github.com/edgexfoundry/edgex-go/internal/pkg/v2/dto/common/base"
	model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/model/core/data"
)

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
	Reading          model.Reading
}
