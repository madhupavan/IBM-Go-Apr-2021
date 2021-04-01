package main

import (
	"errors"
	"fmt"
)

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

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("Invalid arguments")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("Invalid arguments")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

func getCount() func() int {
	var counter = 0

	count := func() int {
		counter += 1
		return counter
	}
	return count
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

	//Handling errors from functions
	quotient, remainder, err := divide(10, 3)
	if err != nil {
		fmt.Println("Something went wrong!!")
		fmt.Println(err)
		fmt.Printf("quotient = %v, remainder = %v\n", quotient, remainder)
		return
	}
	fmt.Printf("Dividing 10 by 0, quotient = %v, remainder = %v\n", quotient, remainder)

	count := getCount()
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

}
