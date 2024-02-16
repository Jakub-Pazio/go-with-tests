package stuctsmethodsinterfaces

import "math"

type Shape interface {
	Area() float64
	Permiter() float64
}

type Rectange struct {
	Width  float64
	Height float64
}

func (r Rectange) Permiter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectange) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Permiter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
