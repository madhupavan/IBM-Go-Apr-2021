package main

import "fmt"

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

/*
type PerishableProduct struct {
	product Product
	expiry  string
}
*/

type PerishableProduct struct {
	Product
	expiry string
}

func main() {

	//p := Product{id: 100, name: "Pen", cost: 10, units: 50, category: "Stationary"}
	p := Product{100, "Pen", 10, 50, "Stationary"}
	var pptr *Product = &p
	fmt.Println(p.name)
	fmt.Println(pptr.name)

	pp := PerishableProduct{
		Product{id: 101, name: "Grapes", cost: 60, units: 100, category: "Food"},
		"2 Days",
	}

	fmt.Println(pp)
	fmt.Println(pp.name)

	//products
	p1 := Product{6, "Pen", 10, 50, "Stationary"}
	p2 := Product{5, "Ten", 30, 20, "Grocery"}
	p3 := Product{4, "Den", 20, 10, "Grocery"}
	p4 := Product{1, "Len", 50, 30, "Stationary"}
	p5 := Product{7, "Zen", 40, 20, "Stationary"}
	p6 := Product{9, "Ken", 40, 20, "Stationary"}

	products := []Product{p1, p2, p3, p4, p5, p6}

	fmt.Println(Index(products, p3))
	fmt.Println("Any stationary products ? :", Any(products, func(p Product) bool {
		return p.category == "Stationary"
	}))
}

func Index(products []Product, p Product) int {
	for idx, product := range products {
		if product == p {
			return idx
		}
	}
	return -1
}

func Includes(products []Product, p Product) bool {
	return Index(products, p) >= 0
}

func Any(products []Product, predicate func(p Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func All(products []Product, predicate func(p Product) bool) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

func Filter(products []Product, predicate func(p Product) bool) []Product {
	result := make([]Product, 0)
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

/*
Index(products, product) => index of product in the products
Includes(products, product) => true/false
Any(products, predicate) => true/false
Filter(products, predicate) => products
*/
