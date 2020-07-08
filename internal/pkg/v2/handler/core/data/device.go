//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"context"

	v2container "github.com/edgexfoundry/edgex-go/internal/pkg/v2/bootstrap/container"
)

// This function will be updated when CheckDevice in v2 core-metadata is available
func checkDevice(device string, ctx context.Context) error {
	mdc := v2container.RegistryClients.MetadataClient
	configuration := v2container.RegistryClients.ConfigClient

	if configuration.Writable.MetaDataCheck {
		_, err := mdc.CheckForDevice(ctx, device)
		if err != nil {
			return err
		}
	}
	return nil
}
