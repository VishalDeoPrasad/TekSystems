package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Ractangle struct {
	height float64
	width  float64
}

// both implementing the Shape interface
//with their respective Area() methods.

func (c Circle) Area() float64 {
	return 3.14 * c.radius

}
func (r Ractangle) Area() float64 {
	return r.height * r.width
}

func main() {
	c := Circle{
		radius: 4,
	}
	fmt.Println("Area of Circle:",c.Area())

	r := Ractangle{
		height: 4,
		width: 6,
	}

	fmt.Println("Area of Rectangle:",r.Area())

}
