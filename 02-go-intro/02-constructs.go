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

}
