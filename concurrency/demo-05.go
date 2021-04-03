package main

import "fmt"

func sum(nos []int, ch chan int) {
	result := 0
	for _, v := range nos {
		result += v
	}
	ch <- result
}

func main() {
	r1 := make(chan int)
	r2 := make(chan int)

	data := []int{8, 2, 5, 3, 7, 1, 9, 6, 11, 14, 17, 57, 14}
	firstSet := data[:len(data)/2]
	secondSet := data[(len(data) / 2):]
	go sum(firstSet, r1)
	go sum(secondSet, r2)
	finalResult := <-r1 + <-r2
	fmt.Println(finalResult)
}
