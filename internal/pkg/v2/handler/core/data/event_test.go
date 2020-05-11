package data

import (
	"context"
	"github.com/stretchr/testify/mock"
	"sync"
	"testing"

	"github.com/edgexfoundry/edgex-go/internal/core/data/config"
	dataMock "github.com/edgexfoundry/edgex-go/internal/pkg/v2/handler/core/data/mocks"
	dbMock "github.com/edgexfoundry/edgex-go/internal/pkg/v2/infrastructure/interfaces/mocks"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	contract "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/edgexfoundry/go-mod-messaging/messaging"
	msgTypes "github.com/edgexfoundry/go-mod-messaging/pkg/types"
)

func newAddEventMockDB(persist bool) *dbMock.DBClient {
	myMock := &dbMock.DBClient{}

	if persist {
		myMock.On("AddEvent", mock.Anything).Return("3c5badcb-2008-47f2-ba78-eb2d992f8422", nil)
	}

	return myMock
}

func TestAddEventWithPersistence(t *testing.T) {
	reset()

	// no need to mock this since it's all in process
	msgClient, _ := messaging.NewMessageClient(msgTypes.MessageBusConfig{
		PublishHost: msgTypes.HostInfo{
			Host:     "*",
			Protocol: "tcp",
			Port:     5563,
		},
		Type: "zero",
	})

	dbClientMock := newAddEventMockDB(true)
	chEvents := make(chan interface{}, 10)
	evt := contract.Event{Device: testDeviceName, Origin: testOrigin, Readings: buildReadings()}
	// wire up handlers to listen for device events
	bitEvents := make([]bool, 2)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go handleDomainEvents(bitEvents, chEvents, &wg, t)

	_, err := AddNewEvent(
		evt,
		context.Background(),
		logger.NewMockClient(),
		dbClientMock,
		chEvents,
		msgClient,
		dataMock.NewMockDeviceClient(),
		&config.ConfigurationStruct{
			Writable: config.WritableInfo{
				PersistData: true,
			},
		})

	if err != nil {
		t.Errorf(err.Error())
	}

	wg.Wait()
	for i, val := range bitEvents {
		if !val {
			t.Errorf("event not received in timely fashion, index %v, TestAddEventWithPersistence", i)
		}
	}

	dbClientMock.AssertExpectations(t)
}
