package temperature_test

import (
	"testing"

	"github.com/Altons/pocket/pkg/temperature"
	"github.com/Altons/pocket/pkg/util"
)

func TestConvert(t *testing.T) {
	for _, tc := range temperatureTestCases {
		temp := temperature.Temperature{}
		got, err := temp.Convert(tc.value, tc.unitIn, tc.unitOut)
		if tc.expectError {
			// we expect error
			if err == nil {
				t.Fatalf("Convert(%f, %q,%q); expected error, got nil.",
					tc.value, tc.unitIn, tc.unitOut)
			}
		} else {
			// we do not expect error
			if err != nil {
				t.Fatalf("Convert(%f, %q, %q) returned unexpected error: %v",
					tc.value, tc.unitIn, tc.unitOut, err)
			}
			if util.Round(got, 0.01) != util.Round(tc.want, 0.01) {
				t.Fatalf("Convert(%f, %q,%q) = %f, want %f.",
					tc.value, tc.unitIn, tc.unitOut, got, tc.want)
			}

		}
	}
}
