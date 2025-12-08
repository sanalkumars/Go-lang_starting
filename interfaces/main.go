package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	circle := Circle{radius: 7}

	printArea(rect)
	printArea(circle)
}
