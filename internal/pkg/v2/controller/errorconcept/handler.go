//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package errorconcept

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"net/http"
)

type ErrorConceptType interface {
	httpErrorCode() int
	isA(err error) bool
	message(err error) string
}

type Handler struct {
	logger logger.LoggingClient
}

type SingleHTTPResponse struct {
	Status  int
	Message string
}

type ErrorHandler interface {
	Handle(err error, ec ErrorConceptType) SingleHTTPResponse
	HandleOneVariant(err error, allowableError ErrorConceptType, defaultError ErrorConceptType) SingleHTTPResponse
	HttpDirectHandle(w http.ResponseWriter, err error, ec ErrorConceptType)
	HttpHandleOneVariant(w http.ResponseWriter, err error, allowableError ErrorConceptType, defaultError ErrorConceptType)
}

func NewErrorHandler(l logger.LoggingClient) ErrorHandler {
	h := Handler{l}
	return &h
}

// Handle applies the specified error and error concept to the single response
func (e *Handler) Handle(err error, ec ErrorConceptType) SingleHTTPResponse {
	message := ec.message(err)
	e.logger.Error(message)
	return SingleHTTPResponse{
		Status:  ec.httpErrorCode(),
		Message: message,
	}
}

// HandleOneVariant applies general error-handling with a single allowable error and a default error to be used as a
// fallback when none of the allowable errors are matched
func (e *Handler) HandleOneVariant(err error, allowableError ErrorConceptType, defaultError ErrorConceptType) SingleHTTPResponse {
	message := allowableError.message(err)
	e.logger.Error(message)
	return SingleHTTPResponse{
		Status:  allowableError.httpErrorCode(),
		Message: message,
	}

}

// Handle applies the specified error and error concept to the HTTP response writer
func (e *Handler) HttpDirectHandle(w http.ResponseWriter, err error, ec ErrorConceptType) {
	message := ec.message(err)
	e.logger.Error(message)
	http.Error(w, message, ec.httpErrorCode())
}

// HandleOneVariant applies general error-handling with a single allowable error and a default error to be used as a
// fallback when none of the allowable errors are matched, call HttpDirectHandle to sent a direct HTTP response
func (e *Handler) HttpHandleOneVariant(w http.ResponseWriter, err error, allowableError ErrorConceptType, defaultError ErrorConceptType) {
	if allowableError != nil && allowableError.isA(err) {
		e.HttpDirectHandle(w, err, allowableError)
		return
	}
	e.HttpDirectHandle(w, err, defaultError)
}
