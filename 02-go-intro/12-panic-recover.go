package main

import "fmt"

func elementAt(nos []int, index int) (value int) {
	//return nos[index]
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("Program recovered!")
			value = nos[len(nos)-1]
		}
	}()
	value = nos[index]
	return
}

func main() {
	nos := []int{1, 2, 3}
	fmt.Println(elementAt(nos, 1))
	fmt.Println(elementAt(nos, 5))
}
