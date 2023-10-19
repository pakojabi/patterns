package factory

import "fmt"

type Polygon interface {}
type Shape int

type polygon struct {
	sides, length int
}

const (
	Triangle Shape = iota
	Square
	Pentagon
	Hexagon
)

func NewPolygon(shape Shape) Polygon {
	switch shape {
	case Triangle:
		return &polygon{3, 10}
	case Square:
		return &polygon{4, 10}
	case Pentagon:
		return &polygon{5, 10}
	case Hexagon:
		return &polygon{6, 10}
	default:
		panic("not supported")
	}
}

func RunPrototypeFactory() {
	fmt.Println("Building a pentagon")
	polygon := NewPolygon(Pentagon)
	fmt.Println(polygon)
}


