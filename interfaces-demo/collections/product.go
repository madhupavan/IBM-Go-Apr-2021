package collections

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p *Product) Print() {
	fmt.Printf("Id = %d, Name = %s, Cost = %v, Units = %v, category = %s\n", p.Id, p.Name, p.Cost, p.Units, p.Category)
}
