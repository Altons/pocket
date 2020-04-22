package volume_test

import (
	"fmt"

	"github.com/Altons/pocket/pkg/volume"
)

func ExampleVolume() {
	c := volume.Cube{Side: 2}
	fmt.Println(volume.Volume(c))
	// Output: 8
}
