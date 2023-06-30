//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"fmt"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

// CheckDuration parses the ISO 8601 time duration string to Duration type
// and evaluates if the duration value is smaller than the suggested minimum duration string
func CheckDuration(value string, min string) (bool, errors.EdgeX) {
	valueDuration, err := time.ParseDuration(value)
	if err != nil {
		return false, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to parse the interval duration string %s to a duration time value", value), err)
	}
	minDuration, err := time.ParseDuration(min)
	if err != nil {
		return false, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to parse the minimum duration string %s to a duration time value", min), err)
	}

	if valueDuration < minDuration {
		// the duration value is smaller than the min
		return false, nil
	}
	return true, nil
}
