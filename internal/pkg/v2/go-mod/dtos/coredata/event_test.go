package coredata

import (
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/dtos/common/base"
	model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod/models/coredata"
	"reflect"
	"testing"
)

var testAddEvent = AddEventRequest{
	Request: base.Request{
		CorrelationID: v2.ExampleUUID,
		RequestID:     v2.ExampleUUID,
	},
	Device:   v2.TestDeviceName,
	Origin:   v2.TestOriginTime,
	Readings: nil,
}

func TestAddEventRequest_Validate(t *testing.T) {
	valid := testAddEvent
	noCoID := testAddEvent
	noCoID.CorrelationID = ""
	noReID := testAddEvent
	noReID.RequestID = ""
	noDevice := testAddEvent
	noDevice.Device = ""
	noOrigin := testAddEvent
	noOrigin.Origin = 0
	tests := []struct {
		name        string
		event       AddEventRequest
		expectError bool
	}{
		{"valid AddEventRequest", valid, false},
		{"invalid AddEventRequest, no CorrelationId", noCoID, true},
		{"invalid AddEventRequest, no CorrelationId", noReID, true},
		{"invalid AddEventRequest, no Device", noDevice, true},
		{"invalid AddEventRequest, no Origin", noOrigin, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.event.Validate()
		})
	}
}

func TestAddEvent_UnmarshalJSON(t *testing.T) {
	valid := testAddEvent
	resultTestBytes, _ := json.Marshal(testAddEvent)
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		addEvent AddEventRequest
		args     args
		wantErr  bool
	}{
		{"unmarshal AddEventRequest with success", valid, args{resultTestBytes}, false},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.addEvent
			if err := tt.addEvent.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("AddEventRequest.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				// if the bytes did unmarshal, make sure they unmarshaled to correct Event by comparing it to expected results
				var unmarshaledResult = tt.addEvent
				if err == nil && !reflect.DeepEqual(expected, unmarshaledResult) {
					fmt.Println(unmarshaledResult)
					t.Errorf("Unmarshal did not result in expected AddEventRequest.")
				}
			}
		})
	}
}

func Test_ToEventModels(t *testing.T) {
	valid := []AddEventRequest{testAddEvent}
	expectedEventModel := []model.Event{{
		CorrelationId: v2.ExampleUUID,
		Device:        v2.TestDeviceName,
		Origin:        v2.TestOriginTime,
		Readings:      []model.Reading{},
	}}
	tests := []struct {
		name        string
		addEvents   []AddEventRequest
		expectError bool
	}{
		{"valid AddEventRequest", valid, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readingModel := ToEventModels(tt.addEvents)
			if !reflect.DeepEqual(expectedEventModel, readingModel) {
				fmt.Println(readingModel)
				t.Errorf("ToEventModels did not result in expected Event model.")
			}
		})
	}
}
