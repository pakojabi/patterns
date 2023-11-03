package decorator

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (s *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", s.Radius)
}

type Square struct {
	Size float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square of size %f", s.Size)
}

// decorator: we embed the interface inside a struct that provides the color and it is also a shape
type ColoredShape struct {
	Shape Shape
	Color string
}

func (s *ColoredShape) Render() string {
	return fmt.Sprintf("%s with color %s", s.Shape.Render(), s.Color)
}

// decorators can be composed
type TransparencyShape struct {
	Shape Shape
	Transparency float32
}

func (s *TransparencyShape) Render() string {
	return fmt.Sprintf("%s with transparency %f", s.Shape.Render(), s.Transparency)
}

func RunDecorator() {
	sq1 := Square{Size: 2.0}
	csq1 := ColoredShape{Shape: &sq1, Color: "Red"}
	tcsq1 := TransparencyShape{Shape: &csq1, Transparency: 10.0}
	fmt.Println(sq1.Render())
	fmt.Println(csq1.Render())
	fmt.Println(tcsq1.Render())
}



