package volume

import (
	"math"
)

//Shape works a placeholder for various geometrical shapes
type Shape interface {
	getVolume() float64
}

//Volume will calculate the volume of a 3-d object.
func Volume(s Shape) float64 {
	return s.getVolume()
}

//Cube contains the params required to calculate the volume of a cube
type Cube struct {
	Side float64
}

func (c Cube) getVolume() float64 {
	return math.Pow(c.Side, 3)
}

//RectanglePrism contains the params required to calculate the volume of a rectangle prism
type RectanglePrism struct {
	Length float64
	Width  float64
	Height float64
}

func (r RectanglePrism) getVolume() float64 {
	return r.Length * r.Width * r.Height
}

//TrianglePrism contains the params required to calculate the volume of a triangle prism
type TrianglePrism struct {
	Base   float64
	Height float64
	Length float64
}

func (t TrianglePrism) getVolume() float64 {
	return t.Base * t.Length * t.Height / 2
}

//Cylinder contains the params required to calculate the volume of a cylinder
type Cylinder struct {
	Radius float64
	Height float64
}

func (c Cylinder) getVolume() float64 {
	return math.Pi * math.Pow(c.Radius, 2) * c.Height
}

//Pyramid contains the params required to calculate the volume of a pyramid
type Pyramid struct {
	Length float64
	Width  float64
	Height float64
}

func (p Pyramid) getVolume() float64 {
	return p.Length * p.Width * p.Height / 3
}

//Cone contains the params required to calculate the volume of a cone
type Cone struct {
	Radius float64
	Height float64
}

func (c Cone) getVolume() float64 {
	return math.Pi * math.Pow(c.Radius, 2) * c.Height / 3
}

//Sphere contains the params required to calculate the volume of a sphere
type Sphere struct {
	Radius float64
}

func (c Sphere) getVolume() float64 {
	return math.Pi * math.Pow(c.Radius, 3) * 4 / 3
}
