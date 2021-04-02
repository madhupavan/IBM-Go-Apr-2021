package main

import "fmt"

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

func (p *Product) ApplyDiscount(percent float32) {
	p.cost = float32((100-percent)/100) * p.cost
}

//Collection & Methods
type Products []Product

func (products Products) Print() {
	for _, product := range products {
		product.Print()
	}
}

func (products Products) Index(p Product) int {
	for idx, product := range products {
		if product == p {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(p Product) bool {
	return products.Index(p) >= 0
}

func (products Products) Any(predicate func(p Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(p Product) bool) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

func (products Products) Filter(predicate func(p Product) bool) Products {
	result := Products{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func main() {

	p1 := Product{6, "Pen", 10, 50, "Stationary"}
	//Print(p1)
	p1.Print()
	p1.ApplyDiscount(10)
	fmt.Println("After 10% discount")
	p1.Print()

	//products
	p2 := Product{5, "Ten", 30, 20, "Grocery"}
	p3 := Product{4, "Den", 20, 10, "Grocery"}
	p4 := Product{1, "Len", 50, 30, "Stationary"}
	p5 := Product{7, "Zen", 40, 20, "Stationary"}
	p6 := Product{9, "Ken", 40, 20, "Stationary"}

	products := Products{p1, p2, p3, p4, p5, p6}

	fmt.Println(products.Index(p3))
	fmt.Println("Any stationary products ? :", products.Any(func(p Product) bool {
		return p.category == "Stationary"
	}))
	stationaryProducts := products.Filter(func(p Product) bool {
		return p.category == "Stationary"
	})
	stationaryProducts.Print()

}
