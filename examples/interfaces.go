package main

import (
	"fmt"
	"math"
)

// An interface defines "All or nothing in go.." A type so must implement all of the methods defined in the interface
// To be considered a valid implementation, and also inferred by go compiler.
type geometry interface {
	area() float64
	perim() float64
}

type rect_2 struct {
	width, height float64
}

type circle struct {
	radius float64
}

// type triangle struct {
// 	base, height float64
// }

func (r rect_2) area() float64 {
	return r.height * r.width
}

func (r rect_2) perim() float64 {
	return 2 * r.width * 2 * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius

}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok { // Circle implicit implements geometry. So in this case is possible to use the type assertion
		fmt.Println("Circle with radius", c.radius)
	} else {
		fmt.Println("It is a rectangle")
	}
}

func ImplementingInterfaces() {
	r := rect_2{width: 3, height: 4}
	c := circle{radius: 5}
	// t := triangle{base: 10, height: 20}
	// Implicit interfaces..
	// Go Philosophy: "If it walks like a duck and it quacks like a duck, then it must be a duck.."
	measure(r)
	measure(c)
	// measure(t)

	detectCircle(r)
	detectCircle(c)
	// detectCircle(t)

}
