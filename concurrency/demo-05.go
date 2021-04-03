package main

import "fmt"

func sum(nos []int) int {
	result := 0
	for _, v := range nos {
		result += v
	}
	return result
}

func main() {
	data := []int{8, 2, 5, 3, 7, 1, 9, 6, 11, 14, 17, 57, 14}
	result := sum(data)
	fmt.Println(result)
}
