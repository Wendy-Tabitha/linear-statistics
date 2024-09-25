package statistics

import (
	"reflect"
	"testing"
)

func TestMean(t *testing.T) {
	dataSet := []float64{10, 20, 30, 40, 50}
	input := Mean(dataSet)
	expected := 30.0

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Test Failed!!!!")
		t.FailNow()
	}
}
