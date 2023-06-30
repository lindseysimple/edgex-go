//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"

	"github.com/stretchr/testify/assert"
)

func TestCheckDuration(t *testing.T) {
	tests := []struct {
		name            string
		interval        string
		min             string
		result          bool
		errorExpected   bool
		expectedErrKind errors.ErrKind
	}{
		{"valid - interval is bigger than the minimum value", "1s", "10ms", true, false, ""},
		{"invalid - interval is bigger than the minimum value", "100us", "1ms", false, false, ""},
		{"invalid - parsing duration string failed", "INVALID", "1ms", false, true, errors.KindServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CheckDuration(tt.interval, tt.min)
			if tt.errorExpected {
				assert.Error(t, err)
				assert.NotEmpty(t, err.Error(), "Error message is empty")
				assert.Equal(t, tt.expectedErrKind, errors.Kind(err), "Error kind not as expected")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.result, result)
			}
		})
	}
}
