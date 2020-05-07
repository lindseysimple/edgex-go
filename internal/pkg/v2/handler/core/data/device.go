//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"context"

	"github.com/edgexfoundry/edgex-go/internal/core/data/config"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/metadata"
)

func checkDevice(
	device string,
	ctx context.Context,
	mdc metadata.DeviceClient,
	configuration *config.ConfigurationStruct) error {

	if configuration.Writable.MetaDataCheck {
		_, err := mdc.CheckForDevice(ctx, device)
		if err != nil {
			return err
		}
	}
	return nil
}
