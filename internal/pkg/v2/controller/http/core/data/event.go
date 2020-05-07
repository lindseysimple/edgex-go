package data

import (
	"net/http"

	"github.com/edgexfoundry/edgex-go/internal/core/data/config"
	"github.com/edgexfoundry/edgex-go/internal/pkg/errorconcept"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/dtos/common/base"
	dto "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/dtos/coredata"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/common"
	handler "github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/core/data"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/infrastructure/interfaces"

	clients "github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/metadata"

	"github.com/edgexfoundry/go-mod-messaging/messaging"
)

func EventController(
	w http.ResponseWriter,
	r *http.Request,
	lc logger.LoggingClient,
	dbClient interfaces.DBClient,
	chEvents chan<- interface{},
	msgClient messaging.MessageClient,
	mdc metadata.DeviceClient,
	httpErrorHandler errorconcept.ErrorHandler,
	configuration *config.ConfigurationStruct) {

	if r.Body != nil {
		defer func() { _ = r.Body.Close() }()
	}

	ctx := r.Context()

	reader := common.NewRequestReader(r, configuration)
	events, err := reader.Read(r.Body, &ctx)
	if err != nil {
		httpErrorHandler.Handle(w, err, errorconcept.Default.InternalServerError)
		return
	}

	// map Event models to AddEventResponse DTOs
	var addResponses []dto.AddEventResponse
	for _, e := range events {
		newId, err := handler.AddNewEvent(e, ctx, lc, dbClient, chEvents, msgClient, mdc, configuration)
		var addEventResponse dto.AddEventResponse
		if err == nil {
			addEventResponse = dto.AddEventResponse{
				Response: base.Response{
					CorrelationID: e.CorrelationId,
					RequestID:     e.RequestId,
					StatusCode:    http.StatusAccepted,
				},
				ID: newId,
			}
		} else {
			addEventResponse = dto.AddEventResponse{
				Response: base.Response{
					CorrelationID: e.CorrelationId,
					RequestID:     e.RequestId,
					StatusCode:    http.StatusBadRequest,
				},
			}
		}
		addResponses = append(addResponses, addEventResponse)
	}
	if err != nil {
		httpErrorHandler.HandleManyVariants(
			w,
			err,
			[]errorconcept.ErrorConceptType{
				errorconcept.ValueDescriptors.NotFound,
				errorconcept.ValueDescriptors.Invalid,
				errorconcept.NewServiceClientHttpError(err),
			},
			errorconcept.Default.InternalServerError)
		return
	}

	w.Header().Set(clients.ContentType, clients.ContentTypeJSON)
	w.WriteHeader(http.StatusMultiStatus)
	common.Encode(addResponses, w, lc)
}
