package shapes

import "fmt"

type CustomShape struct {
	Size float32
}

func (cs CustomShape) Area() float32 {
	return cs.Size * 10
}

func PrintArea(ar Ariable) {
	fmt.Println("Area = ", ar.Area())
}
