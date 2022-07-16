package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
	
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

func totalArea(shapes ...Shape) float64 {
	var total float64
	for _, c := range shapes {
		total += c.area()
	}
	return total
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w
}

// Embeded struct function
func (c *Circle) area() float64 {
	return math.Pi * (c.r * c.r)
}

func main() {
	// Struct initialization via 'new' mem allocator
	c := new(Circle)
	fmt.Println(*c)
	fmt.Println(c.area())

	c2 := Circle{x: 2, y: 1, r: 5}
	fmt.Println(c2)

	// Struct initialization via pointer to function.
	c3 := &Circle{2, 1, 5}
	fmt.Println(*c3)
	fmt.Println(c3.area())
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a * a + b * b)
}