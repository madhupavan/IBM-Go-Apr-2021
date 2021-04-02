package main

import (
	"fmt"
)

var z int = 100

func add(x, y int) int {
	//fmt.Printf("Processing %d and %d\n", x, y)
	return x + y + z
}

func addClient(x, y int) {
	result := add(x, y)
	fmt.Printf("result = %d\n", result)
}

func print(s string) {
	fmt.Println(s)
}

func main() {
	addClient(100, 200)
	z = 1000
	addClient(100, 200)
	/*
		go print("Hello")
		print("World")
	*/
	/* fmt.Println("end of main") */
	//time.Sleep(3 * time.Second)
}
