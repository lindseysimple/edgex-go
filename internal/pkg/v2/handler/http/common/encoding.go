//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"encoding/json"
	"net/http"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
)

func Encode(i interface{}, w http.ResponseWriter, LoggingClient logger.LoggingClient) {
	w.Header().Add(clients.ContentType, clients.ContentTypeJSON)

	enc := json.NewEncoder(w)
	err := enc.Encode(i)
	// Problems encoding
	if err != nil {
		LoggingClient.Error("Error encoding the data: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
