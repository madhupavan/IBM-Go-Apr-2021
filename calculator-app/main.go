package main

import (
	calc "calculator-app/calculator"
	"fmt"
)

func main() {
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println(calc.Multiply(100, 200))
	fmt.Println(calc.Divide(100, 200))
}
