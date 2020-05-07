package coredata

import (
	"fmt"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod"
	model "github.com/edgexfoundry/edgex-go/internal/pkg/v2/go-mod"
	"reflect"
	"testing"
)

var testReading = BaseReading{
	Device:        v2.TestDeviceName,
	Name:          v2.TestReadingName,
	SimpleReading: SimpleReading{
		ValueType: v2.TestValueType,
		Value:     v2.TestValue,
	},
}

func Test_toReadingModel(t *testing.T) {
	valid := testReading
	expected := model.SimpleReading{
		BaseReading:   model.BaseReading{
			Device: v2.TestDeviceName,
			Name:   v2.TestReadingName,
		},
		Value:         v2.TestValue,
		ValueType:     v2.TestValueType,
	}
	tests := []struct {
		name        string
		reading       BaseReading
		expectError bool
	}{
		{"valid Event", valid, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readingModel := toReadingModel(tt.reading, v2.TestDeviceName)
			if !reflect.DeepEqual(expected, readingModel) {
				fmt.Println(readingModel)
				t.Errorf("toReadingModel did not result in expected Reading Model.")
			}
		})
	}
}
