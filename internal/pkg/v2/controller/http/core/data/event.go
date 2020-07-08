package data

import (
	"net/http"

	v2container "github.com/edgexfoundry/edgex-go/internal/pkg/v2/bootstrap/container"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/controller/errorconcept"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/common"
	handler "github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/core/data"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	commonDTO "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	requestDTO "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	respDTO "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
)

func EventController(
	w http.ResponseWriter,
	r *http.Request) {

	if r.Body != nil {
		defer func() { _ = r.Body.Close() }()
	}

	// retrieve all the registry clients
	httpErrorHandler := v2container.RegistryClients.ErrorHandlerClient
	lc := v2container.RegistryClients.LoggingClient
	configuration := v2container.RegistryClients.ConfigClient

	ctx := r.Context()

	reader := common.NewRequestReader(r, configuration)
	addEventReqDTOs, err := reader.Read(r.Body, &ctx)
	if err != nil {
		httpErrorHandler.HttpHandleOneVariant(w, err, errorconcept.NewErrContractInvalidError(err), errorconcept.Default.InternalServerError)
		return
	}
	events := requestDTO.AddEventReqToEventModels(addEventReqDTOs)

	// map Event models to AddEventResponse DTOs
	var addResponses []respDTO.AddEventResponse
	for i, e := range events {
		newId, err := handler.AddEvent(e, ctx)

		var addEventResponse respDTO.AddEventResponse
		// get the requestID from AddEventRequestDTO
		reqId := addEventReqDTOs[i].RequestID

		if err == nil {
			addEventResponse = respDTO.AddEventResponse{
				BaseResponse: commonDTO.BaseResponse{
					RequestID:  reqId,
					Message:    "Add events successfully",
					StatusCode: http.StatusCreated,
				},
				ID: newId,
			}
		} else {
			singleHTTPResponse := httpErrorHandler.HandleOneVariant(
				err,
				errorconcept.NewServiceClientHttpError(err),
				errorconcept.Default.InternalServerError)
			addEventResponse = respDTO.AddEventResponse{
				BaseResponse: commonDTO.BaseResponse{
					RequestID:  reqId,
					StatusCode: http.StatusInternalServerError,
					Message:    singleHTTPResponse.Message,
				},
			}
		}
		addResponses = append(addResponses, addEventResponse)
	}

	w.Header().Set(clients.ContentType, clients.ContentTypeJSON)
	w.WriteHeader(http.StatusMultiStatus)
	common.Encode(addResponses, w, lc)
}
