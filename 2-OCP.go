package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func CalcTotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}
