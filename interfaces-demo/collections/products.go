package collections

import "sort"

type Products []Product

func (products Products) Print() {
	for _, product := range products {
		product.Print()
	}
}

var productsComparers map[string]func(Products, int, int) bool = map[string]func(Products, int, int) bool{
	"name": func(products Products, i, j int) bool {
		return products[i].Name < products[j].Name
	},

	"id": func(products Products, i, j int) bool {
		return products[i].Id < products[j].Id
	},

	"cost": func(products Products, i, j int) bool {
		return products[i].Cost < products[j].Cost
	},

	"units": func(products Products, i, j int) bool {
		return products[i].Units < products[j].Units
	},

	"category": func(products Products, i, j int) bool {
		return products[i].Category < products[j].Category
	},
}

var productsComparer func(products Products, i, j int) bool

func (products Products) Sort(attrName string) {
	productsComparer = productsComparers[attrName]
	sort.Sort(products)
}

//Implementation of the "Interface" interface in "Products" for sorting

func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return productsComparer(products, i, j)
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}
