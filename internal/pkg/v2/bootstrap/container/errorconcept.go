//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/controller/errorconcept"

	"github.com/edgexfoundry/go-mod-bootstrap/di"
)

// ErrorHandler contains the name of the errorconcept.Handler implementation in the DIC.
var ErrorHandlerName = di.TypeInstanceToName(errorconcept.Handler{})

// ErrorHandlerFrom helper function queries the DIC and returns the errorconcept.Handler implementation.
func ErrorHandlerFrom(get di.Get) *errorconcept.Handler {
	return get(ErrorHandlerName).(*errorconcept.Handler)
}
