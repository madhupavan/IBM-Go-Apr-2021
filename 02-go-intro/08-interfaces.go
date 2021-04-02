package main

import (
	"fmt"
	"math"
	"sort"
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

type Ariable interface {
	area() float32
}

type Perimeterable interface {
	perimeter() float32
}

type StandardShape interface {
	Ariable
	Perimeterable
}

func PrintDimentions(s StandardShape) {
	fmt.Println("Area : ", s.area())
	fmt.Println("Perimeter : ", s.perimeter())
}

type CustomShape struct {
	size float32
}

func (cs CustomShape) area() float32 {
	return cs.size * 10
}

func PrintArea(ar Ariable) {
	fmt.Println("Area = ", ar.area())
}

type Products []Product

func (products Products) Print() {
	for _, product := range products {
		product.Print()
	}
}

//Implementation of the "Interface" interface in "Products" for sorting
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].id < products[j].id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

func (p *Product) Print() {
	fmt.Printf("Id = %d, Name = %s, Cost = %v, Units = %v, category = %s\n", p.id, p.name, p.cost, p.units, p.category)
}

func main() {
	r := Rect{10, 20}
	/* fmt.Println(r.area())
	fmt.Println(r.perimeter()) */

	c := Circle{10}
	/* fmt.Println(c.area())
	fmt.Println((c.perimeter())) */

	cs := CustomShape{5}
	PrintArea(r)
	PrintArea(c)
	PrintArea(cs)

	p1 := Product{6, "Pen", 10, 50, "Stationary"}
	p2 := Product{5, "Ten", 30, 20, "Grocery"}
	p3 := Product{4, "Den", 20, 10, "Grocery"}
	p4 := Product{1, "Len", 50, 30, "Stationary"}
	p5 := Product{7, "Zen", 40, 20, "Stationary"}
	p6 := Product{9, "Ken", 40, 20, "Stationary"}

	products := Products{p1, p2, p3, p4, p5, p6}

	fmt.Println("Before sorting")
	products.Print()

	sort.Sort(products)

	fmt.Println("After Sorting")
	products.Print()

}
