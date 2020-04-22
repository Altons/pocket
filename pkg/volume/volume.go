package volume

import (
	"math"
)

type Shape interface {
	getVolume() float64
}

func Volume(s Shape) float64 {
	return s.getVolume()
}

type Cube struct {
	Side float64
}

func (c Cube) getVolume() float64 {
	return math.Pow(c.Side, 3)
}

//RectPrism is a Rectangular prism
type RectanglePrism struct {
	Length float64
	Width  float64
	Height float64
}

func (r RectanglePrism) getVolume() float64 {
	return r.Length * r.Width * r.Height
}

//TrianPrism is a Triangular prism
type TrianglePrism struct {
	Base   float64
	Height float64
	Length float64
}

func (t TrianglePrism) getVolume() float64 {
	return t.Base * t.Length * t.Height / 2
}

//Cylinder
type Cylinder struct {
	Radius float64
	Height float64
}

func (c Cylinder) getVolume() float64 {
	return math.Pi * math.Pow(c.Radius, 2) * c.Height
}

type Pyramid struct {
	Length float64
	Width  float64
	Height float64
}

func (p Pyramid) getVolume() float64 {
	return p.Length * p.Width * p.Height / 3
}

//Cone
type Cone struct {
	Radius float64
	Height float64
}

func (c Cone) getVolume() float64 {
	return math.Pi * math.Pow(c.Radius, 2) * c.Height / 3
}

//Sphere
type Sphere struct {
	Radius float64
}

func (c Sphere) getVolume() float64 {
	return math.Pi * math.Pow(c.Radius, 3) * 4 / 3
}
