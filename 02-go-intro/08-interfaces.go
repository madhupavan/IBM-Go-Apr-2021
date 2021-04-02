package main

import (
	"fmt"
	"math"
)

type Rect struct {
	width, height float32
}

func (r Rect) area() float32 {
	return r.width * r.height
}

func (r Rect) perimeter() float32 {
	return 2*r.width + 2*r.height
}

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}

type Shape interface {
	area() float32
	perimeter() float32
}

func PrintDimentions(s Shape) {
	fmt.Println("Area : ", s.area())
	fmt.Println("Perimeter : ", s.perimeter())
}

func main() {
	r := Rect{10, 20}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())

	c := Circle{10}
	fmt.Println(c.area())
	fmt.Println((c.perimeter()))

	PrintDimentions(r)
	PrintDimentions(c)

}
