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
	fmt.Println(p)

	pp := PerishableProduct{
		Product{id: 101, name: "Grapes", cost: 60, units: 100, category: "Food"},
		"2 Days",
	}

	fmt.Println(pp)
	fmt.Println(pp.name)
}
