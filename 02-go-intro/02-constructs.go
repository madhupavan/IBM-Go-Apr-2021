package main

import "fmt"

func main() {
	//if else construct
	/*
		no := 42
		if no%2 == 0 {
			fmt.Printf("%v is even\n", no)
		} else {
			fmt.Printf("%v is odd\n", no)
		}
	*/

	if no := 42; no%2 == 0 {
		fmt.Printf("%v is even\n", no)
	} else {
		fmt.Printf("%v is odd\n", no)
	}

	//for construct
	/*
		sum := 0
		for i := 0; i < 9; i++ {
			sum += i
		}
		fmt.Println(sum)
	*/

	//without pre/post statements (while alternative)
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	//infinite for
	/*
		for {

		}
	*/

}
