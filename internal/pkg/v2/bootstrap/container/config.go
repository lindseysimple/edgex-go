//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"github.com/edgexfoundry/edgex-go/internal/core/data/config"

	"github.com/edgexfoundry/go-mod-bootstrap/di"
)

// ConfigurationName contains the name of data's config.ConfigurationStruct implementation in the DIC.
var ConfigurationName = di.TypeInstanceToName(config.ConfigurationStruct{})

// ConfigurationFrom helper function queries the DIC and returns datas's config.ConfigurationStruct implementation.
func ConfigurationFrom(get di.Get) *config.ConfigurationStruct {
	return get(ConfigurationName).(*config.ConfigurationStruct)
}
