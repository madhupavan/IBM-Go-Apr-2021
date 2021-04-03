package main

import (
	"fmt"
)

func fn() {
	fmt.Println("fn started")
	defer fmt.Println("fn completed 1")
	defer fmt.Println("fn completed 2")
	fmt.Println("Something went wrong")
	return
	result := 100
	fmt.Println(result)
}
func main() {
	fn()
}
