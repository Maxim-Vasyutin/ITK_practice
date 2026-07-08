package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimetr() float64
}

type Circle struct {
	r int
}

type Rectangle struct {
	width  int
	height int
}

// Circle
func (c *Circle) Area() float64 {
	return math.Pi * float64(c.r)
}

func (c *Circle) Perimetr() float64 {
	return 2 * math.Pi * float64(c.r)
}

// Rectangle
func (rctg *Rectangle) Area() float64 {
	return float64(rctg.width) * float64(rctg.height)
}

func (rctg *Rectangle) Perimetr() float64 {
	return 2 * (float64(rctg.width) + float64(rctg.height))
}

func main() {
	circle := Circle{2}
	rectangle := Rectangle{2, 3}

	fmt.Println("circle area = ", circle.Area())
	fmt.Println("rectangle area = ", rectangle.Area())
	fmt.Println("circle perimetr = ", circle.Perimetr())
	fmt.Println("rectangle perimetr = ", rectangle.Perimetr())
}
