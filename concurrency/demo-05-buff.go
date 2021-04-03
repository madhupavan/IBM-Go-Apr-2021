package main

import (
	"fmt"
	"time"
)

func sum(nos []int, sumC chan int) {
	result := <-sumC
	for _, v := range nos {
		result += v
	}
	sumC <- result
}
func sumInChunks(nos []int, chunkSize int) int {
	ch := make(chan int)
	i := 0
	lenNos := len(nos)
	for i <= lenNos {
		j := i + chunkSize
		if j > lenNos {
			go sum(nos[i:], ch)
		} else {
			go sum(nos[i:j], ch)
		}
		i = j
	}
	ch <- 0
	time.Sleep(30 * time.Second)
	return <-ch
}
func main() {
	data := []int{8, 2, 5, 3, 7, 1, 9, 6, 11, 14, 17, 57, 14}
	chunkSize := 3
	result := sumInChunks(data, chunkSize)
	fmt.Println(result)
}
