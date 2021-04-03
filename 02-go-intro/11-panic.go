package main

import "fmt"

func elementAt(nos []int, index int) int {
	//return nos[index]
	if len(nos) > index {
		return nos[index]
	} else {
		panic("index out of bounds")
	}
}
func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Something really bad happened")
		fmt.Println(r)
	}
	fmt.Println("exiting gracefully!")
}

func main() {
	defer recoverFromPanic()
	nos := []int{1, 2, 3}
	fmt.Println(elementAt(nos, 1))
	fmt.Println(elementAt(nos, 5))
}
