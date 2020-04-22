package temperature_test

import (
	"fmt"

	"github.com/Altons/pocket/pkg/temperature"
)

func ExampleConvert() {
	t := temperature.Temperature{}
	k, _ := t.Convert(10, "c", "k")
	fmt.Println(k)
	// Output: 283.15
}
