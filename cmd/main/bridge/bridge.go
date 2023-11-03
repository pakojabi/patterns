package bridge

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
	// specific vector rendering data
}

func (vr *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius ", radius)
}

type RasterRenderer struct {
	// specific raster rendering data
}

func (rr *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of raidus ", radius)
}

// Bridge

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer, radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

// main

func RunBridge() {
	renderer := VectorRenderer{}
	circle := NewCircle(&renderer, 3.4)

	circle.Draw()
}
