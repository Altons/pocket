package volume_test

import "github.com/Altons/pocket/pkg/volume"

var volumeTestCases = []struct {
	shape volume.Shape
	want  float64
}{
	{shape: volume.Cube{Side: 2}, want: 8},
	{shape: volume.RectanglePrism{Length: 2, Width: 2, Height: 2}, want: 8},
	{shape: volume.TrianglePrism{Base: 2, Height: 2, Length: 2}, want: 4},
	{shape: volume.Cylinder{Radius: 5, Height: 2}, want: 157.08},
	{shape: volume.Pyramid{Length: 3, Width: 3, Height: 3}, want: 9},
	{shape: volume.Cone{Radius: 5, Height: 2}, want: 52.36},
	{shape: volume.Sphere{Radius: 2}, want: 33.51},
}
