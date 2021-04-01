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

func getCount() func() int { //step 1
	var counter = 0 //step 2

	count := func() int { //step 3
		counter += 1 //step 4
		return counter
	}
	return count //step 5
}

func genNos(predicate func(int) bool) func() int {
	no := 0
	return func() int {
		for {
			no += 1
			if predicate(no) {
				return no
			}
		}
	}
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

	/*
		count := getCount()
		fmt.Println(count())
		fmt.Println(count())
		fmt.Println(count())
		fmt.Println(count())
		fmt.Println(count())
		fmt.Println(count())
	*/
	up, down := func() (func() int, func() int) {
		counter := 0
		up := func() int {
			counter += 1
			return counter
		}
		down := func() int {
			counter -= 1
			return counter
		}
		return up, down
	}()

	fmt.Println(up())
	fmt.Println(up())
	fmt.Println(up())
	fmt.Println(up())

	fmt.Println(down())
	fmt.Println(down())
	fmt.Println(down())
	fmt.Println(down())
	fmt.Println(down())

	fmt.Println("Even numbers")
	isEven := func(no int) bool {
		return no%2 == 0
	}
	evenNos := genNos(isEven)
	for i := 0; i < 10; i++ {
		fmt.Println(evenNos())
	}

	fmt.Println("Prime Numbers")
	isPrime := func(n int) bool {
		if n == 1 {
			return false
		}
		for i := 2; i <= (n / 2); i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	primeNos := genNos(isPrime)
	for i := 0; i < 10; i++ {
		fmt.Println(primeNos())
	}
}

/*
up() => 1
up() => 2
up() => 3
up() => 4

down() => 3
down() => 2
down() => 1
down() => 0 */
