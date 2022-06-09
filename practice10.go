//trying to implement interfaces using shapes

package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectange struct {
	length  float64
	breadth float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r rectange) area() float64 {
	return r.length * r.breadth
}

func main() {
	c1 := circle{4.5}
	r1 := rectange{4, 5}
	shapes := []shape{c1, r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

}
