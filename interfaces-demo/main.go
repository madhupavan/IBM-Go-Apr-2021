package main

import (
	"fmt"
	"interfaces-demo/collections"
	"interfaces-demo/shapes"
)

func main() {
	r := shapes.Rect{10, 20}
	/* fmt.Println(r.area())
	fmt.Println(r.perimeter()) */

	c := shapes.Circle{10}
	/* fmt.Println(c.area())
	fmt.Println((c.perimeter())) */

	cs := shapes.CustomShape{5}
	shapes.PrintArea(r)
	shapes.PrintArea(c)
	shapes.PrintArea(cs)

	p1 := collections.Product{6, "Pen", 10, 50, "Stationary"}
	p2 := collections.Product{5, "Ten", 30, 20, "Grocery"}
	p3 := collections.Product{4, "Den", 20, 10, "Grocery"}
	p4 := collections.Product{1, "Len", 50, 30, "Stationary"}
	p5 := collections.Product{7, "Zen", 40, 20, "Stationary"}
	p6 := collections.Product{9, "Ken", 40, 20, "Stationary"}

	products := collections.Products{p1, p2, p3, p4, p5, p6}

	fmt.Println("Before sorting")
	products.Print()

	fmt.Println("After Sorting  - default [id]")
	products.Sort("id")
	products.Print()

	fmt.Println("After Sorting  - default [name]")
	products.Sort("name")
	products.Print()

	fmt.Println("After Sorting  - default [cost]")
	products.Sort("cost")
	products.Print()

	fmt.Println("After Sorting  - default [category]")
	products.Sort("category")
	products.Print()
}
