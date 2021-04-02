package main

import "fmt"

func sum(nos ...int) int {
	result := 0
	for _, no := range nos {
		result += no
	}
	return result
}

/*
Use cases:
sum(10,20)
sum(10,20,30,40,50)
sum(10)
sum()
sum(10,20,30,"40") => 100
sum(10, "20", 30 , "abc") => 60

sum(10,20, []int{30,40,50})
sum(10,20, []int{30,40,"50"})
*/

func main() {
	nos := []int{10, 20, 30, 40, 50}
	/* fmt.Println(sum(nos[0], nos[1], nos[2], nos[3], nos[4])) */
	fmt.Println(sum(nos...))
}
