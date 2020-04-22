package volume_test

import (
	"testing"

	"github.com/Altons/pocket/pkg/util"
	"github.com/Altons/pocket/pkg/volume"
)

func TestVolume(t *testing.T) {
	for _, tt := range volumeTestCases {
		if got := volume.Volume(tt.shape); util.Round(got, 0.01) != util.Round(tt.want, 0.01) {
			t.Errorf("Volume(%v) = %v, want %v", tt.shape, util.Round(got, 0.01), util.Round(tt.want, 0.01))
		}
	}
}
