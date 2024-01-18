package structsmethodinterfaces

import (
	"testing"
	"math"
)

type Rectrangle struct {
	Width float64
	Height float64
}

func (r Rectrangle) Area() float64 {
	return r.Width * r.Height
}

type Circule struct {
	Radius float64
}

func (c Circule) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

type Shape interface {
	Area() float64
}


func TestPerimeter(t *testing.T) {
	rectangle := Rectrangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f want %.2f ", got, want)
	}
}

func TestArea(t *testing.T) {


	areaTest := []struct {
		name string
		shape Shape
		hasArea float64
	} {
		// {Rectrangle{12, 6}, 72.0},
		// {Circule{10}, 314.1592653589793},
		// {Triangle{12, 6}, 36.0},
		{name: "Rectangle", shape: Rectrangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circule{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			hasArea := tt.hasArea
			if got != hasArea {
				t.Errorf("%#v Got %g want %g", tt.shape, got, hasArea)
			}
		})
	}

}