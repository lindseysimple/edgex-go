package data

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/mock"
	"testing"

	dbMock "github.com/edgexfoundry/edgex-go/internal/pkg/v2/infrastructure/interfaces/mocks"

	contract "github.com/edgexfoundry/go-mod-core-contracts/v2/models"
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

	dbClientMock := newAddEventMockDB(true)
	evt := contract.Event{Device: testDeviceName, Origin: testOrigin, Readings: buildReadings()}

	_, err := AddEvent(evt, context.Background())
	if err != nil {
		t.Errorf(err.Error())
	}

	dbClientMock.AssertExpectations(t)
}

func TestAddEventNoPersistence(t *testing.T) {
	reset()

	dbClientMock := newAddEventMockDB(false)
	evt := contract.Event{Device: testDeviceName, Origin: testOrigin, Readings: buildReadings()}

	newId, err := AddEvent(evt, context.Background())

	if err != nil {
		t.Errorf(err.Error())
	}
	if bson.IsObjectIdHex(newId) {
		t.Errorf("unexpected bson id %s received", newId)
	}

	dbClientMock.AssertExpectations(t)
}
