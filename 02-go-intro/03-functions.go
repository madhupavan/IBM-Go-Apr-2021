package main

import "fmt"

/*
func add(x, y int) int {
	return x + y
}
*/

/*
var add = func(x, y int) int {
	return x + y
}
*/

/*
func divide(x, y int) (int, int) {
	return x / y, x % y
}
*/

/*
func divide(x, y int) (quotient, remainder int) {
	return x / y, x % y
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}

func main() {
	/*
		add := func(x, y int) int {
			return x + y
		}
		fmt.Printf("Sum of 10 & 20 is %d\n", add(10, 20))
	*/

	/*
		func(x, y int) {
			result := x + y
			fmt.Printf("result = %v\n", result)
		}(100, 200)
	*/

	quotient, remainder := divide(10, 3)
	fmt.Printf("Dividing 10 by 3, quotient = %v, remainder = %v\n", quotient, remainder)

}
